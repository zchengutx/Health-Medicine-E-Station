package comment

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocket升级器配置
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有来源，生产环境中应该更严格
		return true
	},
}

// 聊天消息结构
type ChatMessage struct {
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	FromID    int32     `json:"from_id"`
	ToID      int32     `json:"to_id"`
	FromName  string    `json:"from_name"`
	Timestamp time.Time `json:"timestamp"`
}

// 客户端连接结构
type Client struct {
	ID     int32
	Name   string
	Role   string // "doctor" 或 "patient"
	Conn   *websocket.Conn
	Send   chan ChatMessage
}

// 一对一聊天管理器
type ChatManager struct {
	clients    map[int32]*Client // key: userID
	register   chan *Client
	unregister chan *Client
	message    chan ChatMessage
	mutex      sync.RWMutex
}

// 全局聊天管理器实例
var chatManager *ChatManager

// 初始化聊天管理器
func init() {
	chatManager = &ChatManager{
		clients:    make(map[int32]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		message:    make(chan ChatMessage),
	}
	go chatManager.run()
}

// 运行聊天管理器
func (cm *ChatManager) run() {
	for {
		select {
		case client := <-cm.register:
			cm.registerClient(client)

		case client := <-cm.unregister:
			cm.unregisterClient(client)

		case message := <-cm.message:
			cm.handleMessage(message)
		}
	}
}

// 注册客户端
func (cm *ChatManager) registerClient(client *Client) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.clients[client.ID] = client
	log.Printf("客户端 %d (%s, %s) 已连接", client.ID, client.Name, client.Role)
}

// 注销客户端
func (cm *ChatManager) unregisterClient(client *Client) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if _, exists := cm.clients[client.ID]; exists {
		delete(cm.clients, client.ID)
		close(client.Send)
		log.Printf("客户端 %d (%s) 断开连接", client.ID, client.Name)
	}
}

// 处理消息（一对一发送）
func (cm *ChatManager) handleMessage(message ChatMessage) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	// 找到目标用户的连接
	if targetClient, exists := cm.clients[message.ToID]; exists {
		select {
		case targetClient.Send <- message:
			log.Printf("消息从 %d 发送到 %d: %s", message.FromID, message.ToID, message.Content)
		default:
			// 如果发送失败，关闭连接
			close(targetClient.Send)
			delete(cm.clients, message.ToID)
		}
	} else {
		log.Printf("目标用户 %d 不在线，消息未发送", message.ToID)
	}
}

// WebSocket连接处理器
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 打印所有请求头用于调试
	log.Printf("WebSocket连接请求，Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			log.Printf("  %s: %s", name, value)
		}
	}

	// 从Header获取token和目标用户信息
	token := r.Header.Get("Authorization")
	userRole := r.Header.Get("User-Role") // "doctor" 或 "patient"
	targetIDStr := r.Header.Get("Target-Id")

	log.Printf("解析到的参数: token=%s, userRole=%s, targetID=%s", token, userRole, targetIDStr)

	// 处理Authorization header中的Bearer前缀
	if token != "" && len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
		log.Printf("去除Bearer前缀后的token: %s", token)
	}

	// 如果Header中有认证信息，先验证
	if token != "" && userRole != "" && targetIDStr != "" {
		log.Printf("Header中有完整认证信息，开始验证")
		
		// 验证JWT token
		claims, errMsg := GetToken(token)
		if claims == nil || errMsg != "" {
			log.Printf("Token验证失败: %s", errMsg)
			http.Error(w, "token无效: "+errMsg, http.StatusUnauthorized)
			return
		}

		userIDFloat, ok := claims["user"].(float64)
		if !ok {
			log.Printf("Token中用户ID格式错误")
			http.Error(w, "token中用户ID格式错误", http.StatusUnauthorized)
			return
		}
		userID := int32(userIDFloat)

		targetID, err := strconv.Atoi(targetIDStr)
		if err != nil {
			log.Printf("无效的目标用户ID: %s", targetIDStr)
			http.Error(w, "无效的目标用户ID", http.StatusBadRequest)
			return
		}

		log.Printf("认证成功，用户ID: %d, 目标ID: %d", userID, targetID)

		// 升级HTTP连接为WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("WebSocket升级失败: %v", err)
			return
		}

		log.Printf("WebSocket连接升级成功")

		// 处理已认证的连接
		handleAuthenticatedWebSocket(conn, userID, userRole, int32(targetID))
		return
	}

	// 如果Header中没有完整信息，升级连接后等待认证消息
	log.Printf("Header信息不完整，升级连接后等待认证消息")
	
	// 升级HTTP连接为WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket升级失败: %v", err)
		return
	}

	log.Printf("WebSocket连接升级成功，等待认证消息")
	handleUnauthenticatedConnection(conn)
}

// 处理已认证的WebSocket连接
func handleAuthenticatedWebSocket(conn *websocket.Conn, userID int32, userRole string, targetID int32) {
	log.Printf("处理已认证的WebSocket连接: userID=%d, userRole=%s, targetID=%d", userID, userRole, targetID)

	// 如果是患者连接医生，检查医生是否在线
	if userRole == "patient" {
		if _, exists := chatManager.clients[targetID]; !exists {
			log.Printf("目标医生 %d 不在线", targetID)
			conn.WriteJSON(map[string]interface{}{
				"type":    "error",
				"message": "医生当前不在线，无法建立连接",
			})
			conn.Close()
			return
		}
	}

	// 创建客户端
	client := &Client{
		ID:   userID,
		Name: "", // 用户名将在需要时从数据库查询
		Role: userRole,
		Conn: conn,
		Send: make(chan ChatMessage, 256),
	}

	// 注册客户端
	chatManager.register <- client

	// 发送连接成功消息
	conn.WriteJSON(map[string]interface{}{
		"type":    "system",
		"message": "连接成功",
	})

	log.Printf("客户端 %d 注册成功", userID)

	// 启动读写协程
	go client.writePump()
	go client.readPump(targetID)
}

// 处理未认证的连接（等待认证消息）
func handleUnauthenticatedConnection(conn *websocket.Conn) {
	defer conn.Close()

	// 设置读取超时
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	// 等待认证消息
	var authMsg map[string]interface{}
	err := conn.ReadJSON(&authMsg)
	if err != nil {
		log.Printf("读取认证消息失败: %v", err)
		return
	}

	// 检查是否为认证消息
	if msgType, ok := authMsg["type"].(string); !ok || msgType != "auth" {
		conn.WriteJSON(map[string]interface{}{
			"type":    "error",
			"message": "需要先发送认证消息",
		})
		return
	}

	// 提取认证信息
	token, _ := authMsg["token"].(string)
	userRole, _ := authMsg["user_role"].(string)
	targetIDFloat, _ := authMsg["target_id"].(float64)
	targetIDStr := strconv.Itoa(int(targetIDFloat))

	if token == "" || userRole == "" || targetIDStr == "" {
		conn.WriteJSON(map[string]interface{}{
			"type":    "error",
			"message": "认证信息不完整",
		})
		return
	}

	// 验证JWT token
	claims, errMsg := GetToken(token)
	if claims == nil || errMsg != "" {
		conn.WriteJSON(map[string]interface{}{
			"type":    "error",
			"message": "token无效: " + errMsg,
		})
		return
	}

	userIDFloat, ok := claims["user"].(float64)
	if !ok {
		conn.WriteJSON(map[string]interface{}{
			"type":    "error",
			"message": "token中用户ID格式错误",
		})
		return
	}
	userID := int32(userIDFloat)

	targetID, err := strconv.Atoi(targetIDStr)
	if err != nil {
		conn.WriteJSON(map[string]interface{}{
			"type":    "error",
			"message": "无效的目标用户ID",
		})
		return
	}

	// 使用认证信息建立连接
	handleAuthenticatedWebSocket(conn, userID, userRole, int32(targetID))
}

// 读取消息
func (c *Client) readPump(targetID int32) {
	defer func() {
		chatManager.unregister <- c
		c.Conn.Close()
	}()

	// 设置读取超时
	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		var message ChatMessage
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket错误: %v", err)
			}
			break
		}

		// 设置消息信息
		message.FromID = c.ID
		message.ToID = targetID
		message.FromName = c.Name
		message.Timestamp = time.Now()

		// 发送消息给目标用户
		chatManager.message <- message
	}
}

// 写入消息
func (c *Client) writePump() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.Conn.WriteJSON(message); err != nil {
				log.Printf("写入消息失败: %v", err)
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// 获取在线用户列表
func GetOnlineUsers() []map[string]interface{} {
	chatManager.mutex.RLock()
	defer chatManager.mutex.RUnlock()

	users := make([]map[string]interface{}, 0)
	for _, client := range chatManager.clients {
		users = append(users, map[string]interface{}{
			"id":   client.ID,
			"name": client.Name,
			"role": client.Role,
		})
	}

	return users
}

// 检查用户是否在线
func IsUserOnline(userID int32) bool {
	chatManager.mutex.RLock()
	defer chatManager.mutex.RUnlock()

	_, exists := chatManager.clients[userID]
	return exists
}

// 发送消息到特定用户
func SendMessageToUser(fromID, toID int32, fromName, content string) error {
	message := ChatMessage{
		Type:      "text",
		Content:   content,
		FromID:    fromID,
		ToID:      toID,
		FromName:  fromName,
		Timestamp: time.Now(),
	}

	chatManager.message <- message
	return nil
}

// 简化的WebSocket测试处理器（用于调试）
func HandleTestWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接为WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket升级失败: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("WebSocket测试连接建立成功")

	// 发送欢迎消息
	welcomeMsg := map[string]interface{}{
		"type":    "system",
		"content": "WebSocket连接测试成功！",
		"time":    time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := conn.WriteJSON(welcomeMsg); err != nil {
		log.Printf("发送欢迎消息失败: %v", err)
		return
	}

	// 简单的回声服务器
	for {
		var msg map[string]interface{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket错误: %v", err)
			}
			break
		}

		log.Printf("收到消息: %+v", msg)

		// 回显消息
		var content string
		if msgContent, ok := msg["content"]; ok {
			content = "收到你的消息: " + msgContent.(string)
		} else {
			content = "收到你的消息"
		}
		
		response := map[string]interface{}{
			"type":     "echo",
			"content":  content,
			"original": msg,
			"time":     time.Now().Format("2006-01-02 15:04:05"),
		}

		if err := conn.WriteJSON(response); err != nil {
			log.Printf("发送回显消息失败: %v", err)
			break
		}
	}
}
package comment

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/websocket"
	"log"
	https "net/http"
	"strconv"
	"sync"
	"time"
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
	RoomID    string    `json:"room_id"`
}

// 客户端连接结构
type Client struct {
	ID     int32
	Name   string
	Role   string // "doctor" 或 "patient"
	Conn   *websocket.Conn
	Send   chan ChatMessage
	RoomID string
}

// 一对一聊天管理器
type ChatManager struct {
	clients    map[string]*Client // key: roomID_userID
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
		clients:    make(map[string]*Client),
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

	clientKey := client.RoomID + "_" + strconv.Itoa(int(client.ID))
	cm.clients[clientKey] = client

	log.Printf("客户端 %d (%s) 连接到房间 %s", client.ID, client.Name, client.RoomID)
}

// 注销客户端
func (cm *ChatManager) unregisterClient(client *Client) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	clientKey := client.RoomID + "_" + strconv.Itoa(int(client.ID))
	if _, exists := cm.clients[clientKey]; exists {
		delete(cm.clients, clientKey)
		close(client.Send)

		// Kiro修改：用户断开连接时清理心跳状态
		userIDStr := strconv.Itoa(int(client.ID))
		RemoveUserHeartbeat(userIDStr)
		log.Printf("Kiro修改：客户端 %d (%s) 断开连接，已清理心跳状态", client.ID, client.Name)
	}
}

// 处理消息（一对一发送）
func (cm *ChatManager) handleMessage(message ChatMessage) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	// Kiro修改：检查目标用户是否在线
	targetUserID := strconv.Itoa(int(message.ToID))
	if !IsUserOnline(targetUserID) {
		log.Printf("Kiro修改：目标用户 %d 不在线，消息发送失败", message.ToID)

		// Kiro修改：向发送者返回错误消息
		senderClientKey := message.RoomID + "_" + strconv.Itoa(int(message.FromID))
		if senderClient, exists := cm.clients[senderClientKey]; exists {
			errorMessage := ChatMessage{
				Type:      "error",
				Content:   "目标用户不在线，消息发送失败",
				FromID:    0, // 系统消息
				ToID:      message.FromID,
				FromName:  "系统",
				Timestamp: time.Now(),
				RoomID:    message.RoomID,
			}
			select {
			case senderClient.Send <- errorMessage:
				log.Printf("Kiro修改：已向用户 %d 发送离线错误提示", message.FromID)
			default:
				log.Printf("Kiro修改：向用户 %d 发送离线错误提示失败", message.FromID)
			}
		}
		return
	}

	// 找到目标用户的连接
	targetClientKey := message.RoomID + "_" + strconv.Itoa(int(message.ToID))
	if targetClient, exists := cm.clients[targetClientKey]; exists {
		select {
		case targetClient.Send <- message:
			log.Printf("消息从 %d 发送到 %d: %s", message.FromID, message.ToID, message.Content)
		default:
			// 如果发送失败，关闭连接
			close(targetClient.Send)
			delete(cm.clients, targetClientKey)
		}
	}
}

// 生成房间ID（基于两个用户ID）
func generateRoomID(userID1, userID2 int32) string {
	if userID1 > userID2 {
		userID1, userID2 = userID2, userID1
	}
	return "room_" + strconv.Itoa(int(userID1)) + "_" + strconv.Itoa(int(userID2))
}

// WebSocket连接处理器
func HandleWebSocket(w http.ResponseWriter, r *http.Request, ctx http.Context) {
	// 从查询参数获取用户信息
	req := ctx.Request()
	value := req.Context().Value("user_id")
	userRole := r.URL.Query().Get("user_role") // "doctor" 或 "patient"
	targetIDStr := r.URL.Query().Get("target_id")

	if targetIDStr == "" {
		https.Error(w, "缺少必要参数", https.StatusBadRequest)
		return
	}

	// 检查userID是否存在且有效
	if value == nil {
		https.Error(w, "Unauthorized", https.StatusUnauthorized)
		return
	}

	userIDFloat, ok := value.(float64)
	if !ok || userIDFloat == 0 {
		https.Error(w, "无效的用户ID", https.StatusBadRequest)
		return
	}

	userID := int32(userIDFloat)
	userIDStr := strconv.Itoa(int(userID))

	targetID, err := strconv.Atoi(targetIDStr)
	if err != nil {
		https.Error(w, "无效的目标用户ID", https.StatusBadRequest)
		return
	}

	// 升级HTTP连接为WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket升级失败: %v", err)
		return
	}

	// Kiro修改：立即更新用户心跳状态，标记为在线
	UpdateUserHeartbeat(userIDStr)
	log.Printf("Kiro修改：用户 %s 连接WebSocket，已更新心跳状态", userIDStr)

	// 生成房间ID
	roomID := generateRoomID(userID, int32(targetID))

	// 创建客户端
	client := &Client{
		ID:     userID,
		Name:   "用户" + userIDStr, // 可以从数据库获取真实姓名
		Role:   userRole,
		Conn:   conn,
		Send:   make(chan ChatMessage, 256),
		RoomID: roomID, // Kiro修改：设置正确的房间ID
	}

	// 注册客户端
	chatManager.register <- client

	// 启动读写协程
	go client.writePump()
	go client.readPump(int32(targetID))

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
		var rawMessage map[string]interface{}
		err := c.Conn.ReadJSON(&rawMessage)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket错误: %v", err)
			}
			break
		}

		// Kiro修改：处理心跳消息
		if msgType, ok := rawMessage["type"].(string); ok && msgType == "heartbeat" {
			// 更新用户心跳时间
			userIDStr := strconv.Itoa(int(c.ID))
			UpdateUserHeartbeat(userIDStr)

			// 发送心跳响应
			heartbeatResponse := CreateHeartbeatResponse()
			if err := c.Conn.WriteJSON(heartbeatResponse); err != nil {
				log.Printf("Kiro修改：发送心跳响应失败: %v", err)
				break
			}
			log.Printf("Kiro修改：处理用户 %d 的心跳消息", c.ID)
			continue
		}

		// Kiro修改：处理普通聊天消息
		var message ChatMessage
		if content, ok := rawMessage["content"].(string); ok {
			message.Type = "text"
			message.Content = content
		} else {
			continue // 跳过无效消息
		}

		// 设置消息信息
		message.FromID = c.ID
		message.ToID = targetID
		message.FromName = c.Name
		message.Timestamp = time.Now()
		message.RoomID = c.RoomID

		// Kiro修改：每次收到消息时也更新心跳（表示用户活跃）
		userIDStr := strconv.Itoa(int(c.ID))
		UpdateUserHeartbeat(userIDStr)

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

// 获取房间中的在线用户
func GetRoomUsers(roomID string) []map[string]interface{} {
	chatManager.mutex.RLock()
	defer chatManager.mutex.RUnlock()

	users := make([]map[string]interface{}, 0)
	for _, client := range chatManager.clients {
		if client.RoomID == roomID {
			users = append(users, map[string]interface{}{
				"id":   client.ID,
				"name": client.Name,
				"role": client.Role,
			})
		}
	}

	return users
}

// 发送消息到特定用户
func SendMessageToUser(roomID string, fromID, toID int32, fromName, content string) error {
	message := ChatMessage{
		Type:      "text",
		Content:   content,
		FromID:    fromID,
		ToID:      toID,
		FromName:  fromName,
		Timestamp: time.Now(),
		RoomID:    roomID,
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

		// Kiro修改：处理心跳消息
		if msgType, ok := msg["type"].(string); ok && msgType == "heartbeat" {
			// 更新心跳时间
			UpdateUserHeartbeat("test_user")
			// 发送心跳响应
			heartbeatResponse := CreateHeartbeatResponse()
			if err := conn.WriteJSON(heartbeatResponse); err != nil {
				log.Printf("发送心跳响应失败: %v", err)
				break
			}
			continue
		}

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

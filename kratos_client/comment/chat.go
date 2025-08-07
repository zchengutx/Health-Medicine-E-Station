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
		log.Printf("客户端 %d (%s) 断开连接", client.ID, client.Name)
	}
}

// 处理消息（一对一发送）
func (cm *ChatManager) handleMessage(message ChatMessage) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

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

	// 创建客户端
	client := &Client{
		ID:   int32(value.(float64)),
		Role: userRole,
		Conn: conn,
		Send: make(chan ChatMessage, 256),
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
		message.RoomID = c.RoomID

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

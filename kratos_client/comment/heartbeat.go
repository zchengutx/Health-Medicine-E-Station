package comment

import (
	"encoding/json"
	"log"
	"sync"
	"time"
)

// Kiro修改：添加心跳检测功能
// HeartbeatManager 心跳管理器
type HeartbeatManager struct {
	// Kiro修改：存储用户最后心跳时间
	userHeartbeats map[string]time.Time
	// Kiro修改：读写锁保护并发访问
	mutex sync.RWMutex
	// Kiro修改：心跳超时时间（默认30秒）
	timeout time.Duration
	// Kiro修改：清理间隔时间（默认10秒）
	cleanupInterval time.Duration
}

// Kiro修改：全局心跳管理器实例
var heartbeatManager *HeartbeatManager

// Kiro修改：初始化心跳管理器
func init() {
	heartbeatManager = &HeartbeatManager{
		userHeartbeats:  make(map[string]time.Time),
		timeout:         30 * time.Second, // 30秒超时
		cleanupInterval: 10 * time.Second, // 10秒清理一次
	}
	// Kiro修改：启动后台清理协程
	go heartbeatManager.startCleanupRoutine()
}

// Kiro修改：更新用户心跳时间
func (hm *HeartbeatManager) UpdateHeartbeat(userID string) {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	hm.userHeartbeats[userID] = time.Now()
	log.Printf("Kiro修改：用户 %s 心跳更新", userID)
}

// Kiro修改：检查用户是否在线
func (hm *HeartbeatManager) IsUserOnline(userID string) bool {
	hm.mutex.RLock()
	defer hm.mutex.RUnlock()
	
	lastHeartbeat, exists := hm.userHeartbeats[userID]
	if !exists {
		return false
	}
	
	// Kiro修改：检查是否超时
	return time.Since(lastHeartbeat) <= hm.timeout
}

// Kiro修改：获取所有在线用户
func (hm *HeartbeatManager) GetOnlineUsers() []string {
	hm.mutex.RLock()
	defer hm.mutex.RUnlock()
	
	var onlineUsers []string
	now := time.Now()
	
	for userID, lastHeartbeat := range hm.userHeartbeats {
		if now.Sub(lastHeartbeat) <= hm.timeout {
			onlineUsers = append(onlineUsers, userID)
		}
	}
	
	log.Printf("Kiro修改：当前在线用户数量: %d", len(onlineUsers))
	return onlineUsers
}

// Kiro修改：移除用户心跳记录
func (hm *HeartbeatManager) RemoveUser(userID string) {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	delete(hm.userHeartbeats, userID)
	log.Printf("Kiro修改：用户 %s 心跳记录已移除", userID)
}

// Kiro修改：后台清理过期心跳记录
func (hm *HeartbeatManager) startCleanupRoutine() {
	ticker := time.NewTicker(hm.cleanupInterval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			hm.cleanupExpiredHeartbeats()
		}
	}
}

// Kiro修改：清理过期的心跳记录
func (hm *HeartbeatManager) cleanupExpiredHeartbeats() {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	
	now := time.Now()
	expiredUsers := make([]string, 0)
	
	for userID, lastHeartbeat := range hm.userHeartbeats {
		if now.Sub(lastHeartbeat) > hm.timeout {
			expiredUsers = append(expiredUsers, userID)
		}
	}
	
	// Kiro修改：删除过期用户
	for _, userID := range expiredUsers {
		delete(hm.userHeartbeats, userID)
		log.Printf("Kiro修改：用户 %s 心跳超时，已标记为离线", userID)
	}
	
	if len(expiredUsers) > 0 {
		log.Printf("Kiro修改：清理了 %d 个过期心跳记录", len(expiredUsers))
	}
}

// Kiro修改：获取用户在线状态详情
func (hm *HeartbeatManager) GetUserStatus(userID string) map[string]interface{} {
	hm.mutex.RLock()
	defer hm.mutex.RUnlock()
	
	status := map[string]interface{}{
		"user_id": userID,
		"online":  false,
		"last_heartbeat": nil,
	}
	
	if lastHeartbeat, exists := hm.userHeartbeats[userID]; exists {
		status["online"] = time.Since(lastHeartbeat) <= hm.timeout
		status["last_heartbeat"] = lastHeartbeat.Format("2006-01-02 15:04:05")
	}
	
	return status
}

// Kiro修改：处理心跳消息
func HandleHeartbeatMessage(userID string, message map[string]interface{}) {
	if msgType, ok := message["type"].(string); ok && msgType == "heartbeat" {
		heartbeatManager.UpdateHeartbeat(userID)
		log.Printf("Kiro修改：收到用户 %s 的心跳消息", userID)
	}
}

// Kiro修改：创建心跳响应消息
func CreateHeartbeatResponse() map[string]interface{} {
	return map[string]interface{}{
		"type":      "heartbeat_response",
		"timestamp": time.Now().Unix(),
		"message":   "heartbeat received",
	}
}

// Kiro修改：公开的API函数
// UpdateUserHeartbeat 更新用户心跳
func UpdateUserHeartbeat(userID string) {
	heartbeatManager.UpdateHeartbeat(userID)
}

// IsUserOnline 检查用户是否在线
func IsUserOnline(userID string) bool {
	return heartbeatManager.IsUserOnline(userID)
}

// GetAllOnlineUsers 获取所有在线用户
func GetAllOnlineUsers() []string {
	return heartbeatManager.GetOnlineUsers()
}

// RemoveUserHeartbeat 移除用户心跳记录
func RemoveUserHeartbeat(userID string) {
	heartbeatManager.RemoveUser(userID)
}

// GetUserOnlineStatus 获取用户在线状态
func GetUserOnlineStatus(userID string) map[string]interface{} {
	return heartbeatManager.GetUserStatus(userID)
}

// Kiro修改：创建心跳检测的WebSocket消息处理器
func ProcessWebSocketMessage(userID string, messageData []byte) ([]byte, bool) {
	var message map[string]interface{}
	if err := json.Unmarshal(messageData, &message); err != nil {
		return nil, false
	}
	
	// Kiro修改：检查是否是心跳消息
	if msgType, ok := message["type"].(string); ok && msgType == "heartbeat" {
		HandleHeartbeatMessage(userID, message)
		response := CreateHeartbeatResponse()
		responseData, _ := json.Marshal(response)
		return responseData, true
	}
	
	return nil, false
}
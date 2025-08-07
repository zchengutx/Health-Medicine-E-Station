package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// 聊天消息模型
type MtChatMessage struct {
	ID          int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement;" json:"id"`
	FromID      int32     `gorm:"column:from_id;type:int;comment:发送者ID;not null;" json:"from_id"`
	ToID        int32     `gorm:"column:to_id;type:int;comment:接收者ID;not null;" json:"to_id"`
	Content     string    `gorm:"column:content;type:text;comment:消息内容;" json:"content"`
	MessageType string    `gorm:"column:message_type;type:varchar(20);comment:消息类型;not null;default:'text';" json:"message_type"`
	RoomID      string    `gorm:"column:room_id;type:varchar(50);comment:房间ID;not null;" json:"room_id"`
	IsRead      bool      `gorm:"column:is_read;type:tinyint(1);comment:是否已读;not null;default:0;" json:"is_read"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);" json:"updated_at"`
}

func (m *MtChatMessage) TableName() string {
	return "mt_chat_message"
}

// 聊天房间模型
type MtChatRoom struct {
	ID              int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement;" json:"id"`
	RoomID          string    `gorm:"column:room_id;type:varchar(50);comment:房间ID;not null;uniqueIndex;" json:"room_id"`
	User1ID         int32     `gorm:"column:user1_id;type:int;comment:用户1ID;not null;" json:"user1_id"`
	User2ID         int32     `gorm:"column:user2_id;type:int;comment:用户2ID;not null;" json:"user2_id"`
	LastMessage     string    `gorm:"column:last_message;type:text;comment:最后一条消息;" json:"last_message"`
	LastMessageTime time.Time `gorm:"column:last_message_time;type:datetime(6);comment:最后消息时间;" json:"last_message_time"`
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);" json:"updated_at"`
}

func (m *MtChatRoom) TableName() string {
	return "mt_chat_room"
}

// 聊天仓库接口
type ChatRepo interface {
	// 保存聊天消息
	SaveChatMessage(ctx context.Context, message *MtChatMessage) error

	// 获取聊天历史记录
	GetChatHistory(ctx context.Context, userID, targetID int32, page, pageSize int32) ([]*MtChatMessage, int32, error)

	// 创建或更新聊天房间
	CreateOrUpdateChatRoom(ctx context.Context, room *MtChatRoom) error

	// 获取用户的聊天房间列表
	GetUserChatRooms(ctx context.Context, userID int32) ([]*MtChatRoom, error)

	// 标记消息为已读
	MarkMessagesAsRead(ctx context.Context, roomID string, userID int32) error

	// 获取未读消息数量
	GetUnreadCount(ctx context.Context, roomID string, userID int32) (int32, error)
}

// 聊天服务
type ChatService struct {
	repo ChatRepo
	log  *log.Helper
}

// 创建聊天服务
func NewChatService(repo ChatRepo, logger log.Logger) *ChatService {
	return &ChatService{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// 保存聊天消息
func (cs *ChatService) SaveChatMessage(ctx context.Context, message *MtChatMessage) error {
	cs.log.WithContext(ctx).Infof("SaveChatMessage: %+v", message)

	// 保存消息
	if err := cs.repo.SaveChatMessage(ctx, message); err != nil {
		return err
	}

	// 更新聊天房间的最后消息
	room := &MtChatRoom{
		RoomID:          message.RoomID,
		User1ID:         message.FromID,
		User2ID:         message.ToID,
		LastMessage:     message.Content,
		LastMessageTime: message.CreatedAt,
	}

	return cs.repo.CreateOrUpdateChatRoom(ctx, room)
}

// 获取聊天历史记录
func (cs *ChatService) GetChatHistory(ctx context.Context, userID, targetID int32, page, pageSize int32) ([]*MtChatMessage, int32, error) {
	cs.log.WithContext(ctx).Infof("GetChatHistory: userID=%d, targetID=%d, page=%d, pageSize=%d", userID, targetID, page, pageSize)
	return cs.repo.GetChatHistory(ctx, userID, targetID, page, pageSize)
}

// 获取用户的聊天房间列表
func (cs *ChatService) GetUserChatRooms(ctx context.Context, userID int32) ([]*MtChatRoom, error) {
	cs.log.WithContext(ctx).Infof("GetUserChatRooms: userID=%d", userID)
	return cs.repo.GetUserChatRooms(ctx, userID)
}

// 标记消息为已读
func (cs *ChatService) MarkMessagesAsRead(ctx context.Context, roomID string, userID int32) error {
	cs.log.WithContext(ctx).Infof("MarkMessagesAsRead: roomID=%s, userID=%d", roomID, userID)
	return cs.repo.MarkMessagesAsRead(ctx, roomID, userID)
}

// 获取未读消息数量
func (cs *ChatService) GetUnreadCount(ctx context.Context, roomID string, userID int32) (int32, error) {
	cs.log.WithContext(ctx).Infof("GetUnreadCount: roomID=%s, userID=%d", roomID, userID)
	return cs.repo.GetUnreadCount(ctx, roomID, userID)
}

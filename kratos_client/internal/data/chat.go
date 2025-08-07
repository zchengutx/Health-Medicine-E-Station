package data

import (
	"context"
	"strconv"

	"kratos_client/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type chatRepo struct {
	data *Data
	log  *log.Helper
}

// NewChatRepo 创建聊天仓库
func NewChatRepo(data *Data, logger log.Logger) biz.ChatRepo {
	return &chatRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 生成房间ID
func generateRoomID(userID1, userID2 int32) string {
	if userID1 > userID2 {
		userID1, userID2 = userID2, userID1
	}
	return "room_" + strconv.Itoa(int(userID1)) + "_" + strconv.Itoa(int(userID2))
}

// 保存聊天消息
func (r *chatRepo) SaveChatMessage(ctx context.Context, message *biz.MtChatMessage) error {
	return r.data.Db.WithContext(ctx).Create(message).Error
}

// 获取聊天历史记录
func (r *chatRepo) GetChatHistory(ctx context.Context, userID, targetID int32, page, pageSize int32) ([]*biz.MtChatMessage, int32, error) {
	roomID := generateRoomID(userID, targetID)
	
	var messages []*biz.MtChatMessage
	var total int64
	
	// 计算总数
	if err := r.data.Db.WithContext(ctx).Model(&biz.MtChatMessage{}).
		Where("room_id = ?", roomID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询消息，按时间倒序
	offset := (page - 1) * pageSize
	if err := r.data.Db.WithContext(ctx).
		Where("room_id = ?", roomID).
		Order("created_at DESC").
		Limit(int(pageSize)).
		Offset(int(offset)).
		Find(&messages).Error; err != nil {
		return nil, 0, err
	}
	
	return messages, int32(total), nil
}

// 创建或更新聊天房间
func (r *chatRepo) CreateOrUpdateChatRoom(ctx context.Context, room *biz.MtChatRoom) error {
	// 确保user1_id < user2_id
	if room.User1ID > room.User2ID {
		room.User1ID, room.User2ID = room.User2ID, room.User1ID
	}
	
	// 尝试更新现有房间
	result := r.data.Db.WithContext(ctx).
		Where("room_id = ?", room.RoomID).
		Updates(map[string]interface{}{
			"last_message":      room.LastMessage,
			"last_message_time": room.LastMessageTime,
		})
	
	if result.Error != nil {
		return result.Error
	}
	
	// 如果没有找到现有房间，创建新房间
	if result.RowsAffected == 0 {
		return r.data.Db.WithContext(ctx).Create(room).Error
	}
	
	return nil
}

// 获取用户的聊天房间列表
func (r *chatRepo) GetUserChatRooms(ctx context.Context, userID int32) ([]*biz.MtChatRoom, error) {
	var rooms []*biz.MtChatRoom
	
	err := r.data.Db.WithContext(ctx).
		Where("user1_id = ? OR user2_id = ?", userID, userID).
		Order("last_message_time DESC").
		Find(&rooms).Error
	
	return rooms, err
}

// 标记消息为已读
func (r *chatRepo) MarkMessagesAsRead(ctx context.Context, roomID string, userID int32) error {
	return r.data.Db.WithContext(ctx).
		Model(&biz.MtChatMessage{}).
		Where("room_id = ? AND to_id = ? AND is_read = ?", roomID, userID, false).
		Update("is_read", true).Error
}

// 获取未读消息数量
func (r *chatRepo) GetUnreadCount(ctx context.Context, roomID string, userID int32) (int32, error) {
	var count int64
	
	err := r.data.Db.WithContext(ctx).
		Model(&biz.MtChatMessage{}).
		Where("room_id = ? AND to_id = ? AND is_read = ?", roomID, userID, false).
		Count(&count).Error
	
	return int32(count), err
}
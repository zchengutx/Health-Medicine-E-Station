package service

import (
	"context"
	chatv1 "kratos_client/api/chat/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type ChatService struct {
	chatv1.UnimplementedChatServer
	log  *log.Helper
}

func NewChatService(logger log.Logger) *ChatService {
	return &ChatService{
		UnimplementedChatServer: chatv1.UnimplementedChatServer{},
		log:                     log.NewHelper(logger),
	}
}

// 获取聊天历史记录
func (s *ChatService) GetChatHistory(ctx context.Context, req *chatv1.GetChatHistoryRequest) (*chatv1.GetChatHistoryReply, error) {
	s.log.Infof("获取聊天历史记录: UserId=%s, TargetId=%s", req.UserId, req.TargetId)
	
	// 简化实现 - 返回空的历史记录
	return &chatv1.GetChatHistoryReply{
		Code:     0,
		Message:  "success",
		Messages: []*chatv1.ChatMessageInfo{},
		Total:    0,
	}, nil
}

// 保存聊天消息
func (s *ChatService) SaveChatMessage(ctx context.Context, req *chatv1.SaveChatMessageRequest) (*chatv1.SaveChatMessageReply, error) {
	s.log.Infof("保存聊天消息: FromId=%s, ToId=%s, Content=%s", req.FromId, req.ToId, req.Content)
	
	// 简化实现 - 仅记录日志
	return &chatv1.SaveChatMessageReply{
		Code:    0,
		Message: "success",
	}, nil
}

// 获取用户的聊天房间列表
func (s *ChatService) GetUserChatRooms(ctx context.Context, req *chatv1.GetUserChatRoomsRequest) (*chatv1.GetUserChatRoomsReply, error) {
	s.log.Infof("获取用户聊天房间列表: UserId=%s", req.UserId)
	
	// 简化实现 - 返回空房间列表
	return &chatv1.GetUserChatRoomsReply{
		Code:    0,
		Message: "success",
		Rooms:   []*chatv1.ChatRoomInfo{},
	}, nil
}

// MarkMessageAsRead 方法已移除，因为protobuf中没有定义相关类型

// 生成房间ID
func generateRoomID(user1ID, user2ID string) string {
	if user1ID < user2ID {
		return user1ID + "_" + user2ID
	}
	return user2ID + "_" + user1ID
}

package service

//
//import (
//	"context"
//	"strconv"
//	"time"
//
//	chatv1 "kratos_client/api/chat/v1"
//	"kratos_client/internal/biz"
//	"kratos_client/internal/data"
//)
//
//type ChatService struct {
//	chatv1.UnimplementedChatServer
//	data *data.Data
//	uc   *biz.ChatService
//}
//
//func NewChatService(uc *biz.ChatService, d *data.Data) *ChatService {
//	return &ChatService{
//		UnimplementedChatServer: chatv1.UnimplementedChatServer{},
//		data:                    d,
//		uc:                      uc,
//	}
//}
//
//// 获取聊天历史记录
//func (s *ChatService) GetChatHistory(ctx context.Context, req *chatv1.GetChatHistoryRequest) (*chatv1.GetChatHistoryReply, error) {
//	// 设置默认分页参数
//	page := req.Page
//	if page <= 0 {
//		page = 1
//	}
//	pageSize := req.PageSize
//	if pageSize <= 0 {
//		pageSize = 20
//	}
//
//	messages, total, err := s.uc.GetChatHistory(ctx, req.UserId, req.TargetId, page, pageSize)
//	if err != nil {
//		return nil, err
//	}
//
//	// 转换为响应格式
//	var messageInfos []*chatv1.ChatMessageInfo
//	for _, msg := range messages {
//		messageInfos = append(messageInfos, &chatv1.ChatMessageInfo{
//			Id:          msg.ID,
//			FromId:      msg.FromID,
//			ToId:        msg.ToID,
//			FromName:    "", // 这里需要根据实际情况获取用户名
//			ToName:      "", // 这里需要根据实际情况获取用户名
//			Content:     msg.Content,
//			MessageType: msg.MessageType,
//			RoomId:      msg.RoomID,
//			CreatedAt:   msg.CreatedAt.Format("2006-01-02 15:04:05"),
//		})
//	}
//
//	return &chatv1.GetChatHistoryReply{
//		Code:     0,
//		Message:  "success",
//		Messages: messageInfos,
//		Total:    total,
//	}, nil
//}
//
//// 保存聊天消息
//func (s *ChatService) SaveChatMessage(ctx context.Context, req *chatv1.SaveChatMessageRequest) (*chatv1.SaveChatMessageReply, error) {
//	// 生成房间ID（如果没有提供）
//	roomID := req.RoomId
//	if roomID == "" {
//		roomID = generateRoomID(req.FromId, req.ToId)
//	}
//
//	message := &biz.MtChatMessage{
//		FromID:      req.FromId,
//		ToID:        req.ToId,
//		Content:     req.Content,
//		MessageType: req.MessageType,
//		RoomID:      roomID,
//		CreatedAt:   time.Now(),
//	}
//
//	err := s.uc.SaveChatMessage(ctx, message)
//	if err != nil {
//		return nil, err
//	}
//
//	return &chatv1.SaveChatMessageReply{
//		Code:    0,
//		Message: "success",
//	}, nil
//}
//
//// 获取用户的聊天房间列表
//func (s *ChatService) GetUserChatRooms(ctx context.Context, req *chatv1.GetUserChatRoomsRequest) (*chatv1.GetUserChatRoomsReply, error) {
//	rooms, err := s.uc.GetUserChatRooms(ctx, req.UserId)
//	if err != nil {
//		return nil, err
//	}
//
//	var roomInfos []*chatv1.ChatRoomInfo
//	for _, room := range rooms {
//		// 确定目标用户ID
//		targetID := room.User1ID
//		if room.User1ID == req.UserId {
//			targetID = room.User2ID
//		}
//
//		// 获取未读消息数量
//		unreadCount, _ := s.uc.GetUnreadCount(ctx, room.RoomID, req.UserId)
//
//		roomInfos = append(roomInfos, &chatv1.ChatRoomInfo{
//			RoomId:          room.RoomID,
//			TargetId:        targetID,
//			TargetName:      "", // 这里需要根据实际情况获取用户名
//			TargetRole:      "", // 这里需要根据实际情况获取用户角色
//			LastMessage:     room.LastMessage,
//			LastMessageTime: room.LastMessageTime.Format("2006-01-02 15:04:05"),
//			UnreadCount:     unreadCount,
//		})
//	}
//
//	return &chatv1.GetUserChatRoomsReply{
//		Code:    0,
//		Message: "success",
//		Rooms:   roomInfos,
//	}, nil
//}
//
//// 生成房间ID的辅助函数
//func generateRoomID(userID1, userID2 int32) string {
//	if userID1 > userID2 {
//		userID1, userID2 = userID2, userID1
//	}
//	return "room_" + strconv.Itoa(int(userID1)) + "_" + strconv.Itoa(int(userID2))
//}

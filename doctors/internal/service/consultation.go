package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	pb "doctors/api/consultation/v1"
	"doctors/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
)

// WebSocket 升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域
	},
}

// WebSocket 连接管理
type WSConnection struct {
	conn           *websocket.Conn
	userID         uint
	userType       string // "doctor" or "patient"
	consultationID uint
	send           chan []byte
	hub            *WSHub
}

// WebSocket 消息结构
type WSMessage struct {
	Type           string      `json:"type"`
	ConsultationID uint        `json:"consultation_id"`
	SenderID       uint        `json:"sender_id"`
	SenderType     string      `json:"sender_type"`
	MessageType    string      `json:"message_type"`
	Content        string      `json:"content"`
	MediaURL       string      `json:"media_url,omitempty"`
	Timestamp      time.Time   `json:"timestamp"`
	Data           interface{} `json:"data,omitempty"`
}

// WebSocket Hub 管理所有连接
type WSHub struct {
	connections map[uint]*WSConnection // consultationID -> connection
	register    chan *WSConnection
	unregister  chan *WSConnection
	broadcast   chan []byte
	mutex       sync.RWMutex
}

// 创建新的 Hub
func NewWSHub() *WSHub {
	return &WSHub{
		connections: make(map[uint]*WSConnection),
		register:    make(chan *WSConnection),
		unregister:  make(chan *WSConnection),
		broadcast:   make(chan []byte),
	}
}

// 运行 Hub
func (h *WSHub) Run() {
	for {
		select {
		case conn := <-h.register:
			h.mutex.Lock()
			h.connections[conn.consultationID] = conn
			h.mutex.Unlock()

		case conn := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.connections[conn.consultationID]; ok {
				delete(h.connections, conn.consultationID)
				close(conn.send)
			}
			h.mutex.Unlock()

		case message := <-h.broadcast:
			h.mutex.RLock()
			for _, conn := range h.connections {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(h.connections, conn.consultationID)
				}
			}
			h.mutex.RUnlock()
		}
	}
}

// 发送消息到特定问诊
func (h *WSHub) SendToConsultation(consultationID uint, message []byte) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	if conn, ok := h.connections[consultationID]; ok {
		select {
		case conn.send <- message:
		default:
			close(conn.send)
			delete(h.connections, consultationID)
		}
	}
}

type ConsultationService struct {
	pb.UnimplementedConsultationServer
	uc  *biz.ConsultationUsecase
	log *log.Helper
	hub *WSHub
}

func NewConsultationService(uc *biz.ConsultationUsecase, logger log.Logger) *ConsultationService {
	hub := NewWSHub()
	go hub.Run() // 启动 WebSocket Hub

	return &ConsultationService{
		uc:  uc,
		log: log.NewHelper(logger),
		hub: hub,
	}
}

// StartConsultation 开始问诊
func (s *ConsultationService) StartConsultation(ctx context.Context, req *pb.StartConsultationReq) (*pb.StartConsultationResp, error) {
	// 验证必填字段
	if req.PatientId == 0 || req.DoctorId == 0 {
		return &pb.StartConsultationResp{
			Message: "患者ID和医生ID不能为空",
			Code:    400,
		}, nil
	}

	// 创建问诊对象
	consultation := &biz.Consultation{
		PatientID:      uint(req.PatientId),
		DoctorID:       uint(req.DoctorId),
		Type:           req.Type,
		ChiefComplaint: req.ChiefComplaint,
		PresentIllness: req.PresentIllness,
		Symptoms:       req.Symptoms,
		Fee:            req.Fee,
	}

	// 设置默认值
	if consultation.Type == "" {
		consultation.Type = "在线问诊"
	}

	// 调用业务逻辑开始问诊
	err := s.uc.StartConsultation(ctx, consultation)
	if err != nil {
		s.log.WithContext(ctx).Errorf("开始问诊失败: %v", err)
		return &pb.StartConsultationResp{
			Message: "开始问诊失败",
			Code:    500,
		}, nil
	}

	// 发送 WebSocket 通知
	wsMessage := WSMessage{
		Type:           "consultation_started",
		ConsultationID: consultation.ID,
		SenderID:       uint(req.PatientId),
		SenderType:     "patient",
		Timestamp:      time.Now(),
		Data: map[string]interface{}{
			"consultation_code": consultation.ConsultationCode,
			"type":              consultation.Type,
			"status":            consultation.Status,
		},
	}
	s.broadcastWSMessage(consultation.ID, wsMessage)

	return &pb.StartConsultationResp{
		Message:          "问诊开始成功",
		Code:             200,
		ConsultationId:   int64(consultation.ID),
		ConsultationCode: consultation.ConsultationCode,
	}, nil
}

// GetConsultationDetail 获取问诊详情
func (s *ConsultationService) GetConsultationDetail(ctx context.Context, req *pb.GetConsultationDetailReq) (*pb.GetConsultationDetailResp, error) {
	consultation, err := s.uc.GetConsultationByID(ctx, uint(req.ConsultationId))
	if err != nil {
		s.log.WithContext(ctx).Errorf("获取问诊详情失败: %v", err)
		return &pb.GetConsultationDetailResp{
			Message: "问诊不存在",
			Code:    404,
		}, nil
	}

	// 构建响应数据
	detail := &pb.ConsultationDetail{
		ConsultationId:   int64(consultation.ID),
		ConsultationCode: consultation.ConsultationCode,
		PatientId:        int64(consultation.PatientID),
		PatientName:      consultation.PatientName,
		DoctorId:         int64(consultation.DoctorID),
		DoctorName:       consultation.DoctorName,
		Type:             consultation.Type,
		Status:           consultation.Status,
		ChiefComplaint:   consultation.ChiefComplaint,
		PresentIllness:   consultation.PresentIllness,
		Symptoms:         consultation.Symptoms,
		Diagnosis:        consultation.Diagnosis,
		Treatment:        consultation.Treatment,
		Prescription:     consultation.Prescription,
		Advice:           consultation.Advice,
		Duration:         int32(consultation.Duration),
		Fee:              consultation.Fee,
		PaymentStatus:    consultation.PaymentStatus,
		Rating:           int32(consultation.Rating),
		Feedback:         consultation.Feedback,
		Remarks:          consultation.Remarks,
		CreatedAt:        consultation.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:        consultation.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	// 处理时间字段
	if consultation.StartTime != nil {
		detail.StartTime = consultation.StartTime.Format("2006-01-02 15:04:05")
	}
	if consultation.EndTime != nil {
		detail.EndTime = consultation.EndTime.Format("2006-01-02 15:04:05")
	}

	return &pb.GetConsultationDetailResp{
		Message: "获取成功",
		Code:    200,
		Detail:  detail,
	}, nil
}

// UpdateConsultationStatus 更新问诊状态
func (s *ConsultationService) UpdateConsultationStatus(ctx context.Context, req *pb.UpdateConsultationStatusReq) (*pb.UpdateConsultationStatusResp, error) {
	consultation := &biz.Consultation{
		ID:           uint(req.ConsultationId),
		Status:       req.Status,
		Diagnosis:    req.Diagnosis,
		Treatment:    req.Treatment,
		Prescription: req.Prescription,
		Advice:       req.Advice,
	}

	err := s.uc.UpdateConsultationStatus(ctx, consultation)
	if err != nil {
		s.log.WithContext(ctx).Errorf("更新问诊状态失败: %v", err)

		if errors.Is(err, biz.ErrConsultationNotFound) {
			return &pb.UpdateConsultationStatusResp{
				Message: "问诊不存在",
				Code:    404,
			}, nil
		}

		if errors.Is(err, biz.ErrConsultationStatusInvalid) {
			return &pb.UpdateConsultationStatusResp{
				Message: "问诊状态无效",
				Code:    400,
			}, nil
		}

		return &pb.UpdateConsultationStatusResp{
			Message: "更新状态失败",
			Code:    500,
		}, nil
	}

	// 发送 WebSocket 通知
	wsMessage := WSMessage{
		Type:           "status_updated",
		ConsultationID: uint(req.ConsultationId),
		Timestamp:      time.Now(),
		Data: map[string]interface{}{
			"status":       req.Status,
			"diagnosis":    req.Diagnosis,
			"treatment":    req.Treatment,
			"prescription": req.Prescription,
			"advice":       req.Advice,
		},
	}
	s.broadcastWSMessage(uint(req.ConsultationId), wsMessage)

	return &pb.UpdateConsultationStatusResp{
		Message: "状态更新成功",
		Code:    200,
	}, nil
}

// EndConsultation 结束问诊
func (s *ConsultationService) EndConsultation(ctx context.Context, req *pb.EndConsultationReq) (*pb.EndConsultationResp, error) {
	consultation := &biz.Consultation{
		ID:           uint(req.ConsultationId),
		Diagnosis:    req.Diagnosis,
		Treatment:    req.Treatment,
		Prescription: req.Prescription,
		Advice:       req.Advice,
		Rating:       int(req.Rating),
		Feedback:     req.Feedback,
	}

	err := s.uc.EndConsultation(ctx, consultation)
	if err != nil {
		s.log.WithContext(ctx).Errorf("结束问诊失败: %v", err)

		if errors.Is(err, biz.ErrConsultationNotFound) {
			return &pb.EndConsultationResp{
				Message: "问诊不存在",
				Code:    404,
			}, nil
		}

		if errors.Is(err, biz.ErrConsultationAlreadyEnded) {
			return &pb.EndConsultationResp{
				Message: "问诊已结束",
				Code:    400,
			}, nil
		}

		return &pb.EndConsultationResp{
			Message: "结束问诊失败",
			Code:    500,
		}, nil
	}

	// 发送 WebSocket 通知
	wsMessage := WSMessage{
		Type:           "consultation_ended",
		ConsultationID: uint(req.ConsultationId),
		Timestamp:      time.Now(),
		Data: map[string]interface{}{
			"diagnosis":    req.Diagnosis,
			"treatment":    req.Treatment,
			"prescription": req.Prescription,
			"advice":       req.Advice,
			"rating":       req.Rating,
			"feedback":     req.Feedback,
		},
	}
	s.broadcastWSMessage(uint(req.ConsultationId), wsMessage)

	return &pb.EndConsultationResp{
		Message: "问诊结束成功",
		Code:    200,
	}, nil
}

// GetConsultationHistory 获取问诊历史
func (s *ConsultationService) GetConsultationHistory(ctx context.Context, req *pb.GetConsultationHistoryReq) (*pb.GetConsultationHistoryResp, error) {
	query := &biz.ConsultationListQuery{
		PatientID: uint(req.PatientId),
		DoctorID:  uint(req.DoctorId),
		Status:    req.Status,
		Type:      req.Type,
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
	}

	result, err := s.uc.GetConsultationHistory(ctx, query)
	if err != nil {
		s.log.WithContext(ctx).Errorf("获取问诊历史失败: %v", err)
		return &pb.GetConsultationHistoryResp{
			Message: "获取问诊历史失败",
			Code:    500,
		}, nil
	}

	// 转换为响应格式
	consultations := make([]*pb.ConsultationDetail, len(result.Consultations))
	for i, consultation := range result.Consultations {
		detail := &pb.ConsultationDetail{
			ConsultationId:   int64(consultation.ID),
			ConsultationCode: consultation.ConsultationCode,
			PatientId:        int64(consultation.PatientID),
			PatientName:      consultation.PatientName,
			DoctorId:         int64(consultation.DoctorID),
			DoctorName:       consultation.DoctorName,
			Type:             consultation.Type,
			Status:           consultation.Status,
			ChiefComplaint:   consultation.ChiefComplaint,
			PresentIllness:   consultation.PresentIllness,
			Symptoms:         consultation.Symptoms,
			Diagnosis:        consultation.Diagnosis,
			Treatment:        consultation.Treatment,
			Prescription:     consultation.Prescription,
			Advice:           consultation.Advice,
			Duration:         int32(consultation.Duration),
			Fee:              consultation.Fee,
			PaymentStatus:    consultation.PaymentStatus,
			Rating:           int32(consultation.Rating),
			Feedback:         consultation.Feedback,
			Remarks:          consultation.Remarks,
			CreatedAt:        consultation.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        consultation.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		// 处理时间字段
		if consultation.StartTime != nil {
			detail.StartTime = consultation.StartTime.Format("2006-01-02 15:04:05")
		}
		if consultation.EndTime != nil {
			detail.EndTime = consultation.EndTime.Format("2006-01-02 15:04:05")
		}

		consultations[i] = detail
	}

	return &pb.GetConsultationHistoryResp{
		Message:       "获取成功",
		Code:          200,
		Consultations: consultations,
		Total:         result.Total,
		Page:          int32(result.Page),
		PageSize:      int32(result.PageSize),
	}, nil
}

// SendMessage 发送消息
func (s *ConsultationService) SendMessage(ctx context.Context, req *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	message := &biz.ConsultationMessage{
		ConsultationID: uint(req.ConsultationId),
		SenderID:       uint(req.SenderId),
		SenderType:     req.SenderType,
		MessageType:    req.MessageType,
		Content:        req.Content,
		MediaURL:       req.MediaURL,
	}

	// 设置默认消息类型
	if message.MessageType == "" {
		message.MessageType = "text"
	}

	err := s.uc.SendMessage(ctx, message)
	if err != nil {
		s.log.WithContext(ctx).Errorf("发送消息失败: %v", err)

		if errors.Is(err, biz.ErrConsultationNotFound) {
			return &pb.SendMessageResp{
				Message: "问诊不存在",
				Code:    404,
			}, nil
		}

		return &pb.SendMessageResp{
			Message: "发送消息失败",
			Code:    500,
		}, nil
	}

	// 发送 WebSocket 消息
	wsMessage := WSMessage{
		Type:           "new_message",
		ConsultationID: uint(req.ConsultationId),
		SenderID:       uint(req.SenderId),
		SenderType:     req.SenderType,
		MessageType:    req.MessageType,
		Content:        req.Content,
		MediaURL:       req.MediaURL,
		Timestamp:      time.Now(),
	}
	s.broadcastWSMessage(uint(req.ConsultationId), wsMessage)

	return &pb.SendMessageResp{
		Message:   "发送成功",
		Code:      200,
		MessageId: int64(message.ID),
	}, nil
}

// GetMessages 获取消息列表
func (s *ConsultationService) GetMessages(ctx context.Context, req *pb.GetMessagesReq) (*pb.GetMessagesResp, error) {
	query := &biz.MessageListQuery{
		ConsultationID: uint(req.ConsultationId),
		Page:           int(req.Page),
		PageSize:       int(req.PageSize),
	}

	result, err := s.uc.GetMessages(ctx, query)
	if err != nil {
		s.log.WithContext(ctx).Errorf("获取消息列表失败: %v", err)
		return &pb.GetMessagesResp{
			Message: "获取消息列表失败",
			Code:    500,
		}, nil
	}

	// 转换为响应格式
	messages := make([]*pb.ConsultationMessage, len(result.Messages))
	for i, message := range result.Messages {
		pbMessage := &pb.ConsultationMessage{
			MessageId:      int64(message.ID),
			ConsultationId: int64(message.ConsultationID),
			SenderId:       int64(message.SenderID),
			SenderType:     message.SenderType,
			MessageType:    message.MessageType,
			Content:        message.Content,
			MediaURL:       message.MediaURL,
			IsRead:         message.IsRead,
			CreatedAt:      message.CreatedAt.Format("2006-01-02 15:04:05"),
		}

		// 处理已读时间
		if message.ReadTime != nil {
			pbMessage.ReadTime = message.ReadTime.Format("2006-01-02 15:04:05")
		}

		messages[i] = pbMessage
	}

	return &pb.GetMessagesResp{
		Message:  "获取成功",
		Code:     200,
		Messages: messages,
		Total:    result.Total,
	}, nil
}

// MarkMessageRead 标记消息已读
func (s *ConsultationService) MarkMessageRead(ctx context.Context, req *pb.MarkMessageReadReq) (*pb.MarkMessageReadResp, error) {
	err := s.uc.MarkMessageAsRead(ctx, uint(req.MessageId), uint(req.ReaderId))
	if err != nil {
		s.log.WithContext(ctx).Errorf("标记消息已读失败: %v", err)
		return &pb.MarkMessageReadResp{
			Message: "标记已读失败",
			Code:    500,
		}, nil
	}

	return &pb.MarkMessageReadResp{
		Message: "标记成功",
		Code:    200,
	}, nil
}

// AddConsultationRecord 添加问诊记录
func (s *ConsultationService) AddConsultationRecord(ctx context.Context, req *pb.AddConsultationRecordReq) (*pb.AddConsultationRecordResp, error) {
	record := &biz.ConsultationRecord{
		ConsultationID: uint(req.ConsultationId),
		RecordType:     req.RecordType,
		Title:          req.Title,
		Content:        req.Content,
		DataValue:      req.DataValue,
		FileURL:        req.FileURL,
		CreatedBy:      uint(req.CreatedBy),
	}

	err := s.uc.AddConsultationRecord(ctx, record)
	if err != nil {
		s.log.WithContext(ctx).Errorf("添加问诊记录失败: %v", err)

		if errors.Is(err, biz.ErrConsultationNotFound) {
			return &pb.AddConsultationRecordResp{
				Message: "问诊不存在",
				Code:    404,
			}, nil
		}

		return &pb.AddConsultationRecordResp{
			Message: "添加记录失败",
			Code:    500,
		}, nil
	}

	return &pb.AddConsultationRecordResp{
		Message:  "添加成功",
		Code:     200,
		RecordId: int64(record.ID),
	}, nil
}

// GetConsultationRecords 获取问诊记录
func (s *ConsultationService) GetConsultationRecords(ctx context.Context, req *pb.GetConsultationRecordsReq) (*pb.GetConsultationRecordsResp, error) {
	records, err := s.uc.GetConsultationRecords(ctx, uint(req.ConsultationId), req.RecordType)
	if err != nil {
		s.log.WithContext(ctx).Errorf("获取问诊记录失败: %v", err)
		return &pb.GetConsultationRecordsResp{
			Message: "获取记录失败",
			Code:    500,
		}, nil
	}

	// 转换为响应格式
	pbRecords := make([]*pb.ConsultationRecord, len(records))
	for i, record := range records {
		pbRecords[i] = &pb.ConsultationRecord{
			RecordId:       int64(record.ID),
			ConsultationId: int64(record.ConsultationID),
			RecordType:     record.RecordType,
			Title:          record.Title,
			Content:        record.Content,
			DataValue:      record.DataValue,
			FileURL:        record.FileURL,
			CreatedBy:      int64(record.CreatedBy),
			CreatedAt:      record.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:      record.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return &pb.GetConsultationRecordsResp{
		Message: "获取成功",
		Code:    200,
		Records: pbRecords,
	}, nil
}

// GetConsultationReport 获取问诊报告
func (s *ConsultationService) GetConsultationReport(ctx context.Context, req *pb.GetConsultationReportReq) (*pb.GetConsultationReportResp, error) {
	consultation, records, messages, err := s.uc.GetConsultationReport(ctx, uint(req.ConsultationId))
	if err != nil {
		s.log.WithContext(ctx).Errorf("获取问诊报告失败: %v", err)

		if errors.Is(err, biz.ErrConsultationNotFound) {
			return &pb.GetConsultationReportResp{
				Message: "问诊不存在",
				Code:    404,
			}, nil
		}

		return &pb.GetConsultationReportResp{
			Message: "获取报告失败",
			Code:    500,
		}, nil
	}

	// 构建问诊详情
	detail := &pb.ConsultationDetail{
		ConsultationId:   int64(consultation.ID),
		ConsultationCode: consultation.ConsultationCode,
		PatientId:        int64(consultation.PatientID),
		PatientName:      consultation.PatientName,
		DoctorId:         int64(consultation.DoctorID),
		DoctorName:       consultation.DoctorName,
		Type:             consultation.Type,
		Status:           consultation.Status,
		ChiefComplaint:   consultation.ChiefComplaint,
		PresentIllness:   consultation.PresentIllness,
		Symptoms:         consultation.Symptoms,
		Diagnosis:        consultation.Diagnosis,
		Treatment:        consultation.Treatment,
		Prescription:     consultation.Prescription,
		Advice:           consultation.Advice,
		Duration:         int32(consultation.Duration),
		Fee:              consultation.Fee,
		PaymentStatus:    consultation.PaymentStatus,
		Rating:           int32(consultation.Rating),
		Feedback:         consultation.Feedback,
		Remarks:          consultation.Remarks,
		CreatedAt:        consultation.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:        consultation.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	// 处理时间字段
	if consultation.StartTime != nil {
		detail.StartTime = consultation.StartTime.Format("2006-01-02 15:04:05")
	}
	if consultation.EndTime != nil {
		detail.EndTime = consultation.EndTime.Format("2006-01-02 15:04:05")
	}

	// 构建记录列表
	pbRecords := make([]*pb.ConsultationRecord, len(records))
	for i, record := range records {
		pbRecords[i] = &pb.ConsultationRecord{
			RecordId:       int64(record.ID),
			ConsultationId: int64(record.ConsultationID),
			RecordType:     record.RecordType,
			Title:          record.Title,
			Content:        record.Content,
			DataValue:      record.DataValue,
			FileURL:        record.FileURL,
			CreatedBy:      int64(record.CreatedBy),
			CreatedAt:      record.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:      record.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	// 构建消息列表
	pbMessages := make([]*pb.ConsultationMessage, len(messages))
	for i, message := range messages {
		pbMessage := &pb.ConsultationMessage{
			MessageId:      int64(message.ID),
			ConsultationId: int64(message.ConsultationID),
			SenderId:       int64(message.SenderID),
			SenderType:     message.SenderType,
			MessageType:    message.MessageType,
			Content:        message.Content,
			MediaURL:       message.MediaURL,
			IsRead:         message.IsRead,
			CreatedAt:      message.CreatedAt.Format("2006-01-02 15:04:05"),
		}

		if message.ReadTime != nil {
			pbMessage.ReadTime = message.ReadTime.Format("2006-01-02 15:04:05")
		}

		pbMessages[i] = pbMessage
	}

	// 构建报告
	report := &pb.ConsultationReport{
		Detail:   detail,
		Records:  pbRecords,
		Messages: pbMessages,
	}

	return &pb.GetConsultationReportResp{
		Message: "获取成功",
		Code:    200,
		Report:  report,
	}, nil
}

// GetConsultationsByType 按类型获取问诊
func (s *ConsultationService) GetConsultationsByType(ctx context.Context, req *pb.GetConsultationsByTypeReq) (*pb.GetConsultationsByTypeResp, error) {
	result, err := s.uc.GetConsultationsByType(ctx, req.Type, int(req.Page), int(req.PageSize), req.Status)
	if err != nil {
		s.log.WithContext(ctx).Errorf("按类型获取问诊失败: %v", err)
		return &pb.GetConsultationsByTypeResp{
			Message: "获取问诊失败",
			Code:    500,
		}, nil
	}

	// 转换为响应格式
	consultations := make([]*pb.ConsultationDetail, len(result.Consultations))
	for i, consultation := range result.Consultations {
		detail := &pb.ConsultationDetail{
			ConsultationId:   int64(consultation.ID),
			ConsultationCode: consultation.ConsultationCode,
			PatientId:        int64(consultation.PatientID),
			PatientName:      consultation.PatientName,
			DoctorId:         int64(consultation.DoctorID),
			DoctorName:       consultation.DoctorName,
			Type:             consultation.Type,
			Status:           consultation.Status,
			ChiefComplaint:   consultation.ChiefComplaint,
			PresentIllness:   consultation.PresentIllness,
			Symptoms:         consultation.Symptoms,
			Diagnosis:        consultation.Diagnosis,
			Treatment:        consultation.Treatment,
			Prescription:     consultation.Prescription,
			Advice:           consultation.Advice,
			Duration:         int32(consultation.Duration),
			Fee:              consultation.Fee,
			PaymentStatus:    consultation.PaymentStatus,
			Rating:           int32(consultation.Rating),
			Feedback:         consultation.Feedback,
			Remarks:          consultation.Remarks,
			CreatedAt:        consultation.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        consultation.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		// 处理时间字段
		if consultation.StartTime != nil {
			detail.StartTime = consultation.StartTime.Format("2006-01-02 15:04:05")
		}
		if consultation.EndTime != nil {
			detail.EndTime = consultation.EndTime.Format("2006-01-02 15:04:05")
		}

		consultations[i] = detail
	}

	return &pb.GetConsultationsByTypeResp{
		Message:       "获取成功",
		Code:          200,
		Consultations: consultations,
		Total:         result.Total,
	}, nil
}

// WebSocket 处理函数
func (s *ConsultationService) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接为 WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.log.Errorf("WebSocket 升级失败: %v", err)
		return
	}
	defer conn.Close()

	// 获取连接参数
	consultationIDStr := r.URL.Query().Get("consultation_id")
	userIDStr := r.URL.Query().Get("user_id")
	userType := r.URL.Query().Get("user_type")

	consultationID, err := strconv.ParseUint(consultationIDStr, 10, 32)
	if err != nil {
		s.log.Errorf("无效的问诊ID: %v", err)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		s.log.Errorf("无效的用户ID: %v", err)
		return
	}

	// 创建 WebSocket 连接
	wsConn := &WSConnection{
		conn:           conn,
		userID:         uint(userID),
		userType:       userType,
		consultationID: uint(consultationID),
		send:           make(chan []byte, 256),
		hub:            s.hub,
	}

	// 注册连接
	s.hub.register <- wsConn

	// 启动读写协程
	go wsConn.writePump()
	go wsConn.readPump(s)
}

// WebSocket 读取消息
func (c *WSConnection) readPump(service *ConsultationService) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var wsMsg WSMessage
		if err := json.Unmarshal(message, &wsMsg); err != nil {
			continue
		}

		// 处理不同类型的消息
		switch wsMsg.Type {
		case "send_message":
			// 保存消息到数据库
			bizMessage := &biz.ConsultationMessage{
				ConsultationID: c.consultationID,
				SenderID:       c.userID,
				SenderType:     c.userType,
				MessageType:    wsMsg.MessageType,
				Content:        wsMsg.Content,
				MediaURL:       wsMsg.MediaURL,
			}

			ctx := context.Background()
			err := service.uc.SendMessage(ctx, bizMessage)
			if err != nil {
				service.log.Errorf("保存WebSocket消息失败: %v", err)
				continue
			}

			// 广播消息
			wsMsg.SenderID = c.userID
			wsMsg.SenderType = c.userType
			wsMsg.ConsultationID = c.consultationID
			wsMsg.Timestamp = time.Now()
			service.broadcastWSMessage(c.consultationID, wsMsg)

		case "typing":
			// 广播正在输入状态
			wsMsg.SenderID = c.userID
			wsMsg.SenderType = c.userType
			wsMsg.ConsultationID = c.consultationID
			wsMsg.Timestamp = time.Now()
			service.broadcastWSMessage(c.consultationID, wsMsg)
		}
	}
}

// WebSocket 写入消息
func (c *WSConnection) writePump() {
	defer c.conn.Close()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

// 广播 WebSocket 消息
func (s *ConsultationService) broadcastWSMessage(consultationID uint, message WSMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		s.log.Errorf("序列化WebSocket消息失败: %v", err)
		return
	}

	s.hub.SendToConsultation(consultationID, data)
}

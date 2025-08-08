package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrConsultationNotFound 问诊不存在
	ErrConsultationNotFound = errors.NotFound("CONSULTATION_NOT_FOUND", "问诊不存在")
	// ErrConsultationStatusInvalid 问诊状态无效
	ErrConsultationStatusInvalid = errors.BadRequest("CONSULTATION_STATUS_INVALID", "问诊状态无效")
	// ErrConsultationAlreadyEnded 问诊已结束
	ErrConsultationAlreadyEnded = errors.BadRequest("CONSULTATION_ALREADY_ENDED", "问诊已结束")
	// ErrMessageNotFound 消息不存在
	ErrMessageNotFound = errors.NotFound("MESSAGE_NOT_FOUND", "消息不存在")
	// ErrRecordNotFound 记录不存在
	ErrRecordNotFound = errors.NotFound("RECORD_NOT_FOUND", "记录不存在")
)

// Consultation 问诊业务实体
type Consultation struct {
	ID               uint       `json:"id"`
	ConsultationCode string     `json:"consultation_code"`
	PatientID        uint       `json:"patient_id"`
	DoctorID         uint       `json:"doctor_id"`
	Type             string     `json:"type"`
	Status           string     `json:"status"`
	ChiefComplaint   string     `json:"chief_complaint"`
	PresentIllness   string     `json:"present_illness"`
	Symptoms         string     `json:"symptoms"`
	Diagnosis        string     `json:"diagnosis"`
	Treatment        string     `json:"treatment"`
	Prescription     string     `json:"prescription"`
	Advice           string     `json:"advice"`
	StartTime        *time.Time `json:"start_time"`
	EndTime          *time.Time `json:"end_time"`
	Duration         int        `json:"duration"`
	Fee              float64    `json:"fee"`
	PaymentStatus    string     `json:"payment_status"`
	Rating           int        `json:"rating"`
	Feedback         string     `json:"feedback"`
	Remarks          string     `json:"remarks"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`

	// 关联信息
	PatientName string `json:"patient_name,omitempty"`
	DoctorName  string `json:"doctor_name,omitempty"`
}

// ConsultationMessage 问诊消息业务实体
type ConsultationMessage struct {
	ID             uint       `json:"id"`
	ConsultationID uint       `json:"consultation_id"`
	SenderID       uint       `json:"sender_id"`
	SenderType     string     `json:"sender_type"`
	MessageType    string     `json:"message_type"`
	Content        string     `json:"content"`
	MediaURL       string     `json:"media_url"`
	IsRead         bool       `json:"is_read"`
	ReadTime       *time.Time `json:"read_time"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// ConsultationRecord 问诊记录业务实体
type ConsultationRecord struct {
	ID             uint      `json:"id"`
	ConsultationID uint      `json:"consultation_id"`
	RecordType     string    `json:"record_type"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	DataValue      string    `json:"data_value"`
	FileURL        string    `json:"file_url"`
	CreatedBy      uint      `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ConsultationListQuery 问诊列表查询参数
type ConsultationListQuery struct {
	PatientID uint   `json:"patient_id"`
	DoctorID  uint   `json:"doctor_id"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}

// ConsultationListResult 问诊列表查询结果
type ConsultationListResult struct {
	Consultations []*Consultation `json:"consultations"`
	Total         int64           `json:"total"`
	Page          int             `json:"page"`
	PageSize      int             `json:"page_size"`
}

// MessageListQuery 消息列表查询参数
type MessageListQuery struct {
	ConsultationID uint `json:"consultation_id"`
	Page           int  `json:"page"`
	PageSize       int  `json:"page_size"`
}

// MessageListResult 消息列表查询结果
type MessageListResult struct {
	Messages []*ConsultationMessage `json:"messages"`
	Total    int64                  `json:"total"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"page_size"`
}

// ConsultationRepo 问诊数据访问接口
type ConsultationRepo interface {
	// 问诊基础操作
	CreateConsultation(ctx context.Context, consultation *Consultation) error
	GetConsultationByID(ctx context.Context, id uint) (*Consultation, error)
	GetConsultationByCode(ctx context.Context, code string) (*Consultation, error)
	UpdateConsultation(ctx context.Context, consultation *Consultation) error
	DeleteConsultation(ctx context.Context, id uint) error

	// 问诊查询操作
	GetConsultationList(ctx context.Context, query *ConsultationListQuery) (*ConsultationListResult, error)
	GetConsultationsByType(ctx context.Context, consultationType string, page, pageSize int, status string) (*ConsultationListResult, error)

	// 消息操作
	CreateMessage(ctx context.Context, message *ConsultationMessage) error
	GetMessageByID(ctx context.Context, id uint) (*ConsultationMessage, error)
	GetMessageList(ctx context.Context, query *MessageListQuery) (*MessageListResult, error)
	UpdateMessage(ctx context.Context, message *ConsultationMessage) error
	MarkMessageAsRead(ctx context.Context, messageID uint, readerID uint) error

	// 记录操作
	CreateRecord(ctx context.Context, record *ConsultationRecord) error
	GetRecordsByConsultationID(ctx context.Context, consultationID uint, recordType string) ([]*ConsultationRecord, error)
	GetRecordByID(ctx context.Context, id uint) (*ConsultationRecord, error)
	UpdateRecord(ctx context.Context, record *ConsultationRecord) error

	// 新增：获取医生信息
	GetDoctorByID(ctx context.Context, id uint) (*Doctor, error)
}

// ConsultationUsecase 问诊业务逻辑
type ConsultationUsecase struct {
	repo ConsultationRepo
	log  *log.Helper
}

// NewConsultationUsecase 创建问诊业务逻辑实例
func NewConsultationUsecase(repo ConsultationRepo, logger log.Logger) *ConsultationUsecase {
	return &ConsultationUsecase{repo: repo, log: log.NewHelper(logger)}
}

// StartConsultation 开始问诊
func (uc *ConsultationUsecase) StartConsultation(ctx context.Context, consultation *Consultation) error {
	// 生成问诊编码
	consultation.ConsultationCode = uc.generateConsultationCode()

	// 设置初始状态
	consultation.Status = "待接诊"
	consultation.PaymentStatus = "未支付"

	// 设置开始时间
	now := time.Now()
	consultation.StartTime = &now

	return uc.repo.CreateConsultation(ctx, consultation)
}

// GetConsultationByID 根据ID获取问诊详情
func (uc *ConsultationUsecase) GetConsultationByID(ctx context.Context, id uint) (*Consultation, error) {
	return uc.repo.GetConsultationByID(ctx, id)
}

// UpdateConsultationStatus 更新问诊状态
func (uc *ConsultationUsecase) UpdateConsultationStatus(ctx context.Context, consultation *Consultation) error {
	// 检查问诊是否存在
	existingConsultation, err := uc.repo.GetConsultationByID(ctx, consultation.ID)
	if err != nil {
		return ErrConsultationNotFound
	}

	// 检查状态是否有效
	if !uc.isValidStatus(consultation.Status) {
		return ErrConsultationStatusInvalid
	}

	// 如果状态变为进行中，更新开始时间
	if consultation.Status == "进行中" && existingConsultation.Status == "待接诊" {
		now := time.Now()
		consultation.StartTime = &now
	}

	// 更新允许修改的字段
	existingConsultation.Status = consultation.Status
	existingConsultation.Diagnosis = consultation.Diagnosis
	existingConsultation.Treatment = consultation.Treatment
	existingConsultation.Prescription = consultation.Prescription
	existingConsultation.Advice = consultation.Advice
	if consultation.StartTime != nil {
		existingConsultation.StartTime = consultation.StartTime
	}

	return uc.repo.UpdateConsultation(ctx, existingConsultation)
}

// EndConsultation 结束问诊
func (uc *ConsultationUsecase) EndConsultation(ctx context.Context, consultation *Consultation) error {
	// 检查问诊是否存在
	existingConsultation, err := uc.repo.GetConsultationByID(ctx, consultation.ID)
	if err != nil {
		return ErrConsultationNotFound
	}

	// 检查问诊是否已结束
	if existingConsultation.Status == "已完成" || existingConsultation.Status == "已取消" {
		return ErrConsultationAlreadyEnded
	}

	// 设置结束时间和状态
	now := time.Now()
	existingConsultation.EndTime = &now
	existingConsultation.Status = "已完成"
	existingConsultation.Diagnosis = consultation.Diagnosis
	existingConsultation.Treatment = consultation.Treatment
	existingConsultation.Prescription = consultation.Prescription
	existingConsultation.Advice = consultation.Advice
	existingConsultation.Rating = consultation.Rating
	existingConsultation.Feedback = consultation.Feedback

	// 计算问诊时长
	if existingConsultation.StartTime != nil {
		duration := now.Sub(*existingConsultation.StartTime)
		existingConsultation.Duration = int(duration.Minutes())
	}

	return uc.repo.UpdateConsultation(ctx, existingConsultation)
}

// GetConsultationHistory 获取问诊历史
func (uc *ConsultationUsecase) GetConsultationHistory(ctx context.Context, query *ConsultationListQuery) (*ConsultationListResult, error) {
	// 设置默认分页参数
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 20
	}
	if query.PageSize > 100 {
		query.PageSize = 100
	}

	return uc.repo.GetConsultationList(ctx, query)
}

// GetConsultationsByType 按类型获取问诊
func (uc *ConsultationUsecase) GetConsultationsByType(ctx context.Context, consultationType string, page, pageSize int, status string) (*ConsultationListResult, error) {
	// 设置默认分页参数
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}

	return uc.repo.GetConsultationsByType(ctx, consultationType, page, pageSize, status)
}

// SendMessage 发送消息
func (uc *ConsultationUsecase) SendMessage(ctx context.Context, message *ConsultationMessage) error {
	// 检查问诊是否存在
	consultation, err := uc.repo.GetConsultationByID(ctx, message.ConsultationID)
	if err != nil {
		return ErrConsultationNotFound
	}

	// 新增：校验医生认证状态，患者发送消息时医生未认证则禁止
	if message.SenderType == "patient" {
		doctor, err := uc.repo.GetDoctorByID(ctx, consultation.DoctorID)
		if err != nil || doctor == nil {
			return errors.BadRequest("DOCTOR_NOT_FOUND", "医生不存在")
		}
		if doctor.Status != "启用" && doctor.Status != "已认证" {
			return errors.Forbidden("DOCTOR_NOT_AUTHENTICATED", "医生未认证，无法发送消息")
		}
	}

	return uc.repo.CreateMessage(ctx, message)
}

// GetMessages 获取消息列表
func (uc *ConsultationUsecase) GetMessages(ctx context.Context, query *MessageListQuery) (*MessageListResult, error) {
	// 设置默认分页参数
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 50
	}
	if query.PageSize > 200 {
		query.PageSize = 200
	}

	return uc.repo.GetMessageList(ctx, query)
}

// MarkMessageAsRead 标记消息已读
func (uc *ConsultationUsecase) MarkMessageAsRead(ctx context.Context, messageID uint, readerID uint) error {
	return uc.repo.MarkMessageAsRead(ctx, messageID, readerID)
}

// AddConsultationRecord 添加问诊记录
func (uc *ConsultationUsecase) AddConsultationRecord(ctx context.Context, record *ConsultationRecord) error {
	// 检查问诊是否存在
	_, err := uc.repo.GetConsultationByID(ctx, record.ConsultationID)
	if err != nil {
		return ErrConsultationNotFound
	}

	return uc.repo.CreateRecord(ctx, record)
}

// GetConsultationRecords 获取问诊记录
func (uc *ConsultationUsecase) GetConsultationRecords(ctx context.Context, consultationID uint, recordType string) ([]*ConsultationRecord, error) {
	return uc.repo.GetRecordsByConsultationID(ctx, consultationID, recordType)
}

// GetConsultationReport 获取问诊报告
func (uc *ConsultationUsecase) GetConsultationReport(ctx context.Context, consultationID uint) (*Consultation, []*ConsultationRecord, []*ConsultationMessage, error) {
	// 获取问诊详情
	consultation, err := uc.repo.GetConsultationByID(ctx, consultationID)
	if err != nil {
		return nil, nil, nil, ErrConsultationNotFound
	}

	// 获取问诊记录
	records, err := uc.repo.GetRecordsByConsultationID(ctx, consultationID, "")
	if err != nil {
		records = []*ConsultationRecord{}
	}

	// 获取消息记录
	messageQuery := &MessageListQuery{
		ConsultationID: consultationID,
		Page:           1,
		PageSize:       1000, // 获取所有消息
	}
	messageResult, err := uc.repo.GetMessageList(ctx, messageQuery)
	messages := []*ConsultationMessage{}
	if err == nil {
		messages = messageResult.Messages
	}

	return consultation, records, messages, nil
}

// generateConsultationCode 生成问诊编码
func (uc *ConsultationUsecase) generateConsultationCode() string {
	now := time.Now()
	return fmt.Sprintf("C%s%06d", now.Format("20060102"), now.Unix()%1000000)
}

// isValidStatus 检查状态是否有效
func (uc *ConsultationUsecase) isValidStatus(status string) bool {
	validStatuses := []string{"待接诊", "进行中", "已完成", "已取消"}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}

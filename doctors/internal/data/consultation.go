package data

import (
	"context"
	"fmt"
	"time"

	"doctors/internal/biz"
	"doctors/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// ConsultationData 问诊数据访问实现
type ConsultationData struct {
	data   *Data
	logger *log.Helper
}

// NewConsultationData 创建问诊数据访问实例
func NewConsultationData(data *Data, logger log.Logger) biz.ConsultationRepo {
	return &ConsultationData{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

// CreateConsultation 创建问诊记录
func (d *ConsultationData) CreateConsultation(ctx context.Context, consultation *biz.Consultation) error {
	consultationModel := &model.Consultations{
		ConsultationCode: consultation.ConsultationCode,
		PatientId:        uint64(consultation.PatientID),
		DoctorId:         uint64(consultation.DoctorID),
		Type:             consultation.Type,
		Status:           consultation.Status,
		ChiefComplaint:   consultation.ChiefComplaint,
		PresentIllness:   consultation.PresentIllness,
		Symptoms:         consultation.Symptoms,
		Fee:              consultation.Fee,
		PaymentStatus:    consultation.PaymentStatus,
	}

	if consultation.StartTime != nil {
		consultationModel.StartTime = *consultation.StartTime
	}

	err := d.data.db.WithContext(ctx).Create(consultationModel).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("创建问诊记录失败: code=%s, error=%v", consultation.ConsultationCode, err)
		return fmt.Errorf("创建问诊记录失败: %w", err)
	}

	consultation.ID = uint(consultationModel.Id)
	consultation.CreatedAt = consultationModel.CreatedAt
	consultation.UpdatedAt = consultationModel.UpdatedAt

	d.logger.WithContext(ctx).Infof("问诊记录创建成功: id=%d, code=%s", consultationModel.Id, consultation.ConsultationCode)
	return nil
}

// GetConsultationByID 根据ID获取问诊信息
func (d *ConsultationData) GetConsultationByID(ctx context.Context, id uint) (*biz.Consultation, error) {
	var consultationModel model.Consultations

	err := d.data.db.WithContext(ctx).Where("id = ?", id).First(&consultationModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("问诊不存在")
		}
		d.logger.WithContext(ctx).Errorf("查询问诊失败: id=%d, error=%v", id, err)
		return nil, fmt.Errorf("查询问诊失败: %w", err)
	}

	consultation := d.modelToEntity(&consultationModel)
	return consultation, nil
}

// GetConsultationByCode 根据编码获取问诊信息
func (d *ConsultationData) GetConsultationByCode(ctx context.Context, code string) (*biz.Consultation, error) {
	var consultationModel model.Consultations

	err := d.data.db.WithContext(ctx).Where("consultation_code = ?", code).First(&consultationModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("问诊不存在")
		}
		d.logger.WithContext(ctx).Errorf("查询问诊失败: code=%s, error=%v", code, err)
		return nil, fmt.Errorf("查询问诊失败: %w", err)
	}

	consultation := d.modelToEntity(&consultationModel)
	return consultation, nil
}

// UpdateConsultation 更新问诊信息
func (d *ConsultationData) UpdateConsultation(ctx context.Context, consultation *biz.Consultation) error {
	consultationModel := &model.Consultations{
		Id:               uint64(consultation.ID),
		ConsultationCode: consultation.ConsultationCode,
		PatientId:        uint64(consultation.PatientID),
		DoctorId:         uint64(consultation.DoctorID),
		Type:             consultation.Type,
		Status:           consultation.Status,
		ChiefComplaint:   consultation.ChiefComplaint,
		PresentIllness:   consultation.PresentIllness,
		Symptoms:         consultation.Symptoms,
		Diagnosis:        consultation.Diagnosis,
		Treatment:        consultation.Treatment,
		Prescription:     consultation.Prescription,
		Advice:           consultation.Advice,
		Duration:         consultation.Duration,
		Fee:              consultation.Fee,
		PaymentStatus:    consultation.PaymentStatus,
		Rating:           consultation.Rating,
		Feedback:         consultation.Feedback,
		Remarks:          consultation.Remarks,
		UpdatedAt:        time.Now(),
	}

	if consultation.StartTime != nil {
		consultationModel.StartTime = *consultation.StartTime
	}
	if consultation.EndTime != nil {
		consultationModel.EndTime = *consultation.EndTime
	}

	err := d.data.db.WithContext(ctx).Save(consultationModel).Error

	if err != nil {
		d.logger.WithContext(ctx).Errorf("更新问诊信息失败: id=%d, error=%v", consultation.ID, err)
		return fmt.Errorf("更新问诊信息失败: %w", err)
	}

	d.logger.WithContext(ctx).Infof("问诊信息更新成功: id=%d", consultation.ID)
	return nil
}

// DeleteConsultation 删除问诊
func (d *ConsultationData) DeleteConsultation(ctx context.Context, id uint) error {
	err := d.data.db.WithContext(ctx).Delete(&model.Consultations{}, id).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("删除问诊失败: id=%d, error=%v", id, err)
		return fmt.Errorf("删除问诊失败: %w", err)
	}

	d.logger.WithContext(ctx).Infof("问诊删除成功: id=%d", id)
	return nil
}

// GetConsultationList 获取问诊列表
func (d *ConsultationData) GetConsultationList(ctx context.Context, query *biz.ConsultationListQuery) (*biz.ConsultationListResult, error) {
	var models []model.Consultations
	var total int64

	db := d.data.db.WithContext(ctx).Model(&model.Consultations{})

	// 添加筛选条件
	if query.PatientID > 0 {
		db = db.Where("patient_id = ?", query.PatientID)
	}
	if query.DoctorID > 0 {
		db = db.Where("doctor_id = ?", query.DoctorID)
	}
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	if query.Type != "" {
		db = db.Where("type = ?", query.Type)
	}

	// 获取总数
	err := db.Count(&total).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("查询问诊总数失败: error=%v", err)
		return nil, fmt.Errorf("查询问诊总数失败: %w", err)
	}

	// 分页查询
	offset := (query.Page - 1) * query.PageSize
	err = db.Offset(offset).Limit(query.PageSize).Order("created_at DESC").Find(&models).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("查询问诊列表失败: error=%v", err)
		return nil, fmt.Errorf("查询问诊列表失败: %w", err)
	}

	// 转换为业务实体
	consultations := make([]*biz.Consultation, len(models))
	for i, model := range models {
		consultations[i] = d.modelToEntity(&model)
	}

	return &biz.ConsultationListResult{
		Consultations: consultations,
		Total:         total,
		Page:          query.Page,
		PageSize:      query.PageSize,
	}, nil
}

// GetConsultationsByType 按类型获取问诊
func (d *ConsultationData) GetConsultationsByType(ctx context.Context, consultationType string, page, pageSize int, status string) (*biz.ConsultationListResult, error) {
	var models []model.Consultations
	var total int64

	db := d.data.db.WithContext(ctx).Model(&model.Consultations{}).Where("type = ?", consultationType)
	if status != "" {
		db = db.Where("status = ?", status)
	}

	// 获取总数
	err := db.Count(&total).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("查询类型问诊总数失败: type=%s, error=%v", consultationType, err)
		return nil, fmt.Errorf("查询类型问诊总数失败: %w", err)
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = db.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&models).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("查询类型问诊列表失败: type=%s, error=%v", consultationType, err)
		return nil, fmt.Errorf("查询类型问诊列表失败: %w", err)
	}

	// 转换为业务实体
	consultations := make([]*biz.Consultation, len(models))
	for i, model := range models {
		consultations[i] = d.modelToEntity(&model)
	}

	return &biz.ConsultationListResult{
		Consultations: consultations,
		Total:         total,
		Page:          page,
		PageSize:      pageSize,
	}, nil
}

// CreateMessage 创建消息
func (d *ConsultationData) CreateMessage(ctx context.Context, message *biz.ConsultationMessage) error {
	messageModel := &model.ConsultationMessages{
		ConsultationId: uint64(message.ConsultationID),
		SenderId:       uint64(message.SenderID),
		SenderType:     message.SenderType,
		MessageType:    message.MessageType,
		Content:        message.Content,
		MediaUrl:       message.MediaURL,
		IsRead:         message.IsRead,
	}

	err := d.data.db.WithContext(ctx).Create(messageModel).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("创建消息失败: consultation_id=%d, error=%v", message.ConsultationID, err)
		return fmt.Errorf("创建消息失败: %w", err)
	}

	message.ID = uint(messageModel.Id)
	message.CreatedAt = messageModel.CreatedAt
	message.UpdatedAt = messageModel.UpdatedAt

	d.logger.WithContext(ctx).Infof("消息创建成功: id=%d, consultation_id=%d", messageModel.Id, message.ConsultationID)
	return nil
}

// GetMessageByID 根据ID获取消息
func (d *ConsultationData) GetMessageByID(ctx context.Context, id uint) (*biz.ConsultationMessage, error) {
	var messageModel model.ConsultationMessages

	err := d.data.db.WithContext(ctx).Where("id = ?", id).First(&messageModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("消息不存在")
		}
		d.logger.WithContext(ctx).Errorf("查询消息失败: id=%d, error=%v", id, err)
		return nil, fmt.Errorf("查询消息失败: %w", err)
	}

	return d.messageModelToEntity(&messageModel), nil
}

// GetMessageList 获取消息列表
func (d *ConsultationData) GetMessageList(ctx context.Context, query *biz.MessageListQuery) (*biz.MessageListResult, error) {
	var models []model.ConsultationMessages
	var total int64

	db := d.data.db.WithContext(ctx).Model(&model.ConsultationMessages{}).Where("consultation_id = ?", query.ConsultationID)

	// 获取总数
	err := db.Count(&total).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("查询消息总数失败: consultation_id=%d, error=%v", query.ConsultationID, err)
		return nil, fmt.Errorf("查询消息总数失败: %w", err)
	}

	// 分页查询
	offset := (query.Page - 1) * query.PageSize
	err = db.Offset(offset).Limit(query.PageSize).Order("created_at ASC").Find(&models).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("查询消息列表失败: consultation_id=%d, error=%v", query.ConsultationID, err)
		return nil, fmt.Errorf("查询消息列表失败: %w", err)
	}

	// 转换为业务实体
	messages := make([]*biz.ConsultationMessage, len(models))
	for i, model := range models {
		messages[i] = d.messageModelToEntity(&model)
	}

	return &biz.MessageListResult{
		Messages: messages,
		Total:    total,
		Page:     query.Page,
		PageSize: query.PageSize,
	}, nil
}

// UpdateMessage 更新消息
func (d *ConsultationData) UpdateMessage(ctx context.Context, message *biz.ConsultationMessage) error {
	messageModel := &model.ConsultationMessages{
		Id:             uint64(message.ID),
		ConsultationId: uint64(message.ConsultationID),
		SenderId:       uint64(message.SenderID),
		SenderType:     message.SenderType,
		MessageType:    message.MessageType,
		Content:        message.Content,
		MediaUrl:       message.MediaURL,
		IsRead:         message.IsRead,
		UpdatedAt:      time.Now(),
	}

	if message.ReadTime != nil {
		messageModel.ReadTime = *message.ReadTime
	}

	err := d.data.db.WithContext(ctx).Save(messageModel).Error

	if err != nil {
		d.logger.WithContext(ctx).Errorf("更新消息失败: id=%d, error=%v", message.ID, err)
		return fmt.Errorf("更新消息失败: %w", err)
	}

	d.logger.WithContext(ctx).Infof("消息更新成功: id=%d", message.ID)
	return nil
}

// MarkMessageAsRead 标记消息已读
func (d *ConsultationData) MarkMessageAsRead(ctx context.Context, messageID uint, readerID uint) error {
	now := time.Now()
	err := d.data.db.WithContext(ctx).Model(&model.ConsultationMessages{}).
		Where("id = ?", messageID).
		Updates(map[string]interface{}{
			"is_read":    true,
			"read_time":  &now,
			"updated_at": now,
		}).Error

	if err != nil {
		d.logger.WithContext(ctx).Errorf("标记消息已读失败: message_id=%d, reader_id=%d, error=%v", messageID, readerID, err)
		return fmt.Errorf("标记消息已读失败: %w", err)
	}

	d.logger.WithContext(ctx).Infof("消息已标记为已读: message_id=%d, reader_id=%d", messageID, readerID)
	return nil
}

// CreateRecord 创建记录
func (d *ConsultationData) CreateRecord(ctx context.Context, record *biz.ConsultationRecord) error {
	recordModel := &model.ConsultationRecords{
		ConsultationId: uint64(record.ConsultationID),
		RecordType:     record.RecordType,
		Title:          record.Title,
		Content:        record.Content,
		DataValue:      record.DataValue,
		FileUrl:        record.FileURL,
		CreatedBy:      uint64(record.CreatedBy),
	}

	err := d.data.db.WithContext(ctx).Create(recordModel).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("创建记录失败: consultation_id=%d, error=%v", record.ConsultationID, err)
		return fmt.Errorf("创建记录失败: %w", err)
	}

	record.ID = uint(recordModel.Id)
	record.CreatedAt = recordModel.CreatedAt
	record.UpdatedAt = recordModel.UpdatedAt

	d.logger.WithContext(ctx).Infof("记录创建成功: id=%d, consultation_id=%d", recordModel.Id, record.ConsultationID)
	return nil
}

// GetRecordsByConsultationID 根据问诊ID获取记录
func (d *ConsultationData) GetRecordsByConsultationID(ctx context.Context, consultationID uint, recordType string) ([]*biz.ConsultationRecord, error) {
	var models []model.ConsultationRecords

	db := d.data.db.WithContext(ctx).Where("consultation_id = ?", consultationID)
	if recordType != "" {
		db = db.Where("record_type = ?", recordType)
	}

	err := db.Order("created_at ASC").Find(&models).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("查询记录失败: consultation_id=%d, error=%v", consultationID, err)
		return nil, fmt.Errorf("查询记录失败: %w", err)
	}

	// 转换为业务实体
	records := make([]*biz.ConsultationRecord, len(models))
	for i, model := range models {
		records[i] = d.recordModelToEntity(&model)
	}

	return records, nil
}

// GetRecordByID 根据ID获取记录
func (d *ConsultationData) GetRecordByID(ctx context.Context, id uint) (*biz.ConsultationRecord, error) {
	var recordModel model.ConsultationRecords

	err := d.data.db.WithContext(ctx).Where("id = ?", id).First(&recordModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("记录不存在")
		}
		d.logger.WithContext(ctx).Errorf("查询记录失败: id=%d, error=%v", id, err)
		return nil, fmt.Errorf("查询记录失败: %w", err)
	}

	return d.recordModelToEntity(&recordModel), nil
}

// UpdateRecord 更新记录
func (d *ConsultationData) UpdateRecord(ctx context.Context, record *biz.ConsultationRecord) error {
	recordModel := &model.ConsultationRecords{
		Id:             uint64(record.ID),
		ConsultationId: uint64(record.ConsultationID),
		RecordType:     record.RecordType,
		Title:          record.Title,
		Content:        record.Content,
		DataValue:      record.DataValue,
		FileUrl:        record.FileURL,
		CreatedBy:      uint64(record.CreatedBy),
		UpdatedAt:      time.Now(),
	}

	err := d.data.db.WithContext(ctx).Save(recordModel).Error

	if err != nil {
		d.logger.WithContext(ctx).Errorf("更新记录失败: id=%d, error=%v", record.ID, err)
		return fmt.Errorf("更新记录失败: %w", err)
	}

	d.logger.WithContext(ctx).Infof("记录更新成功: id=%d", record.ID)
	return nil
}

// modelToEntity 将数据模型转换为业务实体
func (d *ConsultationData) modelToEntity(consultationModel *model.Consultations) *biz.Consultation {
	return &biz.Consultation{
		ID:               uint(consultationModel.Id),
		ConsultationCode: consultationModel.ConsultationCode,
		PatientID:        uint(consultationModel.PatientId),
		DoctorID:         uint(consultationModel.DoctorId),
		Type:             consultationModel.Type,
		Status:           consultationModel.Status,
		ChiefComplaint:   consultationModel.ChiefComplaint,
		PresentIllness:   consultationModel.PresentIllness,
		Symptoms:         consultationModel.Symptoms,
		Diagnosis:        consultationModel.Diagnosis,
		Treatment:        consultationModel.Treatment,
		Prescription:     consultationModel.Prescription,
		Advice:           consultationModel.Advice,
		StartTime:        &consultationModel.StartTime,
		EndTime:          &consultationModel.EndTime,
		Duration:         consultationModel.Duration,
		Fee:              consultationModel.Fee,
		PaymentStatus:    consultationModel.PaymentStatus,
		Rating:           consultationModel.Rating,
		Feedback:         consultationModel.Feedback,
		Remarks:          consultationModel.Remarks,
		CreatedAt:        consultationModel.CreatedAt,
		UpdatedAt:        consultationModel.UpdatedAt,
	}
}

// messageModelToEntity 将消息数据模型转换为业务实体
func (d *ConsultationData) messageModelToEntity(messageModel *model.ConsultationMessages) *biz.ConsultationMessage {
	return &biz.ConsultationMessage{
		ID:             uint(messageModel.Id),
		ConsultationID: uint(messageModel.ConsultationId),
		SenderID:       uint(messageModel.SenderId),
		SenderType:     messageModel.SenderType,
		MessageType:    messageModel.MessageType,
		Content:        messageModel.Content,
		MediaURL:       messageModel.MediaUrl,
		IsRead:         messageModel.IsRead,
		ReadTime:       &messageModel.ReadTime,
		CreatedAt:      messageModel.CreatedAt,
		UpdatedAt:      messageModel.UpdatedAt,
	}
}

// recordModelToEntity 将记录数据模型转换为业务实体
func (d *ConsultationData) recordModelToEntity(recordModel *model.ConsultationRecords) *biz.ConsultationRecord {
	return &biz.ConsultationRecord{
		ID:             uint(recordModel.Id),
		ConsultationID: uint(recordModel.ConsultationId),
		RecordType:     recordModel.RecordType,
		Title:          recordModel.Title,
		Content:        recordModel.Content,
		DataValue:      recordModel.DataValue,
		FileURL:        recordModel.FileUrl,
		CreatedBy:      uint(recordModel.CreatedBy),
		CreatedAt:      recordModel.CreatedAt,
		UpdatedAt:      recordModel.UpdatedAt,
	}
}

// GetDoctorByID 根据ID获取医生信息（补充实现）
func (d *ConsultationData) GetDoctorByID(ctx context.Context, id uint) (*biz.Doctor, error) {
	var doctorModel model.Doctors

	err := d.data.db.WithContext(ctx).Where("id = ?", id).First(&doctorModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("医生不存在")
		}
		d.logger.WithContext(ctx).Errorf("查询医生失败: id=%d, error=%v", id, err)
		return nil, fmt.Errorf("查询医生失败: %w", err)
	}

	doctor := d.modelToDoctorEntity(&doctorModel)
	return doctor, nil
}

// modelToDoctorEntity 将医生数据模型转换为业务实体
func (d *ConsultationData) modelToDoctorEntity(doctorModel *model.Doctors) *biz.Doctor {
	var departmentID, hospitalID *uint
	if doctorModel.DepartmentId != 0 {
		v := uint(doctorModel.DepartmentId)
		departmentID = &v
	}
	if doctorModel.HospitalId != 0 {
		v := uint(doctorModel.HospitalId)
		hospitalID = &v
	}
	return &biz.Doctor{
		ID:            uint(doctorModel.Id),
		DoctorCode:    doctorModel.DoctorCode,
		Phone:         doctorModel.Phone,
		Password:      doctorModel.Password,
		Name:          doctorModel.Name,
		Gender:        doctorModel.Gender,
		BirthDate:     doctorModel.BirthDate,
		Email:         doctorModel.Email,
		Avatar:        doctorModel.Avatar,
		LicenseNumber: doctorModel.LicenseNumber,
		DepartmentID:  departmentID,
		HospitalID:    hospitalID,
		Title:         doctorModel.Title,
		Speciality:    doctorModel.Speciality,
		PracticeScope: doctorModel.PracticeScope,
		Status:        doctorModel.Status,
		LastLoginTime: doctorModel.LastLoginTime,
		LastLoginIP:   doctorModel.LastLoginIp,
		CreatedAt:     doctorModel.CreatedAt,
		UpdatedAt:     doctorModel.UpdatedAt,
	}
}

package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"kratos_client/internal/biz"
)

// 处方数据模型
type MtPrescription struct {
	ID                uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	PrescriptionNo    string          `gorm:"column:prescription_no;size:32;not null;uniqueIndex" json:"prescription_no"`
	DoctorID          uint64          `gorm:"column:doctor_id;not null" json:"doctor_id"`
	PatientID         uint64          `gorm:"column:patient_id;not null" json:"patient_id"`
	MedicalRecordID   *uint64         `gorm:"column:medical_record_id" json:"medical_record_id"`
	PrescriptionDate  time.Time       `gorm:"column:prescription_date;type:date;not null" json:"prescription_date"`
	TotalAmount       decimal.Decimal `gorm:"column:total_amount;type:decimal(10,2);default:0.00;not null" json:"total_amount"`
	PrescriptionType  string          `gorm:"column:prescription_type;size:20;default:'西药'" json:"prescription_type"`
	UsageInstruction  string          `gorm:"column:usage_instruction;type:text" json:"usage_instruction"`
	Status            string          `gorm:"column:status;size:20;default:'已开具';not null" json:"status"`
	AuditorID         *uint64         `gorm:"column:auditor_id" json:"auditor_id"`
	AuditTime         *time.Time      `gorm:"column:audit_time" json:"audit_time"`
	AuditNotes        string          `gorm:"column:audit_notes;size:500" json:"audit_notes"`
	CreatedAt         time.Time       `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (MtPrescription) TableName() string {
	return "mt_prescriptions"
}

// 处方药品明细数据模型
type MtPrescriptionMedicine struct {
	ID             uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	PrescriptionID uint64          `gorm:"column:prescription_id;not null" json:"prescription_id"`
	MedicineID     uint64          `gorm:"column:medicine_id;not null" json:"medicine_id"`
	Quantity       decimal.Decimal `gorm:"column:quantity;type:decimal(10,2);not null" json:"quantity"`
	Unit           string          `gorm:"column:unit;size:20;not null" json:"unit"`
	UnitPrice      decimal.Decimal `gorm:"column:unit_price;type:decimal(10,2);not null" json:"unit_price"`
	TotalPrice     decimal.Decimal `gorm:"column:total_price;type:decimal(10,2);not null" json:"total_price"`
	Dosage         string          `gorm:"column:dosage;size:100" json:"dosage"`
	Frequency      string          `gorm:"column:frequency;size:50" json:"frequency"`
	Duration       string          `gorm:"column:duration;size:50" json:"duration"`
	UsageMethod    string          `gorm:"column:usage_method;size:100" json:"usage_method"`
	Notes          string          `gorm:"column:notes;size:500" json:"notes"`
	CreatedAt      time.Time       `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (MtPrescriptionMedicine) TableName() string {
	return "mt_prescription_medicines"
}

// 处方仓储实现
type prescriptionRepo struct {
	data *Data
	log  *log.Helper
}

// 创建处方仓储
func NewPrescriptionRepo(data *Data, logger log.Logger) biz.PrescriptionRepo {
	return &prescriptionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 转换数据模型到业务模型
func (r *prescriptionRepo) toBizPrescription(do *MtPrescription) *biz.MtPrescription {
	return &biz.MtPrescription{
		ID:                do.ID,
		PrescriptionNo:    do.PrescriptionNo,
		DoctorID:          do.DoctorID,
		PatientID:         do.PatientID,
		MedicalRecordID:   do.MedicalRecordID,
		PrescriptionDate:  do.PrescriptionDate,
		TotalAmount:       do.TotalAmount,
		PrescriptionType:  do.PrescriptionType,
		UsageInstruction:  do.UsageInstruction,
		Status:            do.Status,
		AuditorID:         do.AuditorID,
		AuditTime:         do.AuditTime,
		AuditNotes:        do.AuditNotes,
		CreatedAt:         do.CreatedAt,
		UpdatedAt:         do.UpdatedAt,
	}
}

// 转换处方药品明细
func (r *prescriptionRepo) toBizPrescriptionMedicine(do *MtPrescriptionMedicine) *biz.MtPrescriptionMedicine {
	return &biz.MtPrescriptionMedicine{
		ID:             do.ID,
		PrescriptionID: do.PrescriptionID,
		MedicineID:     do.MedicineID,
		Quantity:       do.Quantity,
		Unit:           do.Unit,
		UnitPrice:      do.UnitPrice,
		TotalPrice:     do.TotalPrice,
		Dosage:         do.Dosage,
		Frequency:      do.Frequency,
		Duration:       do.Duration,
		UsageMethod:    do.UsageMethod,
		Notes:          do.Notes,
		CreatedAt:      do.CreatedAt,
	}
}

// 获取处方列表
func (r *prescriptionRepo) ListPrescriptions(ctx context.Context, req *biz.ListPrescriptionsRequest) ([]*biz.MtPrescription, int64, error) {
	var prescriptions []MtPrescription
	var total int64

	db := r.data.Db.WithContext(ctx).Model(&MtPrescription{})

	// 添加过滤条件
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if req.PrescriptionType != "" {
		db = db.Where("prescription_type = ?", req.PrescriptionType)
	}
	if req.StartDate != nil {
		db = db.Where("prescription_date >= ?", req.StartDate.Format("2006-01-02"))
	}
	if req.EndDate != nil {
		db = db.Where("prescription_date <= ?", req.EndDate.Format("2006-01-02"))
	}

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		r.log.Errorf("查询处方总数失败: %v", err)
		return nil, 0, err
	}

	// 查询列表
	offset := (req.Page - 1) * req.PageSize
	if err := db.Order("created_at DESC").Offset(int(offset)).Limit(int(req.PageSize)).Find(&prescriptions).Error; err != nil {
		r.log.Errorf("查询处方列表失败: %v", err)
		return nil, 0, err
	}

	// 转换为业务模型并填充扩展信息
	bizPrescriptions := make([]*biz.MtPrescription, len(prescriptions))
	for i, prescription := range prescriptions {
		bizPrescription := r.toBizPrescription(&prescription)
		
		// 填充扩展信息
		r.fillPrescriptionExtInfo(ctx, bizPrescription)
		
		bizPrescriptions[i] = bizPrescription
	}

	return bizPrescriptions, total, nil
}

// 获取患者处方列表
func (r *prescriptionRepo) ListPatientPrescriptions(ctx context.Context, req *biz.ListPatientPrescriptionsRequest) ([]*biz.MtPrescription, int64, error) {
	var prescriptions []MtPrescription
	var total int64

	db := r.data.Db.WithContext(ctx).Model(&MtPrescription{}).Where("patient_id = ?", req.PatientID)

	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		r.log.Errorf("查询患者处方总数失败: %v", err)
		return nil, 0, err
	}

	// 查询列表
	offset := (req.Page - 1) * req.PageSize
	if err := db.Order("created_at DESC").Offset(int(offset)).Limit(int(req.PageSize)).Find(&prescriptions).Error; err != nil {
		r.log.Errorf("查询患者处方列表失败: %v", err)
		return nil, 0, err
	}

	// 转换为业务模型并填充扩展信息
	bizPrescriptions := make([]*biz.MtPrescription, len(prescriptions))
	for i, prescription := range prescriptions {
		bizPrescription := r.toBizPrescription(&prescription)
		
		// 填充扩展信息
		r.fillPrescriptionExtInfo(ctx, bizPrescription)
		
		bizPrescriptions[i] = bizPrescription
	}

	return bizPrescriptions, total, nil
}

// 获取医生处方列表
func (r *prescriptionRepo) ListDoctorPrescriptions(ctx context.Context, req *biz.ListDoctorPrescriptionsRequest) ([]*biz.MtPrescription, int64, error) {
	var prescriptions []MtPrescription
	var total int64

	db := r.data.Db.WithContext(ctx).Model(&MtPrescription{}).Where("doctor_id = ?", req.DoctorID)

	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	// 查询总数
	if err := db.Count(&total).Error; err != nil {
		r.log.Errorf("查询医生处方总数失败: %v", err)
		return nil, 0, err
	}

	// 查询列表
	offset := (req.Page - 1) * req.PageSize
	if err := db.Order("created_at DESC").Offset(int(offset)).Limit(int(req.PageSize)).Find(&prescriptions).Error; err != nil {
		r.log.Errorf("查询医生处方列表失败: %v", err)
		return nil, 0, err
	}

	// 转换为业务模型并填充扩展信息
	bizPrescriptions := make([]*biz.MtPrescription, len(prescriptions))
	for i, prescription := range prescriptions {
		bizPrescription := r.toBizPrescription(&prescription)
		
		// 填充扩展信息
		r.fillPrescriptionExtInfo(ctx, bizPrescription)
		
		bizPrescriptions[i] = bizPrescription
	}

	return bizPrescriptions, total, nil
}

// 根据ID获取处方
func (r *prescriptionRepo) GetPrescriptionByID(ctx context.Context, id uint64) (*biz.MtPrescription, error) {
	var prescription MtPrescription
	if err := r.data.Db.WithContext(ctx).Where("id = ?", id).First(&prescription).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.log.Errorf("查询处方失败: %v", err)
		return nil, err
	}

	bizPrescription := r.toBizPrescription(&prescription)
	
	// 填充扩展信息
	r.fillPrescriptionExtInfo(ctx, bizPrescription)

	return bizPrescription, nil
}

// 获取处方药品明细
func (r *prescriptionRepo) GetPrescriptionMedicines(ctx context.Context, prescriptionID uint64) ([]*biz.MtPrescriptionMedicine, error) {
	var medicines []MtPrescriptionMedicine
	if err := r.data.Db.WithContext(ctx).Where("prescription_id = ?", prescriptionID).Order("id ASC").Find(&medicines).Error; err != nil {
		r.log.Errorf("查询处方药品明细失败: %v", err)
		return nil, err
	}

	// 转换为业务模型并填充药品信息
	bizMedicines := make([]*biz.MtPrescriptionMedicine, len(medicines))
	for i, medicine := range medicines {
		bizMedicine := r.toBizPrescriptionMedicine(&medicine)
		
		// 填充药品信息
		r.fillMedicineExtInfo(ctx, bizMedicine)
		
		bizMedicines[i] = bizMedicine
	}

	return bizMedicines, nil
}

// 填充处方扩展信息
func (r *prescriptionRepo) fillPrescriptionExtInfo(ctx context.Context, prescription *biz.MtPrescription) {
	// 这里可以根据实际情况填充医生名称、患者名称、审核医生名称等
	// 由于没有用户表的具体结构，这里先用占位符
	prescription.DoctorName = "医生姓名" // 实际应该从用户表查询
	prescription.PatientName = "患者姓名" // 实际应该从用户表查询
	if prescription.AuditorID != nil {
		prescription.AuditorName = "审核医生姓名" // 实际应该从用户表查询
	}

	// 统计药品种类数量
	var count int64
	if err := r.data.Db.WithContext(ctx).Model(&MtPrescriptionMedicine{}).Where("prescription_id = ?", prescription.ID).Count(&count).Error; err != nil {
		r.log.Errorf("统计处方药品数量失败: %v", err)
	} else {
		prescription.MedicineCount = int32(count)
	}
}

// 填充药品扩展信息
func (r *prescriptionRepo) fillMedicineExtInfo(ctx context.Context, medicine *biz.MtPrescriptionMedicine) {
	// 这里可以根据实际情况从药品表查询药品名称、规格、厂家等信息
	// 由于没有药品表的具体结构，这里先用占位符
	medicine.MedicineName = "药品名称" // 实际应该从药品表查询
	medicine.MedicineSpec = "药品规格" // 实际应该从药品表查询
	medicine.Manufacturer = "生产厂家" // 实际应该从药品表查询
}
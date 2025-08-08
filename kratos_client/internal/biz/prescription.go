package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
)

// 处方模型
type MtPrescription struct {
	ID                uint64          `json:"id"`
	PrescriptionNo    string          `json:"prescription_no"`
	DoctorID          uint64          `json:"doctor_id"`
	PatientID         uint64          `json:"patient_id"`
	MedicalRecordID   *uint64         `json:"medical_record_id"`
	PrescriptionDate  time.Time       `json:"prescription_date"`
	TotalAmount       decimal.Decimal `json:"total_amount"`
	PrescriptionType  string          `json:"prescription_type"`
	UsageInstruction  string          `json:"usage_instruction"`
	Status            string          `json:"status"`
	AuditorID         *uint64         `json:"auditor_id"`
	AuditTime         *time.Time      `json:"audit_time"`
	AuditNotes        string          `json:"audit_notes"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	
	// 扩展信息
	DoctorName        string          `json:"doctor_name"`
	PatientName       string          `json:"patient_name"`
	AuditorName       string          `json:"auditor_name"`
	MedicineCount     int32           `json:"medicine_count"`
}

// 处方药品明细模型
type MtPrescriptionMedicine struct {
	ID             uint64          `json:"id"`
	PrescriptionID uint64          `json:"prescription_id"`
	MedicineID     uint64          `json:"medicine_id"`
	Quantity       decimal.Decimal `json:"quantity"`
	Unit           string          `json:"unit"`
	UnitPrice      decimal.Decimal `json:"unit_price"`
	TotalPrice     decimal.Decimal `json:"total_price"`
	Dosage         string          `json:"dosage"`
	Frequency      string          `json:"frequency"`
	Duration       string          `json:"duration"`
	UsageMethod    string          `json:"usage_method"`
	Notes          string          `json:"notes"`
	CreatedAt      time.Time       `json:"created_at"`
	
	// 扩展信息
	MedicineName   string          `json:"medicine_name"`
	MedicineSpec   string          `json:"medicine_spec"`
	Manufacturer   string          `json:"manufacturer"`
}

// 处方详情
type PrescriptionDetail struct {
	Prescription *MtPrescription            `json:"prescription"`
	Medicines    []*MtPrescriptionMedicine  `json:"medicines"`
}

// 处方列表查询请求
type ListPrescriptionsRequest struct {
	Status           string    `json:"status"`
	PrescriptionType string    `json:"prescription_type"`
	StartDate        *time.Time `json:"start_date"`
	EndDate          *time.Time `json:"end_date"`
	Page             int32     `json:"page"`
	PageSize         int32     `json:"page_size"`
}

// 患者处方列表查询请求
type ListPatientPrescriptionsRequest struct {
	PatientID uint64 `json:"patient_id"`
	Status    string `json:"status"`
	Page      int32  `json:"page"`
	PageSize  int32  `json:"page_size"`
}

// 医生处方列表查询请求
type ListDoctorPrescriptionsRequest struct {
	DoctorID uint64 `json:"doctor_id"`
	Status   string `json:"status"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"page_size"`
}

// 处方仓储接口
type PrescriptionRepo interface {
	// 处方CRUD
	ListPrescriptions(ctx context.Context, req *ListPrescriptionsRequest) ([]*MtPrescription, int64, error)
	ListPatientPrescriptions(ctx context.Context, req *ListPatientPrescriptionsRequest) ([]*MtPrescription, int64, error)
	ListDoctorPrescriptions(ctx context.Context, req *ListDoctorPrescriptionsRequest) ([]*MtPrescription, int64, error)
	GetPrescriptionByID(ctx context.Context, id uint64) (*MtPrescription, error)
	
	// 处方药品明细
	GetPrescriptionMedicines(ctx context.Context, prescriptionID uint64) ([]*MtPrescriptionMedicine, error)
}

// 处方用例
type PrescriptionUsecase struct {
	prescriptionRepo PrescriptionRepo
	log              *log.Helper
}

// 创建处方用例
func NewPrescriptionUsecase(prescriptionRepo PrescriptionRepo, logger log.Logger) *PrescriptionUsecase {
	return &PrescriptionUsecase{
		prescriptionRepo: prescriptionRepo,
		log:              log.NewHelper(logger),
	}
}

// 获取处方列表
func (uc *PrescriptionUsecase) ListPrescriptions(ctx context.Context, req *ListPrescriptionsRequest) ([]*MtPrescription, int64, error) {
	prescriptions, total, err := uc.prescriptionRepo.ListPrescriptions(ctx, req)
	if err != nil {
		uc.log.Errorf("获取处方列表失败: %v", err)
		return nil, 0, err
	}

	return prescriptions, total, nil
}

// 获取患者处方列表
func (uc *PrescriptionUsecase) ListPatientPrescriptions(ctx context.Context, req *ListPatientPrescriptionsRequest) ([]*MtPrescription, int64, error) {
	prescriptions, total, err := uc.prescriptionRepo.ListPatientPrescriptions(ctx, req)
	if err != nil {
		uc.log.Errorf("获取患者处方列表失败: %v", err)
		return nil, 0, err
	}

	return prescriptions, total, nil
}

// 获取医生处方列表
func (uc *PrescriptionUsecase) ListDoctorPrescriptions(ctx context.Context, req *ListDoctorPrescriptionsRequest) ([]*MtPrescription, int64, error) {
	prescriptions, total, err := uc.prescriptionRepo.ListDoctorPrescriptions(ctx, req)
	if err != nil {
		uc.log.Errorf("获取医生处方列表失败: %v", err)
		return nil, 0, err
	}

	return prescriptions, total, nil
}

// 获取处方详情
func (uc *PrescriptionUsecase) GetPrescriptionDetail(ctx context.Context, prescriptionID uint64) (*PrescriptionDetail, error) {
	// 获取处方基本信息
	prescription, err := uc.prescriptionRepo.GetPrescriptionByID(ctx, prescriptionID)
	if err != nil {
		uc.log.Errorf("获取处方信息失败: %v", err)
		return nil, err
	}

	if prescription == nil {
		return nil, fmt.Errorf("处方不存在: ID=%d", prescriptionID)
	}

	// 获取处方药品明细
	medicines, err := uc.prescriptionRepo.GetPrescriptionMedicines(ctx, prescriptionID)
	if err != nil {
		uc.log.Errorf("获取处方药品明细失败: %v", err)
		return nil, err
	}

	return &PrescriptionDetail{
		Prescription: prescription,
		Medicines:    medicines,
	}, nil
}

// 处方状态常量
const (
	PrescriptionStatusCancelled = "已取消"
	PrescriptionStatusIssued    = "已开具"
	PrescriptionStatusAudited   = "已审核"
	PrescriptionStatusDispensed = "已发药"
)

// 处方类型常量
const (
	PrescriptionTypeWestern  = "西药"
	PrescriptionTypeChinese  = "中药"
	PrescriptionTypeMixed    = "中西药"
)

// 获取处方状态文本
func GetPrescriptionStatusText(status string) string {
	switch status {
	case PrescriptionStatusCancelled:
		return "已取消"
	case PrescriptionStatusIssued:
		return "已开具"
	case PrescriptionStatusAudited:
		return "已审核"
	case PrescriptionStatusDispensed:
		return "已发药"
	default:
		return "未知状态"
	}
}

// 获取处方类型文本
func GetPrescriptionTypeText(prescriptionType string) string {
	switch prescriptionType {
	case PrescriptionTypeWestern:
		return "西药"
	case PrescriptionTypeChinese:
		return "中药"
	case PrescriptionTypeMixed:
		return "中西药"
	default:
		return "其他"
	}
}
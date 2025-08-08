package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

// 处方状态常量
const (
	PrescriptionStatusDraft     = "草稿"
	PrescriptionStatusPending   = "待审核"
	PrescriptionStatusConfirmed = "已审核"
	PrescriptionStatusDispensed = "已发药"
	PrescriptionStatusCancelled = "已取消"
)

// 处方类型常量
const (
	PrescriptionTypeWestern = "西药"
	PrescriptionTypeChinese = "中药"
	PrescriptionTypeMixed   = "中西药"
)

// 处方业务实体
type Prescription struct {
	ID               uint64                 `json:"id"`
	PrescriptionNo   string                 `json:"prescription_no"`
	DoctorID         uint64                 `json:"doctor_id"`
	PatientID        uint64                 `json:"patient_id"`
	MedicalRecordID  uint64                 `json:"medical_record_id"`
	AppointmentID    uint64                 `json:"appointment_id"` // 关联挂号表
	PrescriptionDate time.Time              `json:"prescription_date"`
	TotalAmount      float64                `json:"total_amount"`
	PrescriptionType string                 `json:"prescription_type"`
	UsageInstruction string                 `json:"usage_instruction"`
	Status           string                 `json:"status"`
	AuditorID        uint64                 `json:"auditor_id"`
	AuditTime        time.Time              `json:"audit_time"`
	AuditNotes       string                 `json:"audit_notes"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
	Medicines        []PrescriptionMedicine `json:"medicines"`
	ChargeItems      []ChargeItem           `json:"charge_items"` // 收费项目
}

// 处方药品业务实体
type PrescriptionMedicine struct {
	ID             uint64  `json:"id"`
	PrescriptionID uint64  `json:"prescription_id"`
	MedicineID     uint64  `json:"medicine_id"`
	MedicineName   string  `json:"medicine_name"`
	MedicineCode   string  `json:"medicine_code"`
	Specification  string  `json:"specification"`
	Manufacturer   string  `json:"manufacturer"`
	Quantity       float64 `json:"quantity"`
	Unit           string  `json:"unit"`
	UnitPrice      float64 `json:"unit_price"`
	TotalPrice     float64 `json:"total_price"`
	Dosage         string  `json:"dosage"`
	Frequency      string  `json:"frequency"`
	Duration       string  `json:"duration"`
	UsageMethod    string  `json:"usage_method"`
	Notes          string  `json:"notes"`
}

// 药品业务实体
type Medicine struct {
	ID                uint64  `json:"id"`
	MedicineCode      string  `json:"medicine_code"`
	Name              string  `json:"name"`
	GenericName       string  `json:"generic_name"`
	BrandName         string  `json:"brand_name"`
	Specification     string  `json:"specification"`
	DosageForm        string  `json:"dosage_form"`
	Manufacturer      string  `json:"manufacturer"`
	ApprovalNumber    string  `json:"approval_number"`
	Category          string  `json:"category"`
	PrescriptionType  string  `json:"prescription_type"`
	Unit              string  `json:"unit"`
	Price             float64 `json:"price"`
	Indications       string  `json:"indications"`
	Contraindications string  `json:"contraindications"`
	SideEffects       string  `json:"side_effects"`
	DosageUsage       string  `json:"dosage_usage"`
	StorageConditions string  `json:"storage_conditions"`
	ShelfLife         int32   `json:"shelf_life"`
	ImageURL          string  `json:"image_url"`
	Status            string  `json:"status"`
	StockQuantity     float64 `json:"stock_quantity"`
	AvailableQuantity float64 `json:"available_quantity"`
	MinStockLevel     float64 `json:"min_stock_level"`
}

// 药物相互作用业务实体
type DrugInteraction struct {
	ID             uint64 `json:"id"`
	MedicineID1    uint64 `json:"medicine_id_1"`
	MedicineName1  string `json:"medicine_name_1"`
	MedicineID2    uint64 `json:"medicine_id_2"`
	MedicineName2  string `json:"medicine_name_2"`
	Severity       string `json:"severity"`
	Description    string `json:"description"`
	Recommendation string `json:"recommendation"`
}

// 收费项目业务实体
type ChargeItem struct {
	ID          uint64  `json:"id"`
	ItemCode    string  `json:"item_code"`
	ItemName    string  `json:"item_name"`
	ItemType    string  `json:"item_type"` // 药品费、检查费、治疗费等
	UnitPrice   float64 `json:"unit_price"`
	Quantity    float64 `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
	Department  string  `json:"department"`
	Description string  `json:"description"`
}

// 挂号信息业务实体
type Appointment struct {
	ID              uint64    `json:"id"`
	AppointmentNo   string    `json:"appointment_no"`
	PatientID       uint64    `json:"patient_id"`
	DoctorID        uint64    `json:"doctor_id"`
	DepartmentID    uint64    `json:"department_id"`
	ScheduleDate    time.Time `json:"schedule_date"`
	TimeSlot        string    `json:"time_slot"`
	Status          string    `json:"status"`
	RegistrationFee float64   `json:"registration_fee"`
	CreatedAt       time.Time `json:"created_at"`
}

// 处方数据访问接口
type PrescriptionRepo interface {
	// 处方CRUD操作
	CreatePrescription(ctx context.Context, prescription *Prescription) error
	UpdatePrescription(ctx context.Context, prescription *Prescription) error
	GetPrescription(ctx context.Context, id uint64) (*Prescription, error)
	GetPrescriptionByNo(ctx context.Context, prescriptionNo string) (*Prescription, error)
	ListPrescriptions(ctx context.Context, doctorID, patientID uint64, status string, startDate, endDate time.Time, page, pageSize int) ([]*Prescription, int64, error)
	DeletePrescription(ctx context.Context, id uint64) error

	// 处方药品操作
	CreatePrescriptionMedicines(ctx context.Context, medicines []PrescriptionMedicine) error
	UpdatePrescriptionMedicines(ctx context.Context, prescriptionID uint64, medicines []PrescriptionMedicine) error
	GetPrescriptionMedicines(ctx context.Context, prescriptionID uint64) ([]PrescriptionMedicine, error)
	DeletePrescriptionMedicines(ctx context.Context, prescriptionID uint64) error

	// 药品相关操作
	SearchMedicines(ctx context.Context, keyword, category, prescriptionType string, onlyAvailable bool, limit int) ([]*Medicine, error)
	GetMedicine(ctx context.Context, id uint64) (*Medicine, error)
	CheckMedicineStock(ctx context.Context, medicineID uint64, quantity float64) (bool, error)
	GetMedicinesByIDs(ctx context.Context, ids []uint64) ([]*Medicine, error)

	// 药物相互作用检查
	CheckDrugInteractions(ctx context.Context, medicineIDs []uint64) ([]*DrugInteraction, error)

	// 缓存操作
	CachePrescriptionDraft(ctx context.Context, doctorID uint64, prescription *Prescription) error
	GetPrescriptionDraft(ctx context.Context, doctorID uint64) (*Prescription, error)
	DeletePrescriptionDraft(ctx context.Context, doctorID uint64) error
}

// 处方业务逻辑用例
type PrescriptionUsecase struct {
	repo PrescriptionRepo
	log  *log.Helper
}

// 创建处方业务逻辑实例
func NewPrescriptionUsecase(repo PrescriptionRepo, logger log.Logger) *PrescriptionUsecase {
	return &PrescriptionUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// 生成处方编号
func (uc *PrescriptionUsecase) generatePrescriptionNo() string {
	now := time.Now()
	return fmt.Sprintf("RX%s%06d", now.Format("20060102"), now.Unix()%1000000)
}

// 创建处方
func (uc *PrescriptionUsecase) CreatePrescription(ctx context.Context, prescription *Prescription) error {
	// 生成处方编号
	prescription.PrescriptionNo = uc.generatePrescriptionNo()
	prescription.Status = PrescriptionStatusDraft
	prescription.PrescriptionDate = time.Now()
	prescription.CreatedAt = time.Now()
	prescription.UpdatedAt = time.Now()

	// 验证处方数据
	if err := uc.validatePrescription(ctx, prescription); err != nil {
		return err
	}

	// 计算总金额
	if err := uc.calculateTotalAmount(ctx, prescription); err != nil {
		return err
	}

	// 保存处方
	if err := uc.repo.CreatePrescription(ctx, prescription); err != nil {
		uc.log.WithContext(ctx).Errorf("创建处方失败: %v", err)
		return errors.InternalServer("CREATE_PRESCRIPTION_FAILED", "创建处方失败")
	}

	// 保存处方药品
	if len(prescription.Medicines) > 0 {
		for i := range prescription.Medicines {
			prescription.Medicines[i].PrescriptionID = prescription.ID
		}
		if err := uc.repo.CreatePrescriptionMedicines(ctx, prescription.Medicines); err != nil {
			uc.log.WithContext(ctx).Errorf("创建处方药品失败: %v", err)
			return errors.InternalServer("CREATE_PRESCRIPTION_MEDICINES_FAILED", "创建处方药品失败")
		}
	}

	uc.log.WithContext(ctx).Infof("处方创建成功: id=%d, no=%s", prescription.ID, prescription.PrescriptionNo)
	return nil
}

// 验证处方数据
func (uc *PrescriptionUsecase) validatePrescription(ctx context.Context, prescription *Prescription) error {
	if prescription.DoctorID == 0 {
		return errors.BadRequest("INVALID_DOCTOR_ID", "医生ID不能为空")
	}
	if prescription.PatientID == 0 {
		return errors.BadRequest("INVALID_PATIENT_ID", "患者ID不能为空")
	}
	if len(prescription.Medicines) == 0 {
		return errors.BadRequest("EMPTY_MEDICINES", "处方药品不能为空")
	}

	// 验证药品库存
	for _, medicine := range prescription.Medicines {
		available, err := uc.repo.CheckMedicineStock(ctx, medicine.MedicineID, medicine.Quantity)
		if err != nil {
			return err
		}
		if !available {
			return errors.BadRequest("INSUFFICIENT_STOCK", fmt.Sprintf("药品 %s 库存不足", medicine.MedicineName))
		}
	}

	return nil
}

// 计算总金额
func (uc *PrescriptionUsecase) calculateTotalAmount(ctx context.Context, prescription *Prescription) error {
	var total float64

	// 获取药品信息并计算价格
	for i, medicine := range prescription.Medicines {
		medicineInfo, err := uc.repo.GetMedicine(ctx, medicine.MedicineID)
		if err != nil {
			return err
		}

		prescription.Medicines[i].MedicineName = medicineInfo.Name
		prescription.Medicines[i].MedicineCode = medicineInfo.MedicineCode
		prescription.Medicines[i].Specification = medicineInfo.Specification
		prescription.Medicines[i].Manufacturer = medicineInfo.Manufacturer
		prescription.Medicines[i].Unit = medicineInfo.Unit
		prescription.Medicines[i].UnitPrice = medicineInfo.Price
		prescription.Medicines[i].TotalPrice = medicineInfo.Price * medicine.Quantity

		total += prescription.Medicines[i].TotalPrice
	}

	prescription.TotalAmount = total
	return nil
}

// 更新处方
func (uc *PrescriptionUsecase) UpdatePrescription(ctx context.Context, prescription *Prescription) error {
	// 检查处方是否存在
	existingPrescription, err := uc.repo.GetPrescription(ctx, prescription.ID)
	if err != nil {
		return errors.NotFound("PRESCRIPTION_NOT_FOUND", "处方不存在")
	}

	// 检查处方状态是否允许修改
	if existingPrescription.Status != PrescriptionStatusDraft {
		return errors.BadRequest("PRESCRIPTION_NOT_EDITABLE", "处方状态不允许修改")
	}

	// 验证处方数据
	if err := uc.validatePrescription(ctx, prescription); err != nil {
		return err
	}

	// 计算总金额
	if err := uc.calculateTotalAmount(ctx, prescription); err != nil {
		return err
	}

	prescription.UpdatedAt = time.Now()

	// 更新处方
	if err := uc.repo.UpdatePrescription(ctx, prescription); err != nil {
		uc.log.WithContext(ctx).Errorf("更新处方失败: %v", err)
		return errors.InternalServer("UPDATE_PRESCRIPTION_FAILED", "更新处方失败")
	}

	// 更新处方药品
	if err := uc.repo.UpdatePrescriptionMedicines(ctx, prescription.ID, prescription.Medicines); err != nil {
		uc.log.WithContext(ctx).Errorf("更新处方药品失败: %v", err)
		return errors.InternalServer("UPDATE_PRESCRIPTION_MEDICINES_FAILED", "更新处方药品失败")
	}

	uc.log.WithContext(ctx).Infof("处方更新成功: id=%d", prescription.ID)
	return nil
}

// 确认处方
func (uc *PrescriptionUsecase) ConfirmPrescription(ctx context.Context, prescriptionID, doctorID uint64, auditNotes string) error {
	// 获取处方
	prescription, err := uc.repo.GetPrescription(ctx, prescriptionID)
	if err != nil {
		return errors.NotFound("PRESCRIPTION_NOT_FOUND", "处方不存在")
	}

	// 验证权限
	if prescription.DoctorID != doctorID {
		return errors.Forbidden("PERMISSION_DENIED", "无权限操作此处方")
	}

	// 验证状态
	if prescription.Status != PrescriptionStatusDraft && prescription.Status != PrescriptionStatusPending {
		return errors.BadRequest("INVALID_STATUS", "处方状态不允许确认")
	}

	// 检查药品库存
	medicines, err := uc.repo.GetPrescriptionMedicines(ctx, prescriptionID)
	if err != nil {
		return err
	}

	for _, medicine := range medicines {
		available, err := uc.repo.CheckMedicineStock(ctx, medicine.MedicineID, medicine.Quantity)
		if err != nil {
			return err
		}
		if !available {
			return errors.BadRequest("INSUFFICIENT_STOCK", fmt.Sprintf("药品 %s 库存不足", medicine.MedicineName))
		}
	}

	// 更新状态
	prescription.Status = PrescriptionStatusConfirmed
	prescription.AuditorID = doctorID
	prescription.AuditTime = time.Now()
	prescription.AuditNotes = auditNotes
	prescription.UpdatedAt = time.Now()

	if err := uc.repo.UpdatePrescription(ctx, prescription); err != nil {
		uc.log.WithContext(ctx).Errorf("确认处方失败: %v", err)
		return errors.InternalServer("CONFIRM_PRESCRIPTION_FAILED", "确认处方失败")
	}

	uc.log.WithContext(ctx).Infof("处方确认成功: id=%d", prescriptionID)
	return nil
}

// 取消处方
func (uc *PrescriptionUsecase) CancelPrescription(ctx context.Context, prescriptionID, doctorID uint64, cancelReason string) error {
	// 获取处方
	prescription, err := uc.repo.GetPrescription(ctx, prescriptionID)
	if err != nil {
		return errors.NotFound("PRESCRIPTION_NOT_FOUND", "处方不存在")
	}

	// 验证权限
	if prescription.DoctorID != doctorID {
		return errors.Forbidden("PERMISSION_DENIED", "无权限操作此处方")
	}

	// 验证状态
	if prescription.Status == PrescriptionStatusCancelled {
		return errors.BadRequest("ALREADY_CANCELLED", "处方已取消")
	}
	if prescription.Status == PrescriptionStatusDispensed {
		return errors.BadRequest("CANNOT_CANCEL", "已发药的处方不能取消")
	}

	// 更新状态
	prescription.Status = PrescriptionStatusCancelled
	prescription.AuditNotes = cancelReason
	prescription.UpdatedAt = time.Now()

	if err := uc.repo.UpdatePrescription(ctx, prescription); err != nil {
		uc.log.WithContext(ctx).Errorf("取消处方失败: %v", err)
		return errors.InternalServer("CANCEL_PRESCRIPTION_FAILED", "取消处方失败")
	}

	uc.log.WithContext(ctx).Infof("处方取消成功: id=%d", prescriptionID)
	return nil
}

// 获取处方详情
func (uc *PrescriptionUsecase) GetPrescription(ctx context.Context, prescriptionID, doctorID uint64) (*Prescription, error) {
	prescription, err := uc.repo.GetPrescription(ctx, prescriptionID)
	if err != nil {
		return nil, errors.NotFound("PRESCRIPTION_NOT_FOUND", "处方不存在")
	}

	// 验证权限
	if prescription.DoctorID != doctorID {
		return nil, errors.Forbidden("PERMISSION_DENIED", "无权限查看此处方")
	}

	// 获取处方药品
	medicines, err := uc.repo.GetPrescriptionMedicines(ctx, prescriptionID)
	if err != nil {
		return nil, err
	}
	prescription.Medicines = medicines

	return prescription, nil
}

// 获取处方列表
func (uc *PrescriptionUsecase) ListPrescriptions(ctx context.Context, doctorID, patientID uint64, status string, startDate, endDate time.Time, page, pageSize int) ([]*Prescription, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	prescriptions, total, err := uc.repo.ListPrescriptions(ctx, doctorID, patientID, status, startDate, endDate, page, pageSize)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("获取处方列表失败: %v", err)
		return nil, 0, errors.InternalServer("LIST_PRESCRIPTIONS_FAILED", "获取处方列表失败")
	}

	// 获取每个处方的药品信息
	for _, prescription := range prescriptions {
		medicines, err := uc.repo.GetPrescriptionMedicines(ctx, prescription.ID)
		if err != nil {
			uc.log.WithContext(ctx).Warnf("获取处方药品失败: prescription_id=%d, error=%v", prescription.ID, err)
			continue
		}
		prescription.Medicines = medicines
	}

	return prescriptions, total, nil
}

// 搜索药品
func (uc *PrescriptionUsecase) SearchMedicines(ctx context.Context, keyword, category, prescriptionType string, onlyAvailable bool, limit int) ([]*Medicine, error) {
	if limit <= 0 || limit > 100 {
		limit = 50
	}

	medicines, err := uc.repo.SearchMedicines(ctx, keyword, category, prescriptionType, onlyAvailable, limit)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("搜索药品失败: %v", err)
		return nil, errors.InternalServer("SEARCH_MEDICINES_FAILED", "搜索药品失败")
	}

	return medicines, nil
}

// 检查药物相互作用
func (uc *PrescriptionUsecase) CheckDrugInteractions(ctx context.Context, medicineIDs []uint64) ([]*DrugInteraction, bool, error) {
	if len(medicineIDs) < 2 {
		return nil, false, nil
	}

	interactions, err := uc.repo.CheckDrugInteractions(ctx, medicineIDs)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("检查药物相互作用失败: %v", err)
		return nil, false, errors.InternalServer("CHECK_DRUG_INTERACTIONS_FAILED", "检查药物相互作用失败")
	}

	// 检查是否有严重相互作用
	hasSevere := false
	for _, interaction := range interactions {
		if interaction.Severity == "severe" {
			hasSevere = true
			break
		}
	}

	return interactions, hasSevere, nil
}

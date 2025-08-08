package data

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"doctors/internal/biz"
	"doctors/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// PrescriptionData 处方数据访问实现
type PrescriptionData struct {
	data   *Data
	logger *log.Helper
}

// NewPrescriptionData 创建处方数据访问实例
func NewPrescriptionData(data *Data, logger log.Logger) biz.PrescriptionRepo {
	return &PrescriptionData{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

// CreatePrescription 创建处方
func (d *PrescriptionData) CreatePrescription(ctx context.Context, prescription *biz.Prescription) error {
	prescriptionModel := &model.Prescriptions{
		PrescriptionNo:   prescription.PrescriptionNo,
		DoctorId:         prescription.DoctorID,
		PatientId:        prescription.PatientID,
		MedicalRecordId:  prescription.MedicalRecordID,
		PrescriptionDate: prescription.PrescriptionDate,
		TotalAmount:      prescription.TotalAmount,
		PrescriptionType: prescription.PrescriptionType,
		UsageInstruction: prescription.UsageInstruction,
		Status:           prescription.Status,
		AuditorId:        prescription.AuditorID,
		AuditTime:        prescription.AuditTime,
		AuditNotes:       prescription.AuditNotes,
		CreatedAt:        prescription.CreatedAt,
		UpdatedAt:        prescription.UpdatedAt,
	}

	err := d.data.db.WithContext(ctx).Create(prescriptionModel).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("创建处方失败: %v", err)
		return fmt.Errorf("创建处方失败: %w", err)
	}

	prescription.ID = prescriptionModel.Id
	return nil
}

// UpdatePrescription 更新处方
func (d *PrescriptionData) UpdatePrescription(ctx context.Context, prescription *biz.Prescription) error {
	prescriptionModel := &model.Prescriptions{
		Id:               prescription.ID,
		PrescriptionNo:   prescription.PrescriptionNo,
		DoctorId:         prescription.DoctorID,
		PatientId:        prescription.PatientID,
		MedicalRecordId:  prescription.MedicalRecordID,
		PrescriptionDate: prescription.PrescriptionDate,
		TotalAmount:      prescription.TotalAmount,
		PrescriptionType: prescription.PrescriptionType,
		UsageInstruction: prescription.UsageInstruction,
		Status:           prescription.Status,
		AuditorId:        prescription.AuditorID,
		AuditTime:        prescription.AuditTime,
		AuditNotes:       prescription.AuditNotes,
		UpdatedAt:        time.Now(),
	}

	err := d.data.db.WithContext(ctx).Save(prescriptionModel).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("更新处方失败: %v", err)
		return fmt.Errorf("更新处方失败: %w", err)
	}

	return nil
}

// GetPrescription 获取处方
func (d *PrescriptionData) GetPrescription(ctx context.Context, id uint64) (*biz.Prescription, error) {
	var prescriptionModel model.Prescriptions
	err := d.data.db.WithContext(ctx).Where("id = ?", id).First(&prescriptionModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("处方不存在")
		}
		d.logger.WithContext(ctx).Errorf("获取处方失败: %v", err)
		return nil, fmt.Errorf("获取处方失败: %w", err)
	}

	prescription := &biz.Prescription{
		ID:               prescriptionModel.Id,
		PrescriptionNo:   prescriptionModel.PrescriptionNo,
		DoctorID:         prescriptionModel.DoctorId,
		PatientID:        prescriptionModel.PatientId,
		MedicalRecordID:  prescriptionModel.MedicalRecordId,
		PrescriptionDate: prescriptionModel.PrescriptionDate,
		TotalAmount:      prescriptionModel.TotalAmount,
		PrescriptionType: prescriptionModel.PrescriptionType,
		UsageInstruction: prescriptionModel.UsageInstruction,
		Status:           prescriptionModel.Status,
		AuditorID:        prescriptionModel.AuditorId,
		AuditTime:        prescriptionModel.AuditTime,
		AuditNotes:       prescriptionModel.AuditNotes,
		CreatedAt:        prescriptionModel.CreatedAt,
		UpdatedAt:        prescriptionModel.UpdatedAt,
	}

	return prescription, nil
}

// GetPrescriptionByNo 根据处方号获取处方
func (d *PrescriptionData) GetPrescriptionByNo(ctx context.Context, prescriptionNo string) (*biz.Prescription, error) {
	var prescriptionModel model.Prescriptions
	err := d.data.db.WithContext(ctx).Where("prescription_no = ?", prescriptionNo).First(&prescriptionModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("处方不存在")
		}
		d.logger.WithContext(ctx).Errorf("获取处方失败: %v", err)
		return nil, fmt.Errorf("获取处方失败: %w", err)
	}

	prescription := &biz.Prescription{
		ID:               prescriptionModel.Id,
		PrescriptionNo:   prescriptionModel.PrescriptionNo,
		DoctorID:         prescriptionModel.DoctorId,
		PatientID:        prescriptionModel.PatientId,
		MedicalRecordID:  prescriptionModel.MedicalRecordId,
		PrescriptionDate: prescriptionModel.PrescriptionDate,
		TotalAmount:      prescriptionModel.TotalAmount,
		PrescriptionType: prescriptionModel.PrescriptionType,
		UsageInstruction: prescriptionModel.UsageInstruction,
		Status:           prescriptionModel.Status,
		AuditorID:        prescriptionModel.AuditorId,
		AuditTime:        prescriptionModel.AuditTime,
		AuditNotes:       prescriptionModel.AuditNotes,
		CreatedAt:        prescriptionModel.CreatedAt,
		UpdatedAt:        prescriptionModel.UpdatedAt,
	}

	return prescription, nil
}

// ListPrescriptions 获取处方列表
func (d *PrescriptionData) ListPrescriptions(ctx context.Context, doctorID, patientID uint64, status string, startDate, endDate time.Time, page, pageSize int) ([]*biz.Prescription, int64, error) {
	var prescriptionModels []model.Prescriptions
	var total int64

	query := d.data.db.WithContext(ctx).Model(&model.Prescriptions{})

	// 添加查询条件
	if doctorID > 0 {
		query = query.Where("doctor_id = ?", doctorID)
	}
	if patientID > 0 {
		query = query.Where("patient_id = ?", patientID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if !startDate.IsZero() {
		query = query.Where("prescription_date >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("prescription_date <= ?", endDate)
	}

	// 获取总数
	err := query.Count(&total).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("获取处方总数失败: %v", err)
		return nil, 0, fmt.Errorf("获取处方总数失败: %w", err)
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&prescriptionModels).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("获取处方列表失败: %v", err)
		return nil, 0, fmt.Errorf("获取处方列表失败: %w", err)
	}

	// 转换为业务实体
	prescriptions := make([]*biz.Prescription, len(prescriptionModels))
	for i, model := range prescriptionModels {
		prescriptions[i] = &biz.Prescription{
			ID:               model.Id,
			PrescriptionNo:   model.PrescriptionNo,
			DoctorID:         model.DoctorId,
			PatientID:        model.PatientId,
			MedicalRecordID:  model.MedicalRecordId,
			PrescriptionDate: model.PrescriptionDate,
			TotalAmount:      model.TotalAmount,
			PrescriptionType: model.PrescriptionType,
			UsageInstruction: model.UsageInstruction,
			Status:           model.Status,
			AuditorID:        model.AuditorId,
			AuditTime:        model.AuditTime,
			AuditNotes:       model.AuditNotes,
			CreatedAt:        model.CreatedAt,
			UpdatedAt:        model.UpdatedAt,
		}
	}

	return prescriptions, total, nil
}

// DeletePrescription 删除处方
func (d *PrescriptionData) DeletePrescription(ctx context.Context, id uint64) error {
	err := d.data.db.WithContext(ctx).Delete(&model.Prescriptions{}, id).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("删除处方失败: %v", err)
		return fmt.Errorf("删除处方失败: %w", err)
	}
	return nil
}

// CreatePrescriptionMedicines 创建处方药品
func (d *PrescriptionData) CreatePrescriptionMedicines(ctx context.Context, medicines []biz.PrescriptionMedicine) error {
	medicineModels := make([]model.PrescriptionMedicines, len(medicines))
	for i, medicine := range medicines {
		medicineModels[i] = model.PrescriptionMedicines{
			PrescriptionId: medicine.PrescriptionID,
			MedicineId:     medicine.MedicineID,
			Quantity:       medicine.Quantity,
			Unit:           medicine.Unit,
			UnitPrice:      medicine.UnitPrice,
			TotalPrice:     medicine.TotalPrice,
			Dosage:         medicine.Dosage,
			Frequency:      medicine.Frequency,
			Duration:       medicine.Duration,
			UsageMethod:    medicine.UsageMethod,
			Notes:          medicine.Notes,
			CreatedAt:      time.Now(),
		}
	}

	err := d.data.db.WithContext(ctx).Create(&medicineModels).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("创建处方药品失败: %v", err)
		return fmt.Errorf("创建处方药品失败: %w", err)
	}

	// 更新ID
	for i := range medicines {
		medicines[i].ID = medicineModels[i].Id
	}

	return nil
}

// UpdatePrescriptionMedicines 更新处方药品
func (d *PrescriptionData) UpdatePrescriptionMedicines(ctx context.Context, prescriptionID uint64, medicines []biz.PrescriptionMedicine) error {
	// 先删除原有的处方药品
	err := d.data.db.WithContext(ctx).Where("prescription_id = ?", prescriptionID).Delete(&model.PrescriptionMedicines{}).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("删除原有处方药品失败: %v", err)
		return fmt.Errorf("删除原有处方药品失败: %w", err)
	}

	// 创建新的处方药品
	if len(medicines) > 0 {
		return d.CreatePrescriptionMedicines(ctx, medicines)
	}

	return nil
}

// GetPrescriptionMedicines 获取处方药品
func (d *PrescriptionData) GetPrescriptionMedicines(ctx context.Context, prescriptionID uint64) ([]biz.PrescriptionMedicine, error) {
	var medicineModels []model.PrescriptionMedicines
	err := d.data.db.WithContext(ctx).Where("prescription_id = ?", prescriptionID).Find(&medicineModels).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("获取处方药品失败: %v", err)
		return nil, fmt.Errorf("获取处方药品失败: %w", err)
	}

	medicines := make([]biz.PrescriptionMedicine, len(medicineModels))
	for i, model := range medicineModels {
		medicines[i] = biz.PrescriptionMedicine{
			ID:             model.Id,
			PrescriptionID: model.PrescriptionId,
			MedicineID:     model.MedicineId,
			Quantity:       model.Quantity,
			Unit:           model.Unit,
			UnitPrice:      model.UnitPrice,
			TotalPrice:     model.TotalPrice,
			Dosage:         model.Dosage,
			Frequency:      model.Frequency,
			Duration:       model.Duration,
			UsageMethod:    model.UsageMethod,
			Notes:          model.Notes,
		}
	}

	return medicines, nil
}

// DeletePrescriptionMedicines 删除处方药品
func (d *PrescriptionData) DeletePrescriptionMedicines(ctx context.Context, prescriptionID uint64) error {
	err := d.data.db.WithContext(ctx).Where("prescription_id = ?", prescriptionID).Delete(&model.PrescriptionMedicines{}).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("删除处方药品失败: %v", err)
		return fmt.Errorf("删除处方药品失败: %w", err)
	}
	return nil
}

// SearchMedicines 搜索药品
func (d *PrescriptionData) SearchMedicines(ctx context.Context, keyword, category, prescriptionType string, onlyAvailable bool, limit int) ([]*biz.Medicine, error) {
	var medicineModels []model.Medicines
	query := d.data.db.WithContext(ctx).Model(&model.Medicines{}).Where("status = ?", "启用")

	// 添加搜索条件
	if keyword != "" {
		query = query.Where("name LIKE ? OR generic_name LIKE ? OR medicine_code LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if prescriptionType != "" {
		query = query.Where("prescription_type = ?", prescriptionType)
	}

	// 如果只显示有库存的药品，需要关联库存表
	if onlyAvailable {
		query = query.Joins("JOIN medicine_inventory ON medicines.id = medicine_inventory.medicine_id").
			Where("medicine_inventory.available_quantity > 0")
	}

	err := query.Limit(limit).Find(&medicineModels).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("搜索药品失败: %v", err)
		return nil, fmt.Errorf("搜索药品失败: %w", err)
	}

	medicines := make([]*biz.Medicine, len(medicineModels))
	for i, model := range medicineModels {
		medicines[i] = &biz.Medicine{
			ID:                model.Id,
			MedicineCode:      model.MedicineCode,
			Name:              model.Name,
			GenericName:       model.GenericName,
			BrandName:         model.BrandName,
			Specification:     model.Specification,
			DosageForm:        model.DosageForm,
			Manufacturer:      model.Manufacturer,
			ApprovalNumber:    model.ApprovalNumber,
			Category:          model.Category,
			PrescriptionType:  model.PrescriptionType,
			Unit:              model.Unit,
			Price:             model.Price,
			Indications:       model.Indications,
			Contraindications: model.Contraindications,
			SideEffects:       model.SideEffects,
			DosageUsage:       model.DosageUsage,
			StorageConditions: model.StorageConditions,
			ShelfLife:         model.ShelfLife,
			ImageURL:          model.ImageUrl,
			Status:            model.Status,
		}
	}

	return medicines, nil
}

// GetMedicine 获取药品信息
func (d *PrescriptionData) GetMedicine(ctx context.Context, id uint64) (*biz.Medicine, error) {
	var medicineModel model.Medicines
	err := d.data.db.WithContext(ctx).Where("id = ?", id).First(&medicineModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("药品不存在")
		}
		d.logger.WithContext(ctx).Errorf("获取药品失败: %v", err)
		return nil, fmt.Errorf("获取药品失败: %w", err)
	}

	medicine := &biz.Medicine{
		ID:                medicineModel.Id,
		MedicineCode:      medicineModel.MedicineCode,
		Name:              medicineModel.Name,
		GenericName:       medicineModel.GenericName,
		BrandName:         medicineModel.BrandName,
		Specification:     medicineModel.Specification,
		DosageForm:        medicineModel.DosageForm,
		Manufacturer:      medicineModel.Manufacturer,
		ApprovalNumber:    medicineModel.ApprovalNumber,
		Category:          medicineModel.Category,
		PrescriptionType:  medicineModel.PrescriptionType,
		Unit:              medicineModel.Unit,
		Price:             medicineModel.Price,
		Indications:       medicineModel.Indications,
		Contraindications: medicineModel.Contraindications,
		SideEffects:       medicineModel.SideEffects,
		DosageUsage:       medicineModel.DosageUsage,
		StorageConditions: medicineModel.StorageConditions,
		ShelfLife:         medicineModel.ShelfLife,
		ImageURL:          medicineModel.ImageUrl,
		Status:            medicineModel.Status,
	}

	return medicine, nil
}

// CheckMedicineStock 检查药品库存
func (d *PrescriptionData) CheckMedicineStock(ctx context.Context, medicineID uint64, quantity float64) (bool, error) {
	var inventory model.MedicineInventory
	err := d.data.db.WithContext(ctx).Where("medicine_id = ? AND status = ?", medicineID, "正常").First(&inventory).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		d.logger.WithContext(ctx).Errorf("检查药品库存失败: %v", err)
		return false, fmt.Errorf("检查药品库存失败: %w", err)
	}

	return inventory.AvailableQuantity >= quantity, nil
}

// GetMedicinesByIDs 根据ID列表获取药品
func (d *PrescriptionData) GetMedicinesByIDs(ctx context.Context, ids []uint64) ([]*biz.Medicine, error) {
	var medicineModels []model.Medicines
	err := d.data.db.WithContext(ctx).Where("id IN ?", ids).Find(&medicineModels).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("获取药品列表失败: %v", err)
		return nil, fmt.Errorf("获取药品列表失败: %w", err)
	}

	medicines := make([]*biz.Medicine, len(medicineModels))
	for i, model := range medicineModels {
		medicines[i] = &biz.Medicine{
			ID:                model.Id,
			MedicineCode:      model.MedicineCode,
			Name:              model.Name,
			GenericName:       model.GenericName,
			BrandName:         model.BrandName,
			Specification:     model.Specification,
			DosageForm:        model.DosageForm,
			Manufacturer:      model.Manufacturer,
			ApprovalNumber:    model.ApprovalNumber,
			Category:          model.Category,
			PrescriptionType:  model.PrescriptionType,
			Unit:              model.Unit,
			Price:             model.Price,
			Indications:       model.Indications,
			Contraindications: model.Contraindications,
			SideEffects:       model.SideEffects,
			DosageUsage:       model.DosageUsage,
			StorageConditions: model.StorageConditions,
			ShelfLife:         model.ShelfLife,
			ImageURL:          model.ImageUrl,
			Status:            model.Status,
		}
	}

	return medicines, nil
}

// CheckDrugInteractions 检查药物相互作用
func (d *PrescriptionData) CheckDrugInteractions(ctx context.Context, medicineIDs []uint64) ([]*biz.DrugInteraction, error) {
	// 这里需要根据你的药物相互作用表结构来实现
	// 暂时返回空列表
	return []*biz.DrugInteraction{}, nil
}

// CachePrescriptionDraft 缓存处方草稿
func (d *PrescriptionData) CachePrescriptionDraft(ctx context.Context, doctorID uint64, prescription *biz.Prescription) error {
	key := fmt.Sprintf("prescription_draft:%d", doctorID)
	data, err := json.Marshal(prescription)
	if err != nil {
		return err
	}

	err = d.data.rdb.Set(ctx, key, data, time.Hour).Err()
	if err != nil {
		d.logger.WithContext(ctx).Errorf("缓存处方草稿失败: %v", err)
		return fmt.Errorf("缓存处方草稿失败: %w", err)
	}

	return nil
}

// GetPrescriptionDraft 获取处方草稿
func (d *PrescriptionData) GetPrescriptionDraft(ctx context.Context, doctorID uint64) (*biz.Prescription, error) {
	key := fmt.Sprintf("prescription_draft:%d", doctorID)
	data, err := d.data.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		d.logger.WithContext(ctx).Errorf("获取处方草稿失败: %v", err)
		return nil, fmt.Errorf("获取处方草稿失败: %w", err)
	}

	var prescription biz.Prescription
	err = json.Unmarshal([]byte(data), &prescription)
	if err != nil {
		return nil, err
	}

	return &prescription, nil
}

// DeletePrescriptionDraft 删除处方草稿
func (d *PrescriptionData) DeletePrescriptionDraft(ctx context.Context, doctorID uint64) error {
	key := fmt.Sprintf("prescription_draft:%d", doctorID)
	err := d.data.rdb.Del(ctx, key).Err()
	if err != nil {
		d.logger.WithContext(ctx).Errorf("删除处方草稿失败: %v", err)
		return fmt.Errorf("删除处方草稿失败: %w", err)
	}

	return nil
}

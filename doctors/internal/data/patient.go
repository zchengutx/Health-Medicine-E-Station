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

// PatientData 患者数据访问实现
type PatientData struct {
	data   *Data
	logger *log.Helper
}

// NewPatientData 创建患者数据访问实例
func NewPatientData(data *Data, logger log.Logger) biz.PatientRepo {
	return &PatientData{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

// CreatePatient 创建患者记录
func (d *PatientData) CreatePatient(ctx context.Context, patient *biz.Patient) error {
	model := &model.Patients{
		PatientCode:      patient.PatientCode,
		Name:             patient.Name,
		Gender:           patient.Gender,
		Age:              int32(patient.Age),
		Phone:            patient.Phone,
		IdCard:           patient.IdCard,
		Address:          patient.Address,
		EmergencyContact: patient.EmergencyContact,
		EmergencyPhone:   patient.EmergencyPhone,
		MedicalHistory:   patient.MedicalHistory,
		AllergyHistory:   patient.Allergies,
		Status:           patient.Status,
	}

	if patient.BirthDate != nil {
		model.BirthDate = *patient.BirthDate
	}

	err := d.data.db.WithContext(ctx).Create(model).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("创建患者记录失败: name=%s, error=%v", patient.Name, err)
		return fmt.Errorf("创建患者记录失败: %w", err)
	}

	// 更新ID到业务实体
	patient.ID = uint(model.Id)
	patient.PatientCode = model.PatientCode
	patient.CreatedAt = model.CreatedAt
	patient.UpdatedAt = model.UpdatedAt

	d.logger.WithContext(ctx).Infof("患者记录创建成功: id=%d, name=%s", model.Id, patient.Name)
	return nil
}

// GetPatientByID 根据ID获取患者信息
func (d *PatientData) GetPatientByID(ctx context.Context, id uint) (*biz.Patient, error) {
	var model model.Patients

	err := d.data.db.WithContext(ctx).Where("id = ?", id).First(&model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("患者不存在")
		}
		d.logger.WithContext(ctx).Errorf("查询患者失败: id=%d, error=%v", id, err)
		return nil, fmt.Errorf("查询患者失败: %w", err)
	}

	patient := d.modelToEntity(&model)
	return patient, nil
}

// GetPatientByIdCard 根据身份证号获取患者信息
func (d *PatientData) GetPatientByIdCard(ctx context.Context, idCard string) (*biz.Patient, error) {
	var model model.Patients

	err := d.data.db.WithContext(ctx).Where("id_card = ?", idCard).First(&model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("患者不存在")
		}
		d.logger.WithContext(ctx).Errorf("查询患者失败: id_card=%s, error=%v", idCard, err)
		return nil, fmt.Errorf("查询患者失败: %w", err)
	}

	patient := d.modelToEntity(&model)
	return patient, nil
}

// UpdatePatient 更新患者信息
func (d *PatientData) UpdatePatient(ctx context.Context, patient *biz.Patient) error {
	patientModel := &model.Patients{
		Id:               uint64(patient.ID),
		Name:             patient.Name,
		Gender:           patient.Gender,
		Age:              int32(patient.Age),
		Phone:            patient.Phone,
		Address:          patient.Address,
		EmergencyContact: patient.EmergencyContact,
		EmergencyPhone:   patient.EmergencyPhone,
		MedicalHistory:   patient.MedicalHistory,
		AllergyHistory:   patient.Allergies,
		Status:           patient.Status,
		UpdatedAt:        time.Now(),
	}

	if patient.BirthDate != nil {
		patientModel.BirthDate = *patient.BirthDate
	}

	err := d.data.db.WithContext(ctx).Save(patientModel).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("更新患者信息失败: id=%d, error=%v", patient.ID, err)
		return fmt.Errorf("更新患者信息失败: %w", err)
	}

	patient.UpdatedAt = patientModel.UpdatedAt
	d.logger.WithContext(ctx).Infof("患者信息更新成功: id=%d", patient.ID)
	return nil
}

// DeletePatient 删除患者
func (d *PatientData) DeletePatient(ctx context.Context, id uint) error {
	err := d.data.db.WithContext(ctx).Delete(&model.Patients{}, id).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("删除患者失败: id=%d, error=%v", id, err)
		return fmt.Errorf("删除患者失败: %w", err)
	}

	d.logger.WithContext(ctx).Infof("患者删除成功: id=%d", id)
	return nil
}

// GetPatientList 获取患者列表
func (d *PatientData) GetPatientList(ctx context.Context, query *biz.PatientListQuery) (*biz.PatientListResult, error) {
	var models []model.Patients
	var total int64

	db := d.data.db.WithContext(ctx).Model(&model.Patients{})

	// 添加搜索条件
	if query.Keyword != "" {
		db = db.Where("name LIKE ? OR phone LIKE ? OR id_card LIKE ?",
			"%"+query.Keyword+"%", "%"+query.Keyword+"%", "%"+query.Keyword+"%")
	}

	// 添加分类筛选 - 注释掉因为Patients模型中没有Category字段
	// if query.Category != "" {
	//	db = db.Where("category = ?", query.Category)
	// }

	// 添加状态筛选
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	// 获取总数
	err := db.Count(&total).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("查询患者总数失败: error=%v", err)
		return nil, fmt.Errorf("查询患者总数失败: %w", err)
	}

	// 分页查询
	offset := (query.Page - 1) * query.PageSize
	err = db.Offset(offset).Limit(query.PageSize).Order("created_at DESC").Find(&models).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("查询患者列表失败: error=%v", err)
		return nil, fmt.Errorf("查询患者列表失败: %w", err)
	}

	// 转换为业务实体
	patients := make([]*biz.Patient, len(models))
	for i, model := range models {
		patients[i] = d.modelToEntity(&model)
	}

	result := &biz.PatientListResult{
		Patients: patients,
		Total:    total,
		Page:     query.Page,
		PageSize: query.PageSize,
	}

	return result, nil
}

// GetPatientsByCategory 按分类获取患者
func (d *PatientData) GetPatientsByCategory(ctx context.Context, category string, page, pageSize int) (*biz.PatientListResult, error) {
	var models []model.Patients
	var total int64

	// 注释掉因为Patients模型中没有category字段
	db := d.data.db.WithContext(ctx).Model(&model.Patients{}) // .Where("category = ?", category)

	// 获取总数
	err := db.Count(&total).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("查询分类患者总数失败: category=%s, error=%v", category, err)
		return nil, fmt.Errorf("查询分类患者总数失败: %w", err)
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = db.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&models).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("查询分类患者列表失败: category=%s, error=%v", category, err)
		return nil, fmt.Errorf("查询分类患者列表失败: %w", err)
	}

	// 转换为业务实体
	patients := make([]*biz.Patient, len(models))
	for i, model := range models {
		patients[i] = d.modelToEntity(&model)
	}

	result := &biz.PatientListResult{
		Patients: patients,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}

	return result, nil
}

// SearchPatients 搜索患者
func (d *PatientData) SearchPatients(ctx context.Context, keyword string, page, pageSize int) (*biz.PatientListResult, error) {
	query := &biz.PatientListQuery{
		Page:     page,
		PageSize: pageSize,
		Keyword:  keyword,
	}
	return d.GetPatientList(ctx, query)
}

// modelToEntity 将数据模型转换为业务实体
func (d *PatientData) modelToEntity(model *model.Patients) *biz.Patient {
	return &biz.Patient{
		ID:               uint(model.Id),
		PatientCode:      model.PatientCode,
		Name:             model.Name,
		Gender:           model.Gender,
		BirthDate:        &model.BirthDate,
		Age:              int(model.Age),
		Phone:            model.Phone,
		IdCard:           model.IdCard,
		Address:          model.Address,
		EmergencyContact: model.EmergencyContact,
		EmergencyPhone:   model.EmergencyPhone,
		MedicalHistory:   model.MedicalHistory,
		Allergies:        model.AllergyHistory,
		Category:         "", // Patients模型中没有Category字段
		Status:           model.Status,
		Remarks:          "", // Patients模型中没有Remarks字段
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
	}
}

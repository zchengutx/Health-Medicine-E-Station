package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrPatientNotFound 患者不存在
	ErrPatientNotFound = errors.NotFound("PATIENT_NOT_FOUND", "患者不存在")
	// ErrPatientAlreadyExists 患者已存在
	ErrPatientAlreadyExists = errors.Conflict("PATIENT_ALREADY_EXISTS", "患者已存在")
	// ErrInvalidPatientData 患者数据无效
	ErrInvalidPatientData = errors.BadRequest("INVALID_PATIENT_DATA", "患者数据无效")
)

// Patient 患者业务实体
type Patient struct {
	ID               uint       `json:"id"`
	PatientCode      string     `json:"patient_code"`
	Name             string     `json:"name"`
	Gender           string     `json:"gender"`
	BirthDate        *time.Time `json:"birth_date"`
	Age              int        `json:"age"`
	Phone            string     `json:"phone"`
	IdCard           string     `json:"id_card"`
	Address          string     `json:"address"`
	EmergencyContact string     `json:"emergency_contact"`
	EmergencyPhone   string     `json:"emergency_phone"`
	MedicalHistory   string     `json:"medical_history"`
	Allergies        string     `json:"allergies"`
	Category         string     `json:"category"`
	Status           string     `json:"status"`
	Remarks          string     `json:"remarks"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// PatientListQuery 患者列表查询参数
type PatientListQuery struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Keyword  string `json:"keyword"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

// PatientListResult 患者列表查询结果
type PatientListResult struct {
	Patients []*Patient `json:"patients"`
	Total    int64      `json:"total"`
	Page     int        `json:"page"`
	PageSize int        `json:"page_size"`
}

// PatientRepo 患者数据访问接口
type PatientRepo interface {
	// 基础CRUD操作
	CreatePatient(ctx context.Context, patient *Patient) error
	GetPatientByID(ctx context.Context, id uint) (*Patient, error)
	GetPatientByIdCard(ctx context.Context, idCard string) (*Patient, error)
	UpdatePatient(ctx context.Context, patient *Patient) error
	DeletePatient(ctx context.Context, id uint) error

	// 查询操作
	GetPatientList(ctx context.Context, query *PatientListQuery) (*PatientListResult, error)
	GetPatientsByCategory(ctx context.Context, category string, page, pageSize int) (*PatientListResult, error)
	SearchPatients(ctx context.Context, keyword string, page, pageSize int) (*PatientListResult, error)
}

// PatientUsecase 患者业务逻辑
type PatientUsecase struct {
	repo PatientRepo
	log  *log.Helper
}

// NewPatientUsecase 创建患者业务逻辑实例
func NewPatientUsecase(repo PatientRepo, logger log.Logger) *PatientUsecase {
	return &PatientUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreatePatient 创建患者
func (uc *PatientUsecase) CreatePatient(ctx context.Context, patient *Patient) error {
	// 验证必填字段
	if patient.Name == "" {
		return ErrInvalidPatientData
	}

	// 检查身份证号是否已存在
	if patient.IdCard != "" {
		existingPatient, err := uc.repo.GetPatientByIdCard(ctx, patient.IdCard)
		if err == nil && existingPatient != nil {
			return ErrPatientAlreadyExists
		}
	}

	// 生成患者编码
	patient.PatientCode = uc.generatePatientCode()

	// 计算年龄
	if patient.BirthDate != nil {
		patient.Age = uc.calculateAge(*patient.BirthDate)
	}

	// 设置默认状态
	if patient.Status == "" {
		patient.Status = "正常"
	}

	// 设置默认分类
	if patient.Category == "" {
		patient.Category = "普通患者"
	}

	return uc.repo.CreatePatient(ctx, patient)
}

// GetPatientByID 根据ID获取患者信息
func (uc *PatientUsecase) GetPatientByID(ctx context.Context, id uint) (*Patient, error) {
	return uc.repo.GetPatientByID(ctx, id)
}

// UpdatePatientProfile 更新患者档案
func (uc *PatientUsecase) UpdatePatientProfile(ctx context.Context, patient *Patient) error {
	// 检查患者是否存在
	existingPatient, err := uc.repo.GetPatientByID(ctx, patient.ID)
	if err != nil {
		return ErrPatientNotFound
	}

	// 更新允许修改的字段
	existingPatient.Name = patient.Name
	existingPatient.Gender = patient.Gender
	existingPatient.BirthDate = patient.BirthDate
	existingPatient.Phone = patient.Phone
	existingPatient.Address = patient.Address
	existingPatient.EmergencyContact = patient.EmergencyContact
	existingPatient.EmergencyPhone = patient.EmergencyPhone
	existingPatient.MedicalHistory = patient.MedicalHistory
	existingPatient.Allergies = patient.Allergies
	existingPatient.Remarks = patient.Remarks

	// 重新计算年龄
	if existingPatient.BirthDate != nil {
		existingPatient.Age = uc.calculateAge(*existingPatient.BirthDate)
	}

	return uc.repo.UpdatePatient(ctx, existingPatient)
}

// GetPatientList 获取患者列表
func (uc *PatientUsecase) GetPatientList(ctx context.Context, query *PatientListQuery) (*PatientListResult, error) {
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

	return uc.repo.GetPatientList(ctx, query)
}

// GetPatientsByCategory 按分类获取患者
func (uc *PatientUsecase) GetPatientsByCategory(ctx context.Context, category string, page, pageSize int) (*PatientListResult, error) {
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

	return uc.repo.GetPatientsByCategory(ctx, category, page, pageSize)
}

// UpdatePatientCategory 更新患者分类
func (uc *PatientUsecase) UpdatePatientCategory(ctx context.Context, patientID uint, category string) error {
	// 检查患者是否存在
	patient, err := uc.repo.GetPatientByID(ctx, patientID)
	if err != nil {
		return ErrPatientNotFound
	}

	// 更新分类
	patient.Category = category
	return uc.repo.UpdatePatient(ctx, patient)
}

// generatePatientCode 生成患者编码
func (uc *PatientUsecase) generatePatientCode() string {
	now := time.Now()
	return fmt.Sprintf("P%s%06d", now.Format("20060102"), now.Unix()%1000000)
}

// calculateAge 计算年龄
func (uc *PatientUsecase) calculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()

	// 如果还没到生日，年龄减1
	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	if age < 0 {
		age = 0
	}

	return age
}

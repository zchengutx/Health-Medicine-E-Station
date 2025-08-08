package biz

import (
	"context"
	"doctors/utils"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrDoctorNotFound 医生不存在
	ErrDoctorNotFound = errors.NotFound("DOCTOR_NOT_FOUND", "医生不存在")
	// ErrPhoneAlreadyExists 手机号已存在
	ErrPhoneAlreadyExists = errors.Conflict("PHONE_ALREADY_EXISTS", "手机号已注册")
	// ErrInvalidPassword 密码错误
	ErrInvalidPassword = errors.Unauthorized("INVALID_PASSWORD", "密码错误")
	// ErrInvalidVerificationCode 验证码错误
	ErrInvalidVerificationCode = errors.BadRequest("INVALID_VERIFICATION_CODE", "验证码错误")
)

// Doctor 医生业务实体
type Doctor struct {
	ID            uint       `json:"id"`
	DoctorCode    string     `json:"doctor_code"`
	Phone         string     `json:"phone"`
	Password      string     `json:"password"`
	Name          string     `json:"name"`
	Gender        string     `json:"gender"`
	BirthDate     *time.Time `json:"birth_date"`
	Email         string     `json:"email"`
	Avatar        string     `json:"avatar"`
	LicenseNumber string     `json:"license_number"`
	DepartmentID  *uint      `json:"department_id"`
	HospitalID    *uint      `json:"hospital_id"`
	Title         string     `json:"title"`
	Speciality    string     `json:"speciality"`
	PracticeScope string     `json:"practice_scope"`
	Status        string     `json:"status"`
	LastLoginTime *time.Time `json:"last_login_time"`
	LastLoginIP   string     `json:"last_login_ip"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// DoctorRepo 医生数据访问接口
type DoctorRepo interface {
	// Redis相关操作
	SaveVerificationCode(ctx context.Context, phone, code string, expireTime time.Duration) error
	GetVerificationCode(ctx context.Context, phone string) (string, error)
	DeleteVerificationCode(ctx context.Context, phone string) error

	// 数据库相关操作
	CreateDoctor(ctx context.Context, doctor *Doctor) error
	GetDoctorByPhone(ctx context.Context, phone string) (*Doctor, error)
	GetDoctorByID(ctx context.Context, id uint) (*Doctor, error)
	UpdateDoctor(ctx context.Context, doctor *Doctor) error
	DeleteDoctorByID(ctx context.Context, id uint) error
}

// DoctorUsecase 医生业务逻辑
type DoctorUsecase struct {
	repo DoctorRepo
	log  *log.Helper
}

// NewDoctorUsecase 创建医生业务逻辑实例
func NewDoctorUsecase(repo DoctorRepo, logger log.Logger) *DoctorUsecase {
	return &DoctorUsecase{repo: repo, log: log.NewHelper(logger)}
}

// SendSmsCode 发送短信验证码
func (uc *DoctorUsecase) SendSmsCode(ctx context.Context, phone string) error {
	// 生成验证码逻辑
	code := rand.Intn(90000) + 10000

	// 保存验证码到Redis
	err := uc.repo.SaveVerificationCode(ctx, phone, strconv.Itoa(code), 5*time.Minute)
	if err != nil {
		return err
	}

	// 实际发送短信的逻辑
	uc.log.WithContext(ctx).Infof("向手机号 %s 发送验证码: %s", phone, code)

	return nil
}

// VerifySmsCode 验证短信验证码
func (uc *DoctorUsecase) VerifySmsCode(ctx context.Context, phone, code string) (bool, error) {
	savedCode, err := uc.repo.GetVerificationCode(ctx, phone)
	if err != nil {
		return false, err
	}

	if savedCode != code {
		return false, nil
	}

	// 验证成功后删除验证码
	_ = uc.repo.DeleteVerificationCode(ctx, phone)

	return true, nil
}

// generateDoctorCode 生成唯一的医生编码
func (uc *DoctorUsecase) generateDoctorCode() string {
	// 使用时间戳和随机数生成唯一编码
	timestamp := time.Now().Unix()
	randomNum := rand.Intn(9999) + 1000
	return fmt.Sprintf("DR%d%04d", timestamp, randomNum)
}

// RegisterDoctor 注册医生
func (uc *DoctorUsecase) RegisterDoctor(ctx context.Context, doctor *Doctor) error {
	// 检查手机号是否已注册
	existingDoctor, err := uc.repo.GetDoctorByPhone(ctx, doctor.Phone)
	if err == nil && existingDoctor != nil {
		return ErrPhoneAlreadyExists
	}

	// 对密码进行加密
	hashedPassword, err := utils.HashPassword(doctor.Password, 12) // 使用较高的加密强度
	if err != nil {
		uc.log.WithContext(ctx).Errorf("密码加密失败: %v", err)
		return err
	}
	doctor.Password = hashedPassword

	// 生成唯一的医生编码
	doctor.DoctorCode = uc.generateDoctorCode()

	// 创建医生记录
	err = uc.repo.CreateDoctor(ctx, doctor)
	if err != nil {
		return err
	}

	uc.log.WithContext(ctx).Infof("医生注册成功: phone=%s, doctor_code=%s", doctor.Phone, doctor.DoctorCode)
	return nil
}

// LoginDoctor 医生登录
func (uc *DoctorUsecase) LoginDoctor(ctx context.Context, phone, password string) (*Doctor, error) {
	// 根据手机号获取医生信息
	doctor, err := uc.repo.GetDoctorByPhone(ctx, phone)
	if err != nil {
		return nil, ErrDoctorNotFound
	}

	// 验证密码
	isValid, err := utils.VerifyPassword(password, doctor.Password)
	if err != nil || !isValid {
		return nil, ErrInvalidPassword
	}

	// 更新登录信息
	now := time.Now()
	doctor.LastLoginTime = &now
	// doctor.LastLoginIP = 获取IP地址

	// 更新医生信息
	err = uc.repo.UpdateDoctor(ctx, doctor)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("更新医生登录信息失败: %v", err)
	}

	return doctor, nil
}

// GetDoctorByID 根据ID获取医生信息
func (uc *DoctorUsecase) GetDoctorByID(ctx context.Context, id uint) (*Doctor, error) {
	return uc.repo.GetDoctorByID(ctx, id)
}

// GetDoctorByPhone 根据手机号获取医生信息
func (uc *DoctorUsecase) GetDoctorByPhone(ctx context.Context, phone string) (*Doctor, error) {
	return uc.repo.GetDoctorByPhone(ctx, phone)
}

// UpdateDoctorInfo 更新医生信息
func (uc *DoctorUsecase) UpdateDoctorInfo(ctx context.Context, doctor *Doctor) error {
	return uc.repo.UpdateDoctor(ctx, doctor)
}

// UpdateDoctorProfile 更新医生个人资料
func (uc *DoctorUsecase) UpdateDoctorProfile(ctx context.Context, doctor *Doctor) error {
	// 检查医生是否存在
	existingDoctor, err := uc.repo.GetDoctorByID(ctx, doctor.ID)
	if err != nil {
		return ErrDoctorNotFound
	}

	// 更新允许修改的字段
	existingDoctor.Name = doctor.Name
	existingDoctor.Gender = doctor.Gender
	existingDoctor.BirthDate = doctor.BirthDate
	existingDoctor.Email = doctor.Email
	if doctor.Avatar != "" {
		existingDoctor.Avatar = doctor.Avatar
	}
	existingDoctor.Title = doctor.Title
	existingDoctor.Speciality = doctor.Speciality
	existingDoctor.PracticeScope = doctor.PracticeScope

	return uc.repo.UpdateDoctor(ctx, existingDoctor)
}

// ChangePassword 修改密码
func (uc *DoctorUsecase) ChangePassword(ctx context.Context, doctorID uint, oldPassword, newPassword string) error {
	// 获取医生信息
	doctor, err := uc.repo.GetDoctorByID(ctx, doctorID)
	if err != nil {
		return ErrDoctorNotFound
	}

	// 验证旧密码
	isValid, err := utils.VerifyPassword(oldPassword, doctor.Password)
	if err != nil || !isValid {
		return ErrInvalidPassword
	}

	// 加密新密码
	hashedPassword, err := utils.HashPassword(newPassword, 12)
	if err != nil {
		return err
	}

	// 更新密码
	doctor.Password = hashedPassword
	return uc.repo.UpdateDoctor(ctx, doctor)
}

// DeleteDoctorByID 注销医生账号
func (uc *DoctorUsecase) DeleteDoctorByID(ctx context.Context, id uint) error {
	// 检查医生是否存在
	doctor, err := uc.repo.GetDoctorByID(ctx, id)
	if err != nil || doctor == nil {
		return ErrDoctorNotFound
	}
	// 删除医生账号
	return uc.repo.DeleteDoctorByID(ctx, id)
}

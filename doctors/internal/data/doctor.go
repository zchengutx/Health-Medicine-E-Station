package data

import (
	"context"
	"fmt"
	"sync"
	"time"

	"doctors/internal/biz"
	"doctors/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 内存存储验证码的结构
type smsCodeData struct {
	code     string
	expireAt time.Time
}

// 内存存储验证码的map
var (
	smsCodeStore = make(map[string]*smsCodeData)
	smsCodeMutex = sync.RWMutex{}
)

// DoctorData 医生数据访问实现
type DoctorData struct {
	data   *Data
	logger *log.Helper
}

// NewDoctorData 创建医生数据访问实例
func NewDoctorData(data *Data, logger log.Logger) biz.DoctorRepo {
	return &DoctorData{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

// SaveVerificationCode 保存验证码到Redis，失败时使用内存存储
func (d *DoctorData) SaveVerificationCode(ctx context.Context, phone, code string, expireTime time.Duration) error {
	key := fmt.Sprintf("sms_code:%s", phone)

	// 尝试保存到Redis
	err := d.data.rdb.Set(ctx, key, code, expireTime).Err()
	if err != nil {
		d.logger.WithContext(ctx).Warnf("保存验证码到Redis失败，使用内存存储: phone=%s, error=%v", phone, err)

		// Redis失败时使用内存存储
		smsCodeMutex.Lock()
		smsCodeStore[phone] = &smsCodeData{
			code:     code,
			expireAt: time.Now().Add(expireTime),
		}
		smsCodeMutex.Unlock()

		d.logger.WithContext(ctx).Infof("验证码已保存到内存: phone=%s, expire=%v", phone, expireTime)
		return nil
	}

	d.logger.WithContext(ctx).Infof("验证码已保存到Redis: phone=%s, expire=%v", phone, expireTime)
	return nil
}

// GetVerificationCode 从Redis获取验证码，失败时从内存存储获取
func (d *DoctorData) GetVerificationCode(ctx context.Context, phone string) (string, error) {
	key := fmt.Sprintf("sms_code:%s", phone)

	// 尝试从Redis获取
	code, err := d.data.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			d.logger.WithContext(ctx).Infof("Redis中验证码不存在，尝试从内存获取: phone=%s", phone)
		} else {
			d.logger.WithContext(ctx).Warnf("从Redis获取验证码失败，尝试从内存获取: phone=%s, error=%v", phone, err)
		}

		// Redis失败时从内存存储获取
		smsCodeMutex.RLock()
		smsData, exists := smsCodeStore[phone]
		smsCodeMutex.RUnlock()

		if !exists {
			return "", fmt.Errorf("验证码不存在或已过期")
		}

		// 检查是否过期
		if time.Now().After(smsData.expireAt) {
			// 清理过期的验证码
			smsCodeMutex.Lock()
			delete(smsCodeStore, phone)
			smsCodeMutex.Unlock()
			return "", fmt.Errorf("验证码不存在或已过期")
		}

		return smsData.code, nil
	}

	return code, nil
}

// DeleteVerificationCode 从Redis删除验证码，同时从内存存储删除
func (d *DoctorData) DeleteVerificationCode(ctx context.Context, phone string) error {
	key := fmt.Sprintf("sms_code:%s", phone)

	// 尝试从Redis删除
	err := d.data.rdb.Del(ctx, key).Err()
	if err != nil {
		d.logger.WithContext(ctx).Warnf("从Redis删除验证码失败: phone=%s, error=%v", phone, err)
	}

	// 同时从内存存储删除
	smsCodeMutex.Lock()
	delete(smsCodeStore, phone)
	smsCodeMutex.Unlock()

	d.logger.WithContext(ctx).Infof("验证码已删除: phone=%s", phone)
	return nil
}

// CreateDoctor 创建医生记录
func (d *DoctorData) CreateDoctor(ctx context.Context, doctor *biz.Doctor) error {
	doctorModel := &model.Doctors{
		DoctorCode:    doctor.DoctorCode,
		Phone:         doctor.Phone,
		Password:      doctor.Password,
		Gender:        doctor.Gender,
		Email:         doctor.Email,
		Avatar:        doctor.Avatar,
		Title:         doctor.Title,
		Speciality:    doctor.Speciality,
		PracticeScope: doctor.PracticeScope,
		Status:        doctor.Status,
		LastLoginIp:   doctor.LastLoginIP,
	}

	// 只有当Name不为空时才设置
	if doctor.Name != "" {
		doctorModel.Name = doctor.Name
	}

	// 只有当LicenseNumber不为空时才设置
	if doctor.LicenseNumber != "" {
		doctorModel.LicenseNumber = doctor.LicenseNumber
	}

	// 处理指针类型的ID字段
	if doctor.DepartmentID != nil {
		doctorModel.DepartmentId = uint64(*doctor.DepartmentID)
	}
	if doctor.HospitalID != nil {
		doctorModel.HospitalId = uint64(*doctor.HospitalID)
	}

	// 处理字符串类型的时间字段
	if doctor.BirthDate != "" {
		doctorModel.BirthDate = doctor.BirthDate
	}
	if doctor.LastLoginTime != "" {
		doctorModel.LastLoginTime = doctor.LastLoginTime
	}

	err := d.data.db.WithContext(ctx).Create(doctorModel).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("创建医生记录失败: phone=%s, error=%v", doctor.Phone, err)
		return fmt.Errorf("创建医生记录失败: %w", err)
	}

	// 更新ID到业务实体
	doctor.ID = uint(doctorModel.Id)
	doctor.DoctorCode = doctorModel.DoctorCode
	doctor.CreatedAt = doctorModel.CreatedAt
	doctor.UpdatedAt = doctorModel.UpdatedAt

	d.logger.WithContext(ctx).Infof("医生记录创建成功: id=%d, phone=%s", doctorModel.Id, doctor.Phone)
	return nil
}

// GetDoctorByPhone 根据手机号获取医生信息
func (d *DoctorData) GetDoctorByPhone(ctx context.Context, phone string) (*biz.Doctor, error) {
	var doctorModel model.Doctors

	err := d.data.db.WithContext(ctx).Where("phone = ?", phone).First(&doctorModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("医生不存在")
		}
		d.logger.WithContext(ctx).Errorf("查询医生失败: phone=%s, error=%v", phone, err)
		return nil, fmt.Errorf("查询医生失败: %w", err)
	}

	doctor := d.modelToEntity(&doctorModel)
	return doctor, nil
}

// GetDoctorByID 根据ID获取医生信息
func (d *DoctorData) GetDoctorByID(ctx context.Context, id uint) (*biz.Doctor, error) {
	var doctorModel model.Doctors

	err := d.data.db.WithContext(ctx).Where("id = ?", id).First(&doctorModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("医生不存在")
		}
		d.logger.WithContext(ctx).Errorf("查询医生失败: id=%d, error=%v", id, err)
		return nil, fmt.Errorf("查询医生失败: %w", err)
	}

	doctor := d.modelToEntity(&doctorModel)
	return doctor, nil
}

// UpdateDoctor 更新医生信息
func (d *DoctorData) UpdateDoctor(ctx context.Context, doctor *biz.Doctor) error {
	updateData := map[string]any{
		"phone":          doctor.Phone,
		"password":       doctor.Password,
		"name":           doctor.Name,
		"gender":         doctor.Gender,
		"email":          doctor.Email,
		"avatar":         doctor.Avatar,
		"license_number": doctor.LicenseNumber,
		"title":          doctor.Title,
		"speciality":     doctor.Speciality,
		"practice_scope": doctor.PracticeScope,
		"status":         doctor.Status,
		"last_login_ip":  doctor.LastLoginIP,
		"updated_at":     time.Now(),
	}

	// 处理可选的ID字段
	if doctor.DepartmentID != nil {
		updateData["department_id"] = uint64(*doctor.DepartmentID)
	} else {
		updateData["department_id"] = nil
	}

	if doctor.HospitalID != nil {
		updateData["hospital_id"] = uint64(*doctor.HospitalID)
	} else {
		updateData["hospital_id"] = nil
	}

	// 处理出生日期字段（字符串格式）
	if doctor.BirthDate != "" {
		updateData["birth_date"] = doctor.BirthDate
	} else {
		updateData["birth_date"] = nil
	}

	// 处理最后登录时间（字符串格式）
	if doctor.LastLoginTime != "" {
		updateData["last_login_time"] = doctor.LastLoginTime
	}

	err := d.data.db.WithContext(ctx).Model(&model.Doctors{}).Where("id = ?", doctor.ID).Updates(updateData).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("更新医生信息失败: id=%d, error=%v", doctor.ID, err)
		return fmt.Errorf("更新医生信息失败: %w", err)
	}

	doctor.UpdatedAt = updateData["updated_at"].(time.Time)
	d.logger.WithContext(ctx).Infof("医生信息更新成功: id=%d", doctor.ID)
	return nil
}

// modelToEntity 将数据模型转换为业务实体
func (d *DoctorData) modelToEntity(doctorModel *model.Doctors) *biz.Doctor {
	// 处理指针类型的ID字段
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
		Name:          doctorModel.Name,
		Gender:        doctorModel.Gender,
		BirthDate:     doctorModel.BirthDate, // 直接使用字符串
		Phone:         doctorModel.Phone,
		Password:      doctorModel.Password,
		Email:         doctorModel.Email,
		Avatar:        doctorModel.Avatar,
		LicenseNumber: doctorModel.LicenseNumber,
		DepartmentID:  departmentID,
		HospitalID:    hospitalID,
		Title:         doctorModel.Title,
		Speciality:    doctorModel.Speciality,
		PracticeScope: doctorModel.PracticeScope,
		Status:        doctorModel.Status,
		LastLoginTime: doctorModel.LastLoginTime, // 直接使用字符串
		LastLoginIP:   doctorModel.LastLoginIp,
		CreatedAt:     doctorModel.CreatedAt,
		UpdatedAt:     doctorModel.UpdatedAt,
	}
}

// DeleteDoctorByID 删除医生账号
func (d *DoctorData) DeleteDoctorByID(ctx context.Context, id uint) error {
	err := d.data.db.WithContext(ctx).Delete(&model.Doctors{}, id).Error
	if err != nil {
		d.logger.WithContext(ctx).Errorf("删除医生失败: id=%d, error=%v", id, err)
		return fmt.Errorf("删除医生失败: %w", err)
	}
	d.logger.WithContext(ctx).Infof("医生删除成功: id=%d", id)
	return nil
}

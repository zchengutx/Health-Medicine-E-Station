package model

import (
	"time"

	"gorm.io/gorm"
)

// Doctors 医生表模型
type Doctors struct {
	Id            uint64         `gorm:"primaryKey;autoIncrement;comment:医生ID" json:"id"`
	DoctorCode    string         `gorm:"column:doctor_code;type:varchar(32);uniqueIndex:uk_doctor_code;comment:医生编码" json:"doctor_code"`
	Name          string         `gorm:"column:name;type:varchar(50);comment:医生姓名" json:"name"`
	Gender        string         `gorm:"column:gender;type:varchar(10);not null;default:男;comment:性别：男/女" json:"gender"`
	BirthDate     string         `gorm:"column:birth_date;type:varchar(30);comment:出生日期" json:"birth_date"`
	Phone         string         `gorm:"column:phone;type:char(11);not null;uniqueIndex:uk_phone;comment:手机号码" json:"phone"`
	Password      string         `gorm:"column:password;type:varchar(255);not null;comment:密码" json:"password"`
	Email         string         `gorm:"column:email;type:varchar(100);comment:邮箱地址" json:"email"`
	Avatar        string         `gorm:"column:avatar;type:varchar(255);comment:头像URL" json:"avatar"`
	LicenseNumber string         `gorm:"column:license_number;type:varchar(50);uniqueIndex:uk_license_number;comment:执业医师资格证号" json:"license_number"`
	DepartmentId  uint64         `gorm:"column:department_id;type:bigint unsigned;index:idx_department_id;comment:科室ID" json:"department_id"`
	HospitalId    uint64         `gorm:"column:hospital_id;type:bigint unsigned;index:idx_hospital_id;comment:医院ID" json:"hospital_id"`
	Title         string         `gorm:"column:title;type:varchar(50);comment:职称" json:"title"`
	Speciality    string         `gorm:"column:speciality;type:text;comment:专业特长" json:"speciality"`
	PracticeScope string         `gorm:"column:practice_scope;type:text;comment:执业范围" json:"practice_scope"`
	Status        string         `gorm:"column:status;type:varchar(10);default:启用;index:idx_status;comment:状态：禁用/启用/待审核" json:"status"`
	LastLoginTime string         `gorm:"column:last_login_time;type:varchar(30);comment:最后登录时间" json:"last_login_time"`
	LastLoginIp   string         `gorm:"column:last_login_ip;type:varchar(45);comment:最后登录IP" json:"last_login_ip"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6);index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6);comment:更新时间" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);comment:删除时间" json:"deleted_at"`
}

func (d *Doctors) TableName() string {
	return "doctors"
}

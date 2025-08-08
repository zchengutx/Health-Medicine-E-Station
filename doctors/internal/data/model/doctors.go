package model

import (
	"gorm.io/gorm"
	"time"
)

type Doctors struct {
	Id            uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:医生ID;primaryKey;not null;" json:"id"`                                // 医生ID
	DoctorCode    string         `gorm:"column:doctor_code;type:varchar(32);comment:医生编码;not null;" json:"doctor_code"`                          // 医生编码
	Name          string         `gorm:"column:name;type:varchar(50);comment:医生姓名;default:NULL;" json:"name"`                                       // 医生姓名
	Gender        string         `gorm:"column:gender;type:varchar(10);comment:性别：男/女;not null;default:男;" json:"gender"`                         // 性别：男/女
	BirthDate     time.Time      `gorm:"column:birth_date;type:date;comment:出生日期;default:NULL;" json:"birth_date"`                              // 出生日期
	Phone         string         `gorm:"column:phone;type:char(11);comment:手机号码;not null;" json:"phone"`                                        // 手机号码
	Password      string         `gorm:"column:password;type:varchar(255);comment:密码;not null;" json:"password"`                                     // 密码
	Email         string         `gorm:"column:email;type:varchar(100);comment:邮箱地址;default:NULL;" json:"email"`                                // 邮箱地址
	Avatar        string         `gorm:"column:avatar;type:varchar(255);comment:头像URL;default:NULL;" json:"avatar"`                              // 头像URL
	LicenseNumber string         `gorm:"column:license_number;type:varchar(50);comment:执业医师资格证号;default:NULL;" json:"license_number"`               // 执业医师资格证号
	DepartmentId  uint64         `gorm:"column:department_id;type:bigint UNSIGNED;comment:科室ID;default:NULL;" json:"department_id"`               // 科室ID
	HospitalId    uint64         `gorm:"column:hospital_id;type:bigint UNSIGNED;comment:医院ID;default:NULL;" json:"hospital_id"`                  // 医院ID
	Title         string         `gorm:"column:title;type:varchar(50);comment:职称;default:NULL;" json:"title"`                                   // 职称
	Speciality    string         `gorm:"column:speciality;type:text;comment:专业特长;default:NULL;" json:"speciality"`                              // 专业特长
	PracticeScope string         `gorm:"column:practice_scope;type:text;comment:执业范围;default:NULL;" json:"practice_scope"`                       // 执业范围
	Status        string         `gorm:"column:status;type:varchar(10);comment:状态：禁用/启用/待审核;not null;default:启用;" json:"status"`                // 状态：禁用/启用/待审核
	LastLoginTime time.Time      `gorm:"column:last_login_time;type:timestamp;comment:最后登录时间;default:NULL;" json:"last_login_time"`              // 最后登录时间
	LastLoginIp   string         `gorm:"column:last_login_ip;type:varchar(45);comment:最后登录IP;default:NULL;" json:"last_login_ip"`                // 最后登录IP
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime(6);comment:创建时间;not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"` // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime(6);comment:更新时间;not null;default:CURRENT_TIMESTAMP(6);" json:"updated_at"` // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);comment:删除时间;default:NULL;" json:"deleted_at"`                         // 删除时间
}

func (d *Doctors)TableName() string {
	return "doctors"
}

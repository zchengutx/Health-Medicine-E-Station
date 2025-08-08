package model

import (
	"time"

	"gorm.io/gorm"
)

type Appointments struct {
	Id              uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:预约ID;primaryKey;not null;" json:"id"`                               // 预约ID
	AppointmentNo   string         `gorm:"column:appointment_no;type:varchar(32);comment:预约号码;not null;" json:"appointment_no"`                      // 预约号码
	DoctorId        uint64         `gorm:"column:doctor_id;type:bigint UNSIGNED;comment:医生ID;not null;" json:"doctor_id"`                            // 医生ID
	PatientId       uint64         `gorm:"column:patient_id;type:bigint UNSIGNED;comment:患者ID;not null;" json:"patient_id"`                          // 患者ID
	DepartmentId    uint64         `gorm:"column:department_id;type:bigint UNSIGNED;comment:科室ID;not null;" json:"department_id"`                    // 科室ID
	AppointmentDate time.Time      `gorm:"column:appointment_date;type:date;comment:预约日期;not null;" json:"appointment_date"`                         // 预约日期
	AppointmentTime time.Time      `gorm:"column:appointment_time;type:datetime(3);comment:预约时间;not null;" json:"appointment_time"`                  // 预约时间
	AppointmentType string         `gorm:"column:appointment_type;type:varchar(20);comment:预约类型：专家号/普通号;not null;" json:"appointment_type"`          // 预约类型：专家号/普通号
	Fee             float64        `gorm:"column:fee;type:decimal(10, 2);comment:预约费;not null;default:0.00;" json:"fee"`                             // 预约费
	Status          string         `gorm:"column:status;type:varchar(20);comment:状态：已取消/已预约/已就诊/已过期;not null;default:已预约;" json:"status"`            // 状态：已取消/已预约/已就诊/已过期
	Source          string         `gorm:"column:source;type:varchar(20);comment:预约来源：线上/线下;default:线上;" json:"source"`                              // 预约来源：线上/线下
	Notes           string         `gorm:"column:notes;type:varchar(500);comment:备注;default:NULL;" json:"notes"`                                     // 备注
	CreatedAt       time.Time      `gorm:"column:created_at;type:datetime(6);comment:创建时间;not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"` // 创建时间
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:datetime(6);comment:更新时间;not null;default:CURRENT_TIMESTAMP(6);" json:"updated_at"` // 更新时间
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);comment:删除时间;default:NULL;" json:"deleted_at"`                          // 删除时间
}

func (a *Appointments) TableName() string {
	return "appointments"
}

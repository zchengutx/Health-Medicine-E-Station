package model

import (
	"time"

	"gorm.io/gorm"
)

// Consultations 问诊表 - 基于appointments表扩展
type Consultations struct {
	Id               uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:问诊ID;primaryKey;not null;" json:"id"`                          // 问诊ID
	ConsultationCode string         `gorm:"column:consultation_code;type:varchar(32);comment:问诊编码;not null;" json:"consultation_code"`           // 问诊编码
	PatientId        uint64         `gorm:"column:patient_id;type:bigint UNSIGNED;comment:患者ID;not null;" json:"patient_id"`                     // 患者ID
	DoctorId         uint64         `gorm:"column:doctor_id;type:bigint UNSIGNED;comment:医生ID;not null;" json:"doctor_id"`                       // 医生ID
	Type             string         `gorm:"column:type;type:varchar(20);comment:问诊类型;not null;" json:"type"`                                     // 问诊类型
	Status           string         `gorm:"column:status;type:varchar(20);comment:状态;not null;default:待开始;" json:"status"`                       // 状态
	ChiefComplaint   string         `gorm:"column:chief_complaint;type:text;comment:主诉;default:NULL;" json:"chief_complaint"`                    // 主诉
	PresentIllness   string         `gorm:"column:present_illness;type:text;comment:现病史;default:NULL;" json:"present_illness"`                   // 现病史
	Symptoms         string         `gorm:"column:symptoms;type:text;comment:症状;default:NULL;" json:"symptoms"`                                  // 症状
	Diagnosis        string         `gorm:"column:diagnosis;type:text;comment:诊断;default:NULL;" json:"diagnosis"`                                // 诊断
	Treatment        string         `gorm:"column:treatment;type:text;comment:治疗方案;default:NULL;" json:"treatment"`                              // 治疗方案
	Prescription     string         `gorm:"column:prescription;type:text;comment:处方;default:NULL;" json:"prescription"`                          // 处方
	Advice           string         `gorm:"column:advice;type:text;comment:医嘱;default:NULL;" json:"advice"`                                      // 医嘱
	StartTime        time.Time      `gorm:"column:start_time;type:timestamp;comment:开始时间;default:NULL;" json:"start_time"`                       // 开始时间
	EndTime          time.Time      `gorm:"column:end_time;type:timestamp;comment:结束时间;default:NULL;" json:"end_time"`                           // 结束时间
	Duration         int            `gorm:"column:duration;type:int;comment:持续时间(分钟);default:0;" json:"duration"`                                // 持续时间(分钟)
	Fee              float64        `gorm:"column:fee;type:decimal(10,2);comment:费用;default:0.00;" json:"fee"`                                   // 费用
	PaymentStatus    string         `gorm:"column:payment_status;type:varchar(20);comment:支付状态;default:未支付;" json:"payment_status"`              // 支付状态
	Rating           int            `gorm:"column:rating;type:int;comment:评分;default:0;" json:"rating"`                                          // 评分
	Feedback         string         `gorm:"column:feedback;type:text;comment:反馈;default:NULL;" json:"feedback"`                                  // 反馈
	Remarks          string         `gorm:"column:remarks;type:varchar(500);comment:备注;default:NULL;" json:"remarks"`                            // 备注
	CreatedAt        time.Time      `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;comment:删除时间;default:NULL;" json:"deleted_at"`                       // 删除时间
}

func (c *Consultations) TableName() string {
	return "consultations"
}

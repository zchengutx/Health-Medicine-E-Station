package model

import (
	"time"

	"gorm.io/gorm"
)

type FollowUpPlans struct {
	Id               uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:随访计划ID;primaryKey;not null;" json:"id"`                             // 随访计划ID
	PlanNo           string         `gorm:"column:plan_no;type:varchar(32);comment:随访计划编号;not null;" json:"plan_no"`                                  // 随访计划编号
	DoctorId         uint64         `gorm:"column:doctor_id;type:bigint UNSIGNED;comment:医生ID;not null;" json:"doctor_id"`                            // 医生ID
	PatientId        uint64         `gorm:"column:patient_id;type:bigint UNSIGNED;comment:患者ID;not null;" json:"patient_id"`                          // 患者ID
	MedicalRecordId  uint64         `gorm:"column:medical_record_id;type:bigint UNSIGNED;comment:病历ID;default:NULL;" json:"medical_record_id"`        // 病历ID
	PlanName         string         `gorm:"column:plan_name;type:varchar(100);comment:随访计划名称;not null;" json:"plan_name"`                             // 随访计划名称
	FollowUpType     string         `gorm:"column:follow_up_type;type:varchar(20);comment:随访类型：电话/复诊/在线;not null;" json:"follow_up_type"`             // 随访类型：电话/复诊/在线
	StartDate        time.Time      `gorm:"column:start_date;type:date;comment:开始日期;not null;" json:"start_date"`                                     // 开始日期
	EndDate          time.Time      `gorm:"column:end_date;type:date;comment:结束日期;default:NULL;" json:"end_date"`                                     // 结束日期
	Frequency        string         `gorm:"column:frequency;type:varchar(50);comment:随访频率;default:NULL;" json:"frequency"`                            // 随访频率
	NextFollowUpDate time.Time      `gorm:"column:next_follow_up_date;type:date;comment:下次随访日期;default:NULL;" json:"next_follow_up_date"`             // 下次随访日期
	FollowUpContent  string         `gorm:"column:follow_up_content;type:text;comment:随访内容;default:NULL;" json:"follow_up_content"`                   // 随访内容
	Notes            string         `gorm:"column:notes;type:text;comment:备注;default:NULL;" json:"notes"`                                             // 备注
	Status           string         `gorm:"column:status;type:varchar(10);comment:状态：已取消/进行中/已完成/已暂停;not null;" json:"status"`                        // 状态：已取消/进行中/已完成/已暂停
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime(6);comment:创建时间;not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"` // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime(6);comment:更新时间;not null;default:CURRENT_TIMESTAMP(6);" json:"updated_at"` // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);comment:删除时间;default:NULL;" json:"deleted_at"`                          // 删除时间
}

func (f *FollowUpPlans) TableName() string {
	return "follow_up_plans"
}

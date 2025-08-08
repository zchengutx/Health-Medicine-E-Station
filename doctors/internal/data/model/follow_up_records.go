package model

import (
	"time"

	"gorm.io/gorm"
)

type FollowUpRecords struct {
	Id                  uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:随访记录ID;primaryKey;not null;" json:"id"`                             // 随访记录ID
	RecordNo            string         `gorm:"column:record_no;type:varchar(32);comment:随访记录编号;not null;" json:"record_no"`                              // 随访记录编号
	FollowUpPlanId      uint64         `gorm:"column:follow_up_plan_id;type:bigint UNSIGNED;comment:随访计划ID;not null;" json:"follow_up_plan_id"`          // 随访计划ID
	DoctorId            uint64         `gorm:"column:doctor_id;type:bigint UNSIGNED;comment:医生ID;not null;" json:"doctor_id"`                            // 医生ID
	PatientId           uint64         `gorm:"column:patient_id;type:bigint UNSIGNED;comment:患者ID;not null;" json:"patient_id"`                          // 患者ID
	FollowUpDate        time.Time      `gorm:"column:follow_up_date;type:date;comment:随访日期;not null;" json:"follow_up_date"`                             // 随访日期
	FollowUpTime        time.Time      `gorm:"column:follow_up_time;type:time;comment:随访时间;default:NULL;" json:"follow_up_time"`                         // 随访时间
	FollowUpType        string         `gorm:"column:follow_up_type;type:varchar(20);comment:随访类型：电话/复诊/在线;not null;" json:"follow_up_type"`             // 随访类型：电话/复诊/在线
	FollowUpContent     string         `gorm:"column:follow_up_content;type:text;comment:随访内容;default:NULL;" json:"follow_up_content"`                   // 随访内容
	PatientCondition    string         `gorm:"column:patient_condition;type:text;comment:患者情况;default:NULL;" json:"patient_condition"`                   // 患者情况
	TreatmentCompliance string         `gorm:"column:treatment_compliance;type:varchar(50);comment:治疗依从性;default:NULL;" json:"treatment_compliance"`     // 治疗依从性
	SideEffects         string         `gorm:"column:side_effects;type:text;comment:不良反应;default:NULL;" json:"side_effects"`                             // 不良反应
	NextPlan            string         `gorm:"column:next_plan;type:text;comment:下一步计划;default:NULL;" json:"next_plan"`                                  // 下一步计划
	FollowUpResult      string         `gorm:"column:follow_up_result;type:varchar(50);comment:随访结果：好转/稳定/恶化;default:NULL;" json:"follow_up_result"`     // 随访结果：好转/稳定/恶化
	DurationMinutes     int32          `gorm:"column:duration_minutes;type:int;comment:随访时长（分钟）;default:NULL;" json:"duration_minutes"`                  // 随访时长（分钟）
	Notes               string         `gorm:"column:notes;type:text;comment:备注;default:NULL;" json:"notes"`                                             // 备注
	Status              string         `gorm:"column:status;type:varchar(10);comment:状态：已取消/已完成;not null;" json:"status"`                                // 状态：已取消/已完成
	CreatedAt           time.Time      `gorm:"column:created_at;type:datetime(6);comment:创建时间;not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"` // 创建时间
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:datetime(6);comment:更新时间;not null;default:CURRENT_TIMESTAMP(6);" json:"updated_at"` // 更新时间
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);comment:删除时间;default:NULL;" json:"deleted_at"`                          // 删除时间
}

func (f *FollowUpRecords) TableName() string {
	return "follow_up_records"
}

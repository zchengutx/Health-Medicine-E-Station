package model

import (
	"time"
)

type MedicalRecords struct {
	Id                   uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:病历ID;primaryKey;not null;" json:"id"`                                         // 病历ID
	RecordNo             string    `gorm:"column:record_no;type:varchar(32);comment:病历号;not null;" json:"record_no"`                                           // 病历号
	DoctorId             uint64    `gorm:"column:doctor_id;type:bigint UNSIGNED;comment:医生ID;not null;" json:"doctor_id"`                                      // 医生ID
	PatientId            uint64    `gorm:"column:patient_id;type:bigint UNSIGNED;comment:患者ID;not null;" json:"patient_id"`                                    // 患者ID
	AppointmentId        uint64    `gorm:"column:appointment_id;type:bigint UNSIGNED;comment:挂号ID;default:NULL;" json:"appointment_id"`                        // 挂号ID
	VisitDate            time.Time `gorm:"column:visit_date;type:date;comment:就诊日期;not null;" json:"visit_date"`                                               // 就诊日期
	ChiefComplaint       string    `gorm:"column:chief_complaint;type:text;comment:主诉;default:NULL;" json:"chief_complaint"`                                   // 主诉
	PresentIllness       string    `gorm:"column:present_illness;type:text;comment:现病史;default:NULL;" json:"present_illness"`                                  // 现病史
	PastHistory          string    `gorm:"column:past_history;type:text;comment:既往史;default:NULL;" json:"past_history"`                                        // 既往史
	PhysicalExamination  string    `gorm:"column:physical_examination;type:text;comment:体格检查;default:NULL;" json:"physical_examination"`                       // 体格检查
	AuxiliaryExamination string    `gorm:"column:auxiliary_examination;type:text;comment:辅助检查;default:NULL;" json:"auxiliary_examination"`                     // 辅助检查
	Diagnosis            string    `gorm:"column:diagnosis;type:text;comment:诊断;default:NULL;" json:"diagnosis"`                                               // 诊断
	TreatmentPlan        string    `gorm:"column:treatment_plan;type:text;comment:治疗方案;default:NULL;" json:"treatment_plan"`                                   // 治疗方案
	DoctorAdvice         string    `gorm:"column:doctor_advice;type:text;comment:医嘱;default:NULL;" json:"doctor_advice"`                                       // 医嘱
	FollowUpPlan         string    `gorm:"column:follow_up_plan;type:text;comment:随访计划;default:NULL;" json:"follow_up_plan"`                                   // 随访计划
	RecordType           string    `gorm:"column:record_type;type:varchar(20);comment:病历类型：outpatient-门诊，inpatient-住院;default:outpatient;" json:"record_type"` // 病历类型：outpatient-门诊，inpatient-住院
	Status               string    `gorm:"column:status;type:varchar(20);comment:状态：0-草稿，1-已完成，2-已归档;not null;default:1;" json:"status"`                       // 状态：0-草稿，1-已完成，2-已归档
	CreatedAt            time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`                // 创建时间
	UpdatedAt            time.Time `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"`                // 更新时间
}

func (m *MedicalRecords) TableName() string {
	return "medical_records"
}

package model

import (
	"time"
)

type Prescriptions struct {
	Id               uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:处方ID;primaryKey;not null;" json:"id"`                            // 处方ID
	PrescriptionNo   string    `gorm:"column:prescription_no;type:varchar(32);comment:处方号;not null;" json:"prescription_no"`                  // 处方号
	DoctorId         uint64    `gorm:"column:doctor_id;type:bigint UNSIGNED;comment:医生ID;not null;" json:"doctor_id"`                         // 医生ID
	PatientId        uint64    `gorm:"column:patient_id;type:bigint UNSIGNED;comment:患者ID;not null;" json:"patient_id"`                       // 患者ID
	MedicalRecordId  uint64    `gorm:"column:medical_record_id;type:bigint UNSIGNED;comment:病历ID;default:NULL;" json:"medical_record_id"`     // 病历ID
	PrescriptionDate time.Time `gorm:"column:prescription_date;type:date;comment:处方日期;not null;" json:"prescription_date"`                    // 处方日期
	TotalAmount      float64   `gorm:"column:total_amount;type:decimal(10,2);comment:处方总金额;not null;default:0.00;" json:"total_amount"`       // 处方总金额
	PrescriptionType string    `gorm:"column:prescription_type;type:varchar(20);comment:处方类型：西药/中药/中西药;default:西药;" json:"prescription_type"` // 处方类型：西药/中药/中西药
	UsageInstruction string    `gorm:"column:usage_instruction;type:text;comment:用药说明;default:NULL;" json:"usage_instruction"`                // 用药说明
	Status           string    `gorm:"column:status;type:varchar(20);comment:状态：已取消/已开具/已审核/已发药;not null;default:已开具;" json:"status"`         // 状态：已取消/已开具/已审核/已发药
	AuditorId        uint64    `gorm:"column:auditor_id;type:bigint UNSIGNED;comment:审核医生ID;default:NULL;" json:"auditor_id"`                 // 审核医生ID
	AuditTime        time.Time `gorm:"column:audit_time;type:timestamp;comment:审核时间;default:NULL;" json:"audit_time"`                         // 审核时间
	AuditNotes       string    `gorm:"column:audit_notes;type:varchar(500);comment:审核备注;default:NULL;" json:"audit_notes"`                    // 审核备注
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`   // 创建时间
	UpdatedAt        time.Time `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"`   // 更新时间
}

func (p *Prescriptions) TableName() string {
	return "prescriptions"
}

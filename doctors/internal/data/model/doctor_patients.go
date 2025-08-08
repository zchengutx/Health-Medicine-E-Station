package model

import (
	"time"

	"gorm.io/gorm"
)

type DoctorPatients struct {
	Id               uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:ID;primaryKey;not null;" json:"id"`                                        // ID
	DoctorId         uint64         `gorm:"column:doctor_id;type:bigint UNSIGNED;comment:医生ID;not null;" json:"doctor_id"`                                   // 医生ID
	PatientId        uint64         `gorm:"column:patient_id;type:bigint UNSIGNED;comment:患者ID;not null;" json:"patient_id"`                                 // 患者ID
	RelationshipType string         `gorm:"column:relationship_type;type:varchar(20);comment:关系类型: 普通/关注/VIP;not null;default:普通;" json:"relationship_type"` // 关系类型: 普通/关注/VIP
	Tags             string         `gorm:"column:tags;type:varchar(200);comment:患者标签，逗号分隔;default:NULL;" json:"tags"`                                       // 患者标签，逗号分隔
	Notes            string         `gorm:"column:notes;type:text;comment:备注;default:NULL;" json:"notes"`                                                    // 备注
	FirstVisitTime   time.Time      `gorm:"column:first_visit_time;type:timestamp;comment:首次就诊时间;default:NULL;" json:"first_visit_time"`                     // 首次就诊时间
	LastVisitTime    time.Time      `gorm:"column:last_visit_time;type:timestamp;comment:最后就诊时间;default:NULL;" json:"last_visit_time"`                       // 最后就诊时间
	VisitCount       int32          `gorm:"column:visit_count;type:int;comment:就诊次数;default:0;" json:"visit_count"`                                          // 就诊次数
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime(6);comment:创建时间;not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"`        // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime(6);comment:更新时间;not null;default:CURRENT_TIMESTAMP(6);" json:"updated_at"`        // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);comment:删除时间;default:NULL;" json:"deleted_at"`                                 // 删除时间
}

func (d *DoctorPatients) TableName() string {
	return "doctor_patients"
}

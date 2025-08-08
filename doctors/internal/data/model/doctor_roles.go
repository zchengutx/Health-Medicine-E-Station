package model

import "time"

type DoctorRoles struct {
	Id        uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:ID;primaryKey;not null;" json:"id"`                                 // ID
	DoctorId  uint64    `gorm:"column:doctor_id;type:bigint UNSIGNED;comment:医生ID;not null;" json:"doctor_id"`                            // 医生ID
	RoleId    uint64    `gorm:"column:role_id;type:bigint UNSIGNED;comment:角色ID;not null;" json:"role_id"`                                // 角色ID
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(6);comment:创建时间;not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"` // 创建时间
}

func (d *DoctorRoles) TableName() string {
	return "doctor_roles"
}

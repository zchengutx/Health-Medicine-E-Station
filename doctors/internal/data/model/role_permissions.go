package model

import "time"

type RolePermissions struct {
	Id           uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:ID;primaryKey;not null;" json:"id"`                                 // ID
	RoleId       uint64    `gorm:"column:role_id;type:bigint UNSIGNED;comment:角色ID;not null;" json:"role_id"`                                // 角色ID
	PermissionId uint64    `gorm:"column:permission_id;type:bigint UNSIGNED;comment:权限ID;not null;" json:"permission_id"`                    // 权限ID
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(6);comment:创建时间;not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"` // 创建时间
}

func (r *RolePermissions) TableName() string {
	return "role_permissions"
}

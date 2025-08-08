package model

import "time"

type Roles struct {
	Id          uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:角色ID;primaryKey;not null;" json:"id"`                          // 角色ID
	RoleCode    string    `gorm:"column:role_code;type:varchar(32);comment:角色编码;not null;" json:"role_code"`                           // 角色编码
	RoleName    string    `gorm:"column:role_name;type:varchar(50);comment:角色名称;not null;" json:"role_name"`                           // 角色名称
	Description string    `gorm:"column:description;type:varchar(200);comment:角色描述;default:NULL;" json:"description"`                  // 角色描述
	Status      string    `gorm:"column:status;type:varchar(10);comment:状态：禁用/启用;not null;default:启用;" json:"status"`                  // 状态：禁用/启用
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (r *Roles) TableName() string {
	return "roles"
}

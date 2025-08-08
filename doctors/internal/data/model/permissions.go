package model

import "time"

type Permissions struct {
	Id             uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:权限ID;primaryKey;not null;" json:"id"`                          // 权限ID
	PermissionCode string    `gorm:"column:permission_code;type:varchar(64);comment:权限编码;not null;" json:"permission_code"`               // 权限编码
	PermissionName string    `gorm:"column:permission_name;type:varchar(100);comment:权限名称;not null;" json:"permission_name"`              // 权限名称
	ResourceType   string    `gorm:"column:resource_type;type:varchar(32);comment:资源类型：菜单/按钮/接口;not null;" json:"resource_type"`          // 资源类型：菜单/按钮/接口
	ResourcePath   string    `gorm:"column:resource_path;type:varchar(200);comment:资源路径;default:NULL;" json:"resource_path"`              // 资源路径
	ParentId       uint64    `gorm:"column:parent_id;type:bigint UNSIGNED;comment:父权限ID;default:0;" json:"parent_id"`                     // 父权限ID
	SortOrder      int32     `gorm:"column:sort_order;type:int;comment:排序;default:0;" json:"sort_order"`                                  // 排序
	Status         string    `gorm:"column:status;type:varchar(10);comment:状态：禁用/启用;not null;default:启用;" json:"status"`                  // 状态：禁用/启用
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (p *Permissions) TableName() string {
	return "permissions"
}

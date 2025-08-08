package model

import "time"

type SystemConfigs struct {
	Id          uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:配置ID;primaryKey;not null;" json:"id"`                                                      // 配置ID
	ConfigKey   string    `gorm:"column:config_key;type:varchar(100);comment:配置键;not null;uniqueIndex:uk_config_key;" json:"config_key"`                           // 配置键
	ConfigValue string    `gorm:"column:config_value;type:text;comment:配置值;default:NULL;" json:"config_value"`                                                     // 配置值
	ConfigType  string    `gorm:"column:config_type;type:varchar(20);comment:配置类型：string-字符串，number-数字，boolean-布尔，json-JSON;default:'';" json:"config_type"`       // 配置类型：string-字符串，number-数字，boolean-布尔，json-JSON
	Description string    `gorm:"column:description;type:varchar(200);comment:配置描述;default:NULL;" json:"description"`                                              // 配置描述
	GroupName   string    `gorm:"column:group_name;type:varchar(50);comment:配置分组;default:NULL;index:idx_group_name;" json:"group_name"`                            // 配置分组
	IsSystem    string    `gorm:"column:is_system;type:varchar(10);comment:是否系统配置：否/是;default:否;" json:"is_system"`                                                // 是否系统配置：否/是
	Status      string    `gorm:"column:status;type:varchar(10);comment:状态: 禁用/启用;not null;default:启用;index:idx_status;" json:"status"`                            // 状态: 禁用/启用
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`                             // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (s *SystemConfigs) TableName() string {
	return "system_configs"
}

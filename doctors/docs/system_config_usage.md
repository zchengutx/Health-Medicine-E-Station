# 系统配置使用说明

## 概述

`SystemConfigModel` 是系统配置表的数据模型，用于存储系统的各种配置信息。

## 数据库表结构

```sql
CREATE TABLE `system_configs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '配置ID',
  `config_key` varchar(100) NOT NULL COMMENT '配置键',
  `config_value` text COMMENT '配置值',
  `config_type` varchar(20) DEFAULT '' COMMENT '配置类型：string-字符串，number-数字，boolean-布尔，json-JSON',
  `description` varchar(200) DEFAULT NULL COMMENT '配置描述',
  `group_name` varchar(50) DEFAULT NULL COMMENT '配置分组',
  `is_system` varchar(10) DEFAULT '否' COMMENT '是否系统配置：否/是',
  `status` varchar(10) NOT NULL DEFAULT '启用' COMMENT '状态：禁用/启用',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_config_key` (`config_key`),
  KEY `idx_group_name` (`group_name`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB COMMENT='系统配置表';
```

## Go 结构体

```go
type SystemConfigModel struct {
    ID          uint      `gorm:"primaryKey;autoIncrement;comment:配置ID" json:"id"`
    ConfigKey   string    `gorm:"uniqueIndex:uk_config_key;size:100;not null;comment:配置键" json:"config_key"`
    ConfigValue string    `gorm:"type:text;comment:配置值" json:"config_value"`
    ConfigType  string    `gorm:"size:20;default:'';comment:配置类型" json:"config_type"`
    Description string    `gorm:"size:200;comment:配置描述" json:"description"`
    GroupName   string    `gorm:"size:50;index:idx_group_name;comment:配置分组" json:"group_name"`
    IsSystem    string    `gorm:"size:10;default:否;comment:是否系统配置" json:"is_system"`
    Status      string    `gorm:"size:10;not null;default:启用;index:idx_status;comment:状态" json:"status"`
    CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
    UpdatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}
```

## 使用方法

### 1. 创建配置

```go
config := &SystemConfigModel{
    ConfigKey:   "sms.provider",
    ConfigValue: "aliyun",
    ConfigType:  "string",
    Description: "短信服务提供商",
    GroupName:   "短信配置",
    IsSystem:    "否",
    Status:      "启用",
}

err := systemConfigData.CreateConfig(ctx, config)
```

### 2. 获取配置

```go
// 根据配置键获取
config, err := systemConfigData.GetConfigByKey(ctx, "sms.provider")

// 获取配置值（带默认值）
provider := systemConfigData.GetConfigValue(ctx, "sms.provider", "aliyun")

// 获取布尔配置
enabled := systemConfigData.GetConfigBool(ctx, "sms.enabled", true)
```

### 3. 更新配置

```go
// 更新整个配置
config.ConfigValue = "tencent"
err := systemConfigData.UpdateConfig(ctx, config)

// 只更新配置值
err := systemConfigData.UpdateConfigValue(ctx, "sms.provider", "tencent")
```

### 4. 获取分组配置

```go
configs, err := systemConfigData.GetConfigsByGroup(ctx, "短信配置")
```

### 5. 删除配置

```go
err := systemConfigData.DeleteConfig(ctx, "sms.provider")
// 注意：系统配置不能删除
```

## 配置类型

- `string`: 字符串类型
- `number`: 数字类型
- `boolean`: 布尔类型（true/false, 1/0, 是/否）
- `json`: JSON 格式

## 常用配置示例

### 短信配置
```go
{
    ConfigKey: "sms.provider",
    ConfigValue: "aliyun",
    ConfigType: "string",
    GroupName: "短信配置",
    Description: "短信服务提供商"
}

{
    ConfigKey: "sms.enabled",
    ConfigValue: "true",
    ConfigType: "boolean",
    GroupName: "短信配置",
    Description: "是否启用短信服务"
}
```

### 系统配置
```go
{
    ConfigKey: "system.maintenance",
    ConfigValue: "false",
    ConfigType: "boolean",
    GroupName: "系统配置",
    Description: "系统维护模式",
    IsSystem: "是"
}
```

## 注意事项

1. **配置键唯一性**: `config_key` 必须唯一
2. **系统配置保护**: `is_system = "是"` 的配置不能删除
3. **状态控制**: 只有 `status = "启用"` 的配置才会被查询到
4. **分组管理**: 使用 `group_name` 对配置进行分组管理
5. **类型安全**: 使用对应的方法获取不同类型的配置值

## 最佳实践

1. **命名规范**: 使用点号分隔的命名方式，如 `sms.provider`、`system.maintenance`
2. **分组管理**: 相关配置使用相同的分组名
3. **描述完整**: 为每个配置添加清晰的描述
4. **默认值**: 获取配置时总是提供合理的默认值
5. **类型标记**: 正确设置 `config_type` 以便于管理和验证
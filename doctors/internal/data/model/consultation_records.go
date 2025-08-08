package model

import (
	"time"
)

// ConsultationRecords 问诊记录表
type ConsultationRecords struct {
	Id             uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:记录ID;primaryKey;not null;" json:"id"`                          // 记录ID
	ConsultationId uint64    `gorm:"column:consultation_id;type:bigint UNSIGNED;comment:问诊ID;not null;" json:"consultation_id"`           // 问诊ID
	RecordType     string    `gorm:"column:record_type;type:varchar(50);comment:记录类型;not null;" json:"record_type"`                       // 记录类型
	Title          string    `gorm:"column:title;type:varchar(200);comment:标题;default:NULL;" json:"title"`                                // 标题
	Content        string    `gorm:"column:content;type:text;comment:内容;default:NULL;" json:"content"`                                    // 内容
	DataValue      string    `gorm:"column:data_value;type:text;comment:数据值;default:NULL;" json:"data_value"`                             // 数据值
	FileUrl        string    `gorm:"column:file_url;type:varchar(255);comment:文件URL;default:NULL;" json:"file_url"`                       // 文件URL
	CreatedBy      uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者ID;not null;" json:"created_by"`                    // 创建者ID
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (c *ConsultationRecords) TableName() string {
	return "consultation_records"
}

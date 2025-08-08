package model

import (
	"time"
)

// ConsultationMessages 问诊消息表
type ConsultationMessages struct {
	Id             uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:消息ID;primaryKey;not null;" json:"id"`                          // 消息ID
	ConsultationId uint64    `gorm:"column:consultation_id;type:bigint UNSIGNED;comment:问诊ID;not null;" json:"consultation_id"`           // 问诊ID
	SenderId       uint64    `gorm:"column:sender_id;type:bigint UNSIGNED;comment:发送者ID;not null;" json:"sender_id"`                      // 发送者ID
	SenderType     string    `gorm:"column:sender_type;type:varchar(20);comment:发送者类型;not null;" json:"sender_type"`                      // 发送者类型
	MessageType    string    `gorm:"column:message_type;type:varchar(20);comment:消息类型;not null;" json:"message_type"`                     // 消息类型
	Content        string    `gorm:"column:content;type:text;comment:消息内容;default:NULL;" json:"content"`                                  // 消息内容
	MediaUrl       string    `gorm:"column:media_url;type:varchar(255);comment:媒体URL;default:NULL;" json:"media_url"`                     // 媒体URL
	IsRead         bool      `gorm:"column:is_read;type:boolean;comment:是否已读;default:false;" json:"is_read"`                              // 是否已读
	ReadTime       time.Time `gorm:"column:read_time;type:timestamp;comment:阅读时间;default:NULL;" json:"read_time"`                         // 阅读时间
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (c *ConsultationMessages) TableName() string {
	return "consultation_messages"
}

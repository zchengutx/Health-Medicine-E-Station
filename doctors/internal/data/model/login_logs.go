package model

import "time"

type LoginLogs struct {
	Id            uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:日志ID;primaryKey;not null;" json:"id"`                               // 日志ID
	DoctorId      uint64    `gorm:"column:doctor_id;type:bigint UNSIGNED;comment:医生ID;not null;" json:"doctor_id"`                            // 医生ID
	LoginType     string    `gorm:"column:login_type;type:varchar(20);comment:登录类型：password-密码，sms-短信;not null;" json:"login_type"`           // 登录类型：password-密码，sms-短信
	LoginIp       string    `gorm:"column:login_ip;type:varchar(45);comment:登录IP;not null;" json:"login_ip"`                                  // 登录IP
	UserAgent     string    `gorm:"column:user_agent;type:varchar(500);comment:用户代理;default:NULL;" json:"user_agent"`                         // 用户代理
	LoginStatus   string    `gorm:"column:login_status;type:varchar(6);comment:登录状态：失败/成功;not null;" json:"login_status"`                     // 登录状态：失败/成功
	FailureReason string    `gorm:"column:failure_reason;type:varchar(200);comment:失败原因;default:NULL;" json:"failure_reason"`                 // 失败原因
	LoginTime     time.Time `gorm:"column:login_time;type:datetime(6);comment:登录时间;not null;default:CURRENT_TIMESTAMP(6);" json:"login_time"` // 登录时间
}

func (l *LoginLogs) TableName() string {
	return "login_logs"
}

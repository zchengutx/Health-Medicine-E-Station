package model

import "time"

type OperationLogs struct {
	Id            uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:日志ID;primaryKey;not null;" json:"id"`                                                // 日志ID
	UserId        uint64    `gorm:"column:user_id;type:bigint UNSIGNED;comment:用户ID;not null;" json:"user_id"`                                                 // 用户ID
	UserType      string    `gorm:"column:user_type;type:varchar(20);comment:用户类型：医生/管理员;not null;default:医生;" json:"user_type"`                               // 用户类型：医生/管理员
	OperationType string    `gorm:"column:operation_type;type:varchar(50);comment:操作类型：create-创建，update-更新，delete-删除，view-查看;not null;" json:"operation_type"` // 操作类型：create-创建，update-更新，delete-删除，view-查看
	Module        string    `gorm:"column:module;type:varchar(50);comment:操作模块：patient-患者，medical_record-病历等;not null;" json:"module"`                         // 操作模块：patient-患者，medical_record-病历等
	OperationDesc string    `gorm:"column:operation_desc;type:varchar(200);comment:操作描述;not null;" json:"operation_desc"`                                      // 操作描述
	RequestUrl    string    `gorm:"column:request_url;type:varchar(200);comment:请求URL;default:NULL;" json:"request_url"`                                       // 请求URL
	RequestMethod string    `gorm:"column:request_method;type:varchar(10);comment:请求方法：GET，POST，PUT，DELETE;default:NULL;" json:"request_method"`               // 请求方法：GET，POST，PUT，DELETE
	RequestParams string    `gorm:"column:request_params;type:text;comment:请求参数;default:NULL;" json:"request_params"`                                          // 请求参数
	ResponseData  string    `gorm:"column:response_data;type:text;comment:响应数据;default:NULL;" json:"response_data"`                                            // 响应数据
	IpAddress     string    `gorm:"column:ip_address;type:varchar(45);comment:IP地址;default:NULL;" json:"ip_address"`                                           // IP地址
	UserAgent     string    `gorm:"column:user_agent;type:varchar(500);comment:用户代理;default:NULL;" json:"user_agent"`                                          // 用户代理
	ExecutionTime int32     `gorm:"column:execution_time;type:int;comment:执行时间（毫秒）;default:NULL;" json:"execution_time"`                                       // 执行时间（毫秒）
	Status        string    `gorm:"column:status;type:varchar(20);comment:状态：0-失败，1-成功;not null;default:1;" json:"status"`                                     // 状态：0-失败，1-成功
	ErrorMessage  string    `gorm:"column:error_message;type:varchar(500);comment:错误信息;default:NULL;" json:"error_message"`                                    // 错误信息
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`                       // 创建时间
}

func (o *OperationLogs) TableName() string {
	return "operation_logs"
}

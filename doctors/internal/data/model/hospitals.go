package model

import (
	"gorm.io/gorm"
	"time"
)

type Hospitals struct {
	Id            uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:医院ID;primaryKey;not null;" json:"id"`                               // 医院ID
	HospitalCode  string         `gorm:"column:hospital_code;type:varchar(32);comment:医院编码;not null;" json:"hospital_code"`                         // 医院编码
	Name          string         `gorm:"column:name;type:varchar(100);comment:医院名称;not null;" json:"name"`                                        // 医院名称
	ShortName     string         `gorm:"column:short_name;type:varchar(50);comment:医院简称;default:NULL;" json:"short_name"`                          // 医院简称
	Level         string         `gorm:"column:level;type:varchar(20);comment:医院等级：三甲，三乙，二甲等;default:NULL;" json:"level"`                       // 医院等级：三甲，三乙，二甲等
	Type          string         `gorm:"column:type;type:varchar(20);comment:医院类型：综合，专科等;default:NULL;" json:"type"`                             // 医院类型：综合，专科等
	Address       string         `gorm:"column:address;type:varchar(200);comment:医院地址;default:NULL;" json:"address"`                               // 医院地址
	Phone         string         `gorm:"column:phone;type:varchar(20);comment:联系电话;default:NULL;" json:"phone"`                                    // 联系电话
	Website       string         `gorm:"column:website;type:varchar(100);comment:官方网站;default:NULL;" json:"website"`                               // 官方网站
	Introduction  string         `gorm:"column:introduction;type:text;comment:医院介绍;default:NULL;" json:"introduction"`                             // 医院介绍
	LicenseNumber string         `gorm:"column:license_number;type:varchar(50);comment:医疗机构执业许可证号;default:NULL;" json:"license_number"`            // 医疗机构执业许可证号
	Status        string         `gorm:"column:status;type:varchar(10);comment:状态：禁用/启用;not null;" json:"status"`                                 // 状态：禁用/启用
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime(6);comment:创建时间;not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"`  // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime(6);comment:更新时间;not null;default:CURRENT_TIMESTAMP(6);" json:"updated_at"`  // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);comment:删除时间;default:NULL;" json:"deleted_at"`                          // 删除时间
}

func (h *Hospitals)TableName() string {
	return "hospitals"
}

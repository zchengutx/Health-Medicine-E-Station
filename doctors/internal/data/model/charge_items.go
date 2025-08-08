package model

import (
	"time"

	"gorm.io/gorm"
)

type ChargeItems struct {
	Id                    uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:收费项目ID;primaryKey;not null;" json:"id"`                                 // 收费项目ID
	ItemCode              string         `gorm:"column:item_code;type:varchar(32);comment:项目编码;not null;" json:"item_code"`                                    // 项目编码
	ItemName              string         `gorm:"column:item_name;type:varchar(100);comment:项目名称;not null;" json:"item_name"`                                   // 项目名称
	Category              string         `gorm:"column:category;type:varchar(50);comment:项目分类：挂号费/诊疗费/检查费等;not null;" json:"category"`                         // 项目分类：挂号费/诊疗费/检查费等
	Unit                  string         `gorm:"column:unit;type:varchar(20);comment:单位;default:NULL;" json:"unit"`                                            // 单位
	UnitPrice             float64        `gorm:"column:unit_price;type:decimal(10,2);comment:单价;not null;default:0.00;" json:"unit_price"`                     // 单价
	Description           string         `gorm:"column:description;type:text;comment:项目描述;default:NULL;" json:"description"`                                   // 项目描述
	IsMedicalInsurance    string         `gorm:"column:is_medical_insurance;type:varchar(10);comment:是否医保项目：否/是;default:否;" json:"is_medical_insurance"`       // 是否医保项目：否/是
	MedicalInsuranceRatio float64        `gorm:"column:medical_insurance_ratio;type:decimal(5,2);comment:医保报销比例;default:0.00;" json:"medical_insurance_ratio"` // 医保报销比例
	Status                string         `gorm:"column:status;type:varchar(10);comment:状态：停用/启用;not null;default:停用;" json:"status"`                           // 状态：停用/启用
	CreatedAt             time.Time      `gorm:"column:created_at;type:datetime(6);comment:创建时间;not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"`     // 创建时间
	UpdatedAt             time.Time      `gorm:"column:updated_at;type:datetime(6);comment:更新时间;not null;default:CURRENT_TIMESTAMP(6);" json:"updated_at"`     // 更新时间
	DeletedAt             gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);comment:删除时间;default:NULL;" json:"deleted_at"`                              // 删除时间
}

func (c *ChargeItems) TableName() string {
	return "charge_items"
}

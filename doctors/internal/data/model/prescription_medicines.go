package model

import "time"

type PrescriptionMedicines struct {
	Id             uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:ID;primaryKey;not null;" json:"id"`                            // ID
	PrescriptionId uint64    `gorm:"column:prescription_id;type:bigint UNSIGNED;comment:处方ID;not null;" json:"prescription_id"`           // 处方ID
	MedicineId     uint64    `gorm:"column:medicine_id;type:bigint UNSIGNED;comment:药品ID;not null;" json:"medicine_id"`                   // 药品ID
	Quantity       float64   `gorm:"column:quantity;type:decimal(10, 2);comment:数量;not null;" json:"quantity"`                            // 数量
	Unit           string    `gorm:"column:unit;type:varchar(20);comment:单位;not null;" json:"unit"`                                       // 单位
	UnitPrice      float64   `gorm:"column:unit_price;type:decimal(10, 2);comment:单价;not null;" json:"unit_price"`                        // 单价
	TotalPrice     float64   `gorm:"column:total_price;type:decimal(10, 2);comment:总价;not null;" json:"total_price"`                      // 总价
	Dosage         string    `gorm:"column:dosage;type:varchar(100);comment:用量;default:NULL;" json:"dosage"`                              // 用量
	Frequency      string    `gorm:"column:frequency;type:varchar(50);comment:频次;default:NULL;" json:"frequency"`                         // 频次
	Duration       string    `gorm:"column:duration;type:varchar(50);comment:疗程;default:NULL;" json:"duration"`                           // 疗程
	UsageMethod    string    `gorm:"column:usage_method;type:varchar(100);comment:用法;default:NULL;" json:"usage_method"`                  // 用法
	Notes          string    `gorm:"column:notes;type:varchar(500);comment:备注;default:NULL;" json:"notes"`                                // 备注
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
}

func (p *PrescriptionMedicines) TableName() string {
	return "prescription_medicines"
}

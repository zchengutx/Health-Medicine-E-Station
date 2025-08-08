package model

import "time"

type MedicineInventory struct {
	Id                uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:ID;primaryKey;not null;" json:"id"`                                    // ID
	MedicineId        uint64    `gorm:"column:medicine_id;type:bigint UNSIGNED;comment:药品ID;not null;" json:"medicine_id"`                           // 药品ID
	HospitalId        uint64    `gorm:"column:hospital_id;type:bigint UNSIGNED;comment:医院ID;not null;" json:"hospital_id"`                           // 医院ID
	BatchNumber       string    `gorm:"column:batch_number;type:varchar(50);comment:批次号;not null;" json:"batch_number"`                              // 批次号
	ProductionDate    time.Time `gorm:"column:production_date;type:date;comment:生产日期;default:NULL;" json:"production_date"`                          // 生产日期
	ExpiryDate        time.Time `gorm:"column:expiry_date;type:date;comment:有效期;default:NULL;" json:"expiry_date"`                                   // 有效期
	StockQuantity     float64   `gorm:"column:stock_quantity;type:decimal(10, 2);comment:库存数量;not null;default:0.00;" json:"stock_quantity"`         // 库存数量
	ReservedQuantity  float64   `gorm:"column:reserved_quantity;type:decimal(10, 2);comment:预留数量;not null;default:0.00;" json:"reserved_quantity"`   // 预留数量
	AvailableQuantity float64   `gorm:"column:available_quantity;type:decimal(10, 2);comment:可用数量;not null;default:0.00;" json:"available_quantity"` // 可用数量
	PurchasePrice     float64   `gorm:"column:purchase_price;type:decimal(10, 2);comment:进价;default:0.00;" json:"purchase_price"`                    // 进价
	SellingPrice      float64   `gorm:"column:selling_price;type:decimal(10, 2);comment:售价;default:0.00;" json:"selling_price"`                      // 售价
	Supplier          string    `gorm:"column:supplier;type:varchar(100);comment:供应商;default:NULL;" json:"supplier"`                                 // 供应商
	StorageLocation   string    `gorm:"column:storage_location;type:varchar(100);comment:存储位置;default:NULL;" json:"storage_location"`                // 存储位置
	MinStockLevel     float64   `gorm:"column:min_stock_level;type:decimal(10, 2);comment:最低库存;default:0.00;" json:"min_stock_level"`                // 最低库存
	MaxStockLevel     float64   `gorm:"column:max_stock_level;type:decimal(10, 2);comment:最高库存;default:0.00;" json:"max_stock_level"`                // 最高库存
	Status            string    `gorm:"column:status;type:varchar(20);comment:状态：停用/正常/过期/召回;not null;default:正常;" json:"status"`                    // 状态：停用/正常/过期/召回
	CreatedAt         time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`         // 创建时间
	UpdatedAt         time.Time `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"`         // 更新时间
}

func (m *MedicineInventory) TableName() string {
	return "medicine_inventory"
}

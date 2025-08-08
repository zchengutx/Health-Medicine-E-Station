package model

import "time"

type Medicines struct {
	Id                uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:药品ID;primaryKey;not null;" json:"id"`                                // 药品ID
	MedicineCode      string    `gorm:"column:medicine_code;type:varchar(32);comment:药品编码;not null;" json:"medicine_code"`                          // 药品编码
	Name              string    `gorm:"column:name;type:varchar(100);comment:药品名称;not null;" json:"name"`                                         // 药品名称
	GenericName       string    `gorm:"column:generic_name;type:varchar(100);comment:通用名;default:NULL;" json:"generic_name"`                       // 通用名
	BrandName         string    `gorm:"column:brand_name;type:varchar(100);comment:商品名;default:NULL;" json:"brand_name"`                          // 商品名
	Specification     string    `gorm:"column:specification;type:varchar(100);comment:规格;default:NULL;" json:"specification"`                      // 规格
	DosageForm        string    `gorm:"column:dosage_form;type:varchar(50);comment:剂型：片剂，胶囊，注射液等;default:NULL;" json:"dosage_form"`               // 剂型：片剂，胶囊，注射液等
	Manufacturer      string    `gorm:"column:manufacturer;type:varchar(100);comment:生产厂家;default:NULL;" json:"manufacturer"`                     // 生产厂家
	ApprovalNumber    string    `gorm:"column:approval_number;type:varchar(50);comment:批准文号;default:NULL;" json:"approval_number"`                 // 批准文号
	Category          string    `gorm:"column:category;type:varchar(50);comment:药品分类;default:NULL;" json:"category"`                              // 药品分类
	PrescriptionType  string    `gorm:"column:prescription_type;type:varchar(20);comment:处方类型：处方药/非处方药;default:处方药;" json:"prescription_type"`    // 处方类型：处方药/非处方药
	Unit              string    `gorm:"column:unit;type:varchar(20);comment:单位：盒，瓶，支等;default:NULL;" json:"unit"`                                // 单位：盒，瓶，支等
	Price             float64   `gorm:"column:price;type:decimal(10,2);comment:单价;default:0.00;" json:"price"`                                     // 单价
	Indications       string    `gorm:"column:indications;type:text;comment:适应症;default:NULL;" json:"indications"`                                // 适应症
	Contraindications string    `gorm:"column:contraindications;type:text;comment:禁忌症;default:NULL;" json:"contraindications"`                      // 禁忌症
	SideEffects       string    `gorm:"column:side_effects;type:text;comment:不良反应;default:NULL;" json:"side_effects"`                             // 不良反应
	DosageUsage       string    `gorm:"column:dosage_usage;type:text;comment:用法用量;default:NULL;" json:"dosage_usage"`                             // 用法用量
	StorageConditions string    `gorm:"column:storage_conditions;type:varchar(200);comment:储存条件;default:NULL;" json:"storage_conditions"`           // 储存条件
	ShelfLife         int32     `gorm:"column:shelf_life;type:int;comment:保质期（月）;default:NULL;" json:"shelf_life"`                                // 保质期（月）
	ImageUrl          string    `gorm:"column:image_url;type:varchar(255);comment:药品图片URL;default:NULL;" json:"image_url"`                        // 药品图片URL
	Status            string    `gorm:"column:status;type:varchar(20);comment:状态：停用/启用;not null;default:启用;" json:"status"`                        // 状态：停用/启用
	CreatedAt         time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`       // 创建时间
	UpdatedAt         time.Time `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"`       // 更新时间
}

func (m *Medicines)TableName() string {
	return "medicines"
}

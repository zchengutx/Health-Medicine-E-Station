package model

import (
	"time"

	"gorm.io/gorm"
)

type ChargeDetails struct {
	Id                     uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:ID;primaryKey;not null;" json:"id"`                                       // ID
	DetailNo               string         `gorm:"column:detail_no;type:varchar(32);comment:明细编号;not null;" json:"detail_no"`                                      // 明细编号
	PatientId              uint64         `gorm:"column:patient_id;type:bigint UNSIGNED;comment:患者ID;not null;" json:"patient_id"`                                // 患者ID
	DoctorId               uint64         `gorm:"column:doctor_id;type:bigint UNSIGNED;comment:医生ID;not null;" json:"doctor_id"`                                  // 医生ID
	MedicalRecordId        uint64         `gorm:"column:medical_record_id;type:bigint UNSIGNED;comment:病历ID;default:NULL;" json:"medical_record_id"`              // 病历ID
	PrescriptionId         uint64         `gorm:"column:prescription_id;type:bigint UNSIGNED;comment:处方ID;default:NULL;" json:"prescription_id"`                  // 处方ID
	ChargeItemId           uint64         `gorm:"column:charge_item_id;type:bigint UNSIGNED;comment:收费项目ID;not null;" json:"charge_item_id"`                      // 收费项目ID
	ChargeDate             time.Time      `gorm:"column:charge_date;type:date;comment:收费日期;not null;" json:"charge_date"`                                         // 收费日期
	Quantity               float64        `gorm:"column:quantity;type:decimal(10, 2);comment:数量;not null;default:1.00;" json:"quantity"`                          // 数量
	UnitPrice              float64        `gorm:"column:unit_price;type:decimal(10, 2);comment:单价;not null;" json:"unit_price"`                                   // 单价
	TotalAmount            float64        `gorm:"column:total_amount;type:decimal(10, 2);comment:总金额;not null;" json:"total_amount"`                              // 总金额
	DiscountAmount         float64        `gorm:"column:discount_amount;type:decimal(10, 2);comment:优惠金额;default:0.00;" json:"discount_amount"`                   // 优惠金额
	ActualAmount           float64        `gorm:"column:actual_amount;type:decimal(10, 2);comment:实际金额;not null;" json:"actual_amount"`                           // 实际金额
	MedicalInsuranceAmount float64        `gorm:"column:medical_insurance_amount;type:decimal(10, 2);comment:医保金额;default:0.00;" json:"medical_insurance_amount"` // 医保金额
	SelfPayAmount          float64        `gorm:"column:self_pay_amount;type:decimal(10, 2);comment:自费金额;not null;" json:"self_pay_amount"`                       // 自费金额
	PaymentStatus          string         `gorm:"column:payment_status;type:varchar(20);comment:支付状态：未支付/已支付/已退费;not null;default:未支付;" json:"payment_status"`    // 支付状态：未支付/已支付/已退费
	PaymentMethod          string         `gorm:"column:payment_method;type:varchar(20);comment:支付方式：现金/刷卡/支付宝/微信;default:NULL;" json:"payment_method"`           // 支付方式：现金/刷卡/支付宝/微信
	PaymentTime            time.Time      `gorm:"column:payment_time;type:timestamp;comment:支付时间;default:NULL;" json:"payment_time"`                              // 支付时间
	RefundAmount           float64        `gorm:"column:refund_amount;type:decimal(10, 2);comment:退费金额;default:0.00;" json:"refund_amount"`                       // 退费金额
	RefundTime             time.Time      `gorm:"column:refund_time;type:timestamp;comment:退费时间;default:NULL;" json:"refund_time"`                                // 退费时间
	Notes                  string         `gorm:"column:notes;type:varchar(500);comment:备注;default:NULL;" json:"notes"`                                           // 备注
	CreatedAt              time.Time      `gorm:"column:created_at;type:datetime(6);comment:创建时间;not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"`       // 创建时间
	UpdatedAt              time.Time      `gorm:"column:updated_at;type:datetime(6);comment:更新时间;not null;default:CURRENT_TIMESTAMP(6);" json:"updated_at"`       // 更新时间
	DeletedAt              gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6);default:NULL;" json:"deleted_at"`
}

func (c *ChargeDetails) TableName() string {
	return "charge_details"
}

// 自动生成模板MtOrdersDrug
package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mtOrdersDrug表 结构体  MtOrdersDrug
type MtOrdersDrug struct {
	global.GVA_MODEL
	OrderId     *int    `json:"orderId" form:"orderId" gorm:"comment:订单id;column:order_id;size:19;" binding:"required"`                                          //订单id
	DrugId      *int    `json:"drugId" form:"drugId" gorm:"comment:药品id;column:drug_id;size:19;" binding:"required"`                                             //药品id
	UserId      *int    `json:"userId" form:"userId" gorm:"comment:患者id;column:user_id;size:19;" binding:"required"`                                             //患者id
	Quantity    *int    `json:"quantity" form:"quantity" gorm:"comment:数量;column:quantity;size:10;" binding:"required"`                                          //数量
	OrderStatus *string `json:"orderStatus" form:"orderStatus" gorm:"comment:订单状态:1-待发货，2-待收货，3-已收货;column:order_status;size:10;" binding:"required"` //订单状态:1-待发货，2-待收货，3-已收货
	CreatedBy   uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint    `gorm:"column:deleted_by;comment:删除者"`

	Order string `gorm:"-" json:"order"`
	Drug  string `gorm:"-" json:"drug"`
	User  string `gorm:"-" json:"user"`
}

// TableName mtOrdersDrug表 MtOrdersDrug自定义表名 mt_orders_drug
func (MtOrdersDrug) TableName() string {
	return "mt_orders_drug"
}

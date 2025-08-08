
// 自动生成模板MtOrders
package medicine
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mtOrders表 结构体  MtOrders
type MtOrders struct {
    global.GVA_MODEL
  OrderNo  *string `json:"orderNo" form:"orderNo" gorm:"comment:订单编号;column:order_no;size:50;" binding:"required"`  //订单编号
  UserName  *string `json:"userName" form:"userName" gorm:"comment:患者姓名;column:user_name;size:50;" binding:"required"`  //患者姓名
  UserPhone  *string `json:"userPhone" form:"userPhone" gorm:"comment:患者电话;column:user_phone;size:20;" binding:"required"`  //患者电话
  DoctorName  *string `json:"doctorName" form:"doctorName" gorm:"comment:医生姓名;column:doctor_name;size:50;" binding:"required"`  //医生姓名
  AddressDetail  *string `json:"addressDetail" form:"addressDetail" gorm:"comment:地址详情;column:address_detail;size:200;" binding:"required"`  //地址详情
  TotalAmount  *float64 `json:"totalAmount" form:"totalAmount" gorm:"comment:总金额;column:total_amount;" binding:"required"`  //总金额
  PayType  *string `json:"payType" form:"payType" gorm:"comment:支付方式：1-微信，2-支付宝，3-银行卡;column:pay_type;size:10;" binding:"required"`  //支付方式：1-微信，2-支付宝，3-银行卡
  Status  *string `json:"status" form:"status" gorm:"comment:订单状态：1-待支付，2-已支付，3-配药中，4-已发货，5-已完成，6-已取消;column:status;size:20;" binding:"required"`  //订单状态：1-待支付，2-已支付，3-配药中，4-已发货，5-已完成，6-已取消
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName mtOrders表 MtOrders自定义表名 mt_orders
func (MtOrders) TableName() string {
    return "mt_orders"
}






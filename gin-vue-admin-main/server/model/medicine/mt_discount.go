
// 自动生成模板MtDiscount
package medicine
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// mtDiscount表 结构体  MtDiscount
type MtDiscount struct {
    global.GVA_MODEL
  DiscountName  *string `json:"discountName" form:"discountName" gorm:"comment:优惠券名称;column:discount_name;size:50;" binding:"required"`  //优惠券名称
  Classify  *string `json:"classify" form:"classify" gorm:"comment:券的来源分类;column:classify;size:20;" binding:"required"`  //券的来源分类
  StoreId  *int `json:"storeId" form:"storeId" gorm:"comment:店铺id，平台卷为null;column:store_id;size:10;" binding:"required"`  //店铺id，平台卷为null
  DiscountAmout  *float64 `json:"discountAmout" form:"discountAmout" gorm:"comment:优惠金额;column:discount_amout;" binding:"required"`  //优惠金额
  MinOrderAmount  *float64 `json:"minOrderAmount" form:"minOrderAmount" gorm:"comment:最低消费门槛;column:min_order_amount;" binding:"required"`  //最低消费门槛
  StartTime  *time.Time `json:"startTime" form:"startTime" gorm:"comment:有效期开始;column:start_time;" binding:"required"`  //有效期开始
  EndTime  *time.Time `json:"endTime" form:"endTime" gorm:"comment:有效期结束;column:end_time;" binding:"required"`  //有效期结束
  MaxIssue  *int `json:"maxIssue" form:"maxIssue" gorm:"comment:总发行量，0=不限;column:max_issue;" binding:"required"`  //总发行量，0=不限
  MaxPerUser  *bool `json:"maxPerUser" form:"maxPerUser" gorm:"comment:每人限领;column:max_per_user;" binding:"required"`  //每人限领
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName mtDiscount表 MtDiscount自定义表名 mt_discount
func (MtDiscount) TableName() string {
    return "mt_discount"
}






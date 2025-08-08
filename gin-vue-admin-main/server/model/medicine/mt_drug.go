// 自动生成模板MtDrug
package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mtDrug表 结构体  MtDrug
type MtDrug struct {
	global.GVA_MODEL
	DrugName      *string  `json:"drugName" form:"drugName" gorm:"comment:药品名称;column:drug_name;size:20;" binding:"required"`             //药品名称
	Guide         *int     `json:"guide" form:"guide" gorm:"comment:用药指导id;column:guide;size:19;" binding:"required"`                     //用药指导id
	Explain       *int     `json:"explain" form:"explain" gorm:"comment:说明书id;column:explain;size:19;" binding:"required"`                //说明书id
	Specification *string  `json:"specification" form:"specification" gorm:"comment:规格;column:specification;size:20;" binding:"required"` //规格
	Price         *float64 `json:"price" form:"price" gorm:"comment:价格;column:price;size:22;" binding:"required"`                         //价格
	SalesVolume   *float64 `json:"salesVolume" form:"salesVolume" gorm:"comment:销量;column:sales_volume;size:22;" binding:"required"`      //销量
	Inventory     *int     `json:"inventory" form:"inventory" gorm:"comment:库存;column:inventory;size:19;" binding:"required"`             //库存
	Status        *string  `json:"status" form:"status" gorm:"comment:状态;column:status;size:10;" binding:"required"`                      //状态
	CreatedBy     uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy     uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy     uint     `gorm:"column:deleted_by;comment:删除者"`

	Guides   string `json:"guides" gorm:"-"`
	Explains string `json:"explains" gorm:"-"`
}

// TableName mtDrug表 MtDrug自定义表名 mt_drug
func (MtDrug) TableName() string {
	return "mt_drug"
}

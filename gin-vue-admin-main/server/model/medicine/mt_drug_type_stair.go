package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MtDrugTypeStair 药品类型阶梯表 结构体
type MtDrugTypeStair struct {
	global.GVA_MODEL
	Name        *string `json:"name" form:"name" gorm:"comment:阶梯名称;column:name;size:50;" binding:"required"`     //阶梯名称
	Code        *string `json:"code" form:"code" gorm:"comment:阶梯代码;column:code;size:20;" binding:"required"`     //阶梯代码
	Description *string `json:"description" form:"description" gorm:"comment:描述;column:description;size:200;"`    //描述
	Sort        *int    `json:"sort" form:"sort" gorm:"comment:排序;column:sort;" binding:"required"`               //排序
	Status      *string `json:"status" form:"status" gorm:"comment:状态;column:status;size:10;" binding:"required"` //状态
	CreatedBy   uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName MtDrugTypeStair自定义表名 mt_drug_type_stair
func (MtDrugTypeStair) TableName() string {
	return "mt_drug_type_stair"
}

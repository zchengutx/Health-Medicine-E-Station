
// 自动生成模板MtUser
package medicine
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mtUser表 结构体  MtUser
type MtUser struct {
    global.GVA_MODEL
  NickName  *string `json:"nickName" form:"nickName" gorm:"comment:昵称;column:nick_name;size:20;" binding:"required"`  //昵称
  Mobile  *string `json:"mobile" form:"mobile" gorm:"comment:手机号;column:mobile;" binding:"required"`  //手机号
  Avatar  *string `json:"avatar" form:"avatar" gorm:"comment:头像;column:avatar;size:255;"`  //头像
  Status  *string `json:"status" form:"status" gorm:"default:1;comment:用户账号状态;column:status;size:10;" binding:"required"`  //用户账号状态
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName mtUser表 MtUser自定义表名 mt_user
func (MtUser) TableName() string {
    return "mt_user"
}






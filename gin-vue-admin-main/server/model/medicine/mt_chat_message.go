// 自动生成模板MtChatMessage
package medicine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// mtChatMessage表 结构体  MtChatMessage
type MtChatMessage struct {
	global.GVA_MODEL
	FromId      *int    `json:"fromId" form:"fromId" gorm:"comment:发送者id;column:from_id;size:10;" binding:"required"`                //发送者id
	ToId        *int    `json:"toId" form:"toId" gorm:"comment:接收者id;column:to_id;size:10;" binding:"required"`                      //接收者id
	Content     *string `json:"content" form:"content" gorm:"comment:消息内容;column:content;size:50;" binding:"required"`              //消息内容
	MessageType *string `json:"messageType" form:"messageType" gorm:"comment:消息类型;column:message_type;size:20;" binding:"required"` //消息类型
	RoomId      *int    `json:"roomId" form:"roomId" gorm:"comment:房间id;column:room_id;size:10;" binding:"required"`                  //房间id
	CreatedBy   uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName mtChatMessage表 MtChatMessage自定义表名 mt_chat_message
func (MtChatMessage) TableName() string {
	return "mt_chat_message"
}

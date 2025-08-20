package medicine

import "time"

type MtChatRoom struct {
	Id              int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	RoomId          string    `gorm:"column:room_id;type:varchar(50);comment:房间ID;default:NULL;" json:"room_id"`                                       // 房间ID
	RoomName        string    `gorm:"column:room_name;type:varchar(50);comment:房间名称;default:NULL;" json:"room_name"`                                   // 房间名称
	User1Id         int32     `gorm:"column:user1_id;type:int;comment:用户1id;not null;default:0;" json:"user1_id"`                                      // 用户1id
	User2Id         int32     `gorm:"column:user2_id;type:int;comment:用户2id;not null;default:0;" json:"user2_id"`                                      // 用户2id
	LastMessage     string    `gorm:"column:last_message;type:varchar(50);comment:最后一条消息;default:NULL;" json:"last_message"`                           // 最后一条消息
	LastMessageTime time.Time `gorm:"column:last_message_time;type:datetime(6);comment:最后消息时间;default:CURRENT_TIMESTAMP(6);" json:"last_message_time"` // 最后消息时间
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP(6);" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP(6);" json:"updated_at"`
}

func (MtChatRoom) TableName() string {
	return "mt_char_room"
}

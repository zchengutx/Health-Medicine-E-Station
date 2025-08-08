package model

import "time"

type Greeter struct {
	Id        int64     `gorm:"column:id;type:bigint;primaryKey;not null;" json:"id"`
	Hello     string    `gorm:"column:hello;type:varchar(255);not null;" json:"hello"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"updated_at"`
}

func (g *Greeter) TableName() string {
	return "greeter"
}

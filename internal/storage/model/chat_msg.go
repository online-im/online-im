package model

import "gorm.io/gorm"

type RoomType int

const (
	ChatMsgTableName = "chat_msg"

	RoomTypeUser     RoomType = 0
	RoomTypeChatRoom RoomType = 1
)

// ChatMsg store message information and it's chat information
type ChatMsg struct {
	gorm.Model
	Content  string   `gorm:"column:content;not null"`
	From     int64    `gorm:"column:from;not null"`
	To       int64    `gorm:"column:to;not null"`
	RoomType RoomType `gorm:"column:room_type"`
}

func (ChatMsg) TableName() string {
	return ChatMsgTableName
}

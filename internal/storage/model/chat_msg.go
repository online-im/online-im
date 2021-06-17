package model

import (
	"github.com/online-im/online-im/pkg/constant"
	"gorm.io/gorm"
)

const (
	ChatMsgTableName = "chat_msg"
)

// ChatMsg store message information and it's chat information
type ChatMsg struct {
	gorm.Model
	Content  string               `gorm:"column:content;not null"`
	From     int64                `gorm:"column:from;not null"`
	To       int64                `gorm:"column:to;not null"`
	RoomType constant.PublishType `gorm:"column:room_type"`
}

func (ChatMsg) TableName() string {
	return ChatMsgTableName
}

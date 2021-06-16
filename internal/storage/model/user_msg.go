package model

import "gorm.io/gorm"

const (
	UserMsgTableName = "user_msg"
)

// UserMsg stores relationship between user(not include chatroom) and message
type UserMsg struct {
	gorm.Model
	Receiver  int64 `gorm:"column:receiver;not null"`
	MsgID     int64 `gorm:"column:msg_id;not null"`
	ChatMsgID int64 `gorm:"column:chat_msg_id;not null"`
}

func (UserMsg) TableName() string {
	return UserMsgTableName
}

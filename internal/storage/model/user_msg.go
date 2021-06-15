package model

import "gorm.io/gorm"

const (
	UserMsgTableName = "user_msg"
)

// UserMsg stores relationship between user and message
type UserMsg struct {
	gorm.Model
	UserID int64 `gorm:"column:user_id;not null"`
	MsgID  int64 `gorm:"column:msg_id;not null"`
}

func (UserMsg) TableName()string {
	return UserMsgTableName
}
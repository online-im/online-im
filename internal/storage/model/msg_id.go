package model

import "gorm.io/gorm"

const (
	MsgIDTableName = "msg_id"
)

// MsgID stores currently next msg id of receiver user(not chatroom)
type MsgID struct {
	gorm.Model
	Receiver int64 `gorm:"column:receiver;not null"`
	MaxMsgID int64 `gorm:"column:max_msg_id;not null"`
}

func (MsgID) TableName() string {
	return MsgIDTableName
}

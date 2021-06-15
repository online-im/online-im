package dal

import (
	"github.com/online-im/online-im/internal/storage/model"
	"gorm.io/gorm"
)

type Dal interface {
	GetMessageWithID(db *gorm.DB, msgID int64) (*model.ChatMsg, error)
	GetOwnReceiveMessage(db *gorm.DB, userID int64, fromMsgID int64, limit int) ([]model.UserMsg, error)
	// AddMessage add a chat message
	AddMessage(db *gorm.DB, msg *model.ChatMsg) error
	// SendMessageTo store the message to be sent to the user
	SendMessageTo(db *gorm.DB, sendInfo *model.UserMsg) error
}


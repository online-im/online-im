package dal

import (
	"github.com/online-im/online-im/internal/storage/model"
	"gorm.io/gorm"
)

type Dal interface {
	GetMsgID(db *gorm.DB, receiver int64) (int64, error)
	BatchGetMessageContent(db *gorm.DB, chatMsgIDs []int64) ([]model.ChatMsg, error)
	GetMessageContent(db *gorm.DB, chatID, from int64, limit int) ([]model.ChatMsg, error)
	GetMessages(db *gorm.DB, receiver, start int64, limit int) ([]model.UserMsg, error)
	// AddMessageContent add a chat message
	AddMessageContent(db *gorm.DB, msg *model.ChatMsg) (int64, error)
	// AddMessage store the message to be sent to the user
	AddMessage(db *gorm.DB, msg *model.UserMsg) error
}

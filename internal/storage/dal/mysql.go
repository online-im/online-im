package dal

import (
	"github.com/online-im/online-im/internal/storage/model"
	"gorm.io/gorm"
)

type mysqlDal struct{}

func (d *mysqlDal) GetMessageWithID(db *gorm.DB, msgID int64) (*model.ChatMsg, error) {
	msg := &model.ChatMsg{}
	err := db.Model(msg).Where("id = ?", msgID).First(msg).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return msg, nil
}

func (d *mysqlDal) GetOwnReceiveMessage(db *gorm.DB,
	userID int64, fromMsgID int64, limit int) ([]model.UserMsg, error) {
	msgs := make([]model.UserMsg, 0)
	err := db.Model(&model.UserMsg{}).
		Where("user_id = ? AND msg_id > ?", userID, fromMsgID).
		Limit(limit).Find(msgs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return msgs, nil
}

// AddMessage add a chat message
func (d *mysqlDal) AddMessage(db *gorm.DB, msg *model.ChatMsg) error {
	return db.Create(msg).Error
}

// SendMessageTo store the message to be sent to the user
func (d *mysqlDal) SendMessageTo(db *gorm.DB, sendInfo *model.UserMsg) error {
	return db.Create(sendInfo).Error
}

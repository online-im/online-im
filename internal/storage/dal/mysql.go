package dal

import (
	"github.com/online-im/online-im/internal/storage/model"
	"gorm.io/gorm"
)

const (
	BatchSize = 200
)

type mysqlDal struct{}

func NewMysqlDal() Dal {
	return &mysqlDal{}
}

func (d *mysqlDal) GetMsgID(db *gorm.DB, receiver int64) (int64, error) {
	var msgID int64
	if err := db.Transaction(func(tx *gorm.DB) error {
		msgIDInfo := &model.MsgID{}
		if err := tx.Model(msgIDInfo).Where("receiver = ?", receiver).
			FirstOrCreate(msgIDInfo).Error; err != nil {
			return err
		}
		msgIDInfo.MaxMsgID += 1
		if err := tx.Model(msgIDInfo).Where("receiver = ?", receiver).
			Update("max_msg_id", msgIDInfo.MaxMsgID).Error; err != nil {
			return err
		}
		msgID = msgIDInfo.MaxMsgID - 1
		return nil
	}); err != nil {
		return -1, err
	}
	return msgID, nil
}

func (d *mysqlDal) BatchGetMessageContent(db *gorm.DB, chatMsgIDs []int64) ([]model.ChatMsg, error) {
	msgs := make([]model.ChatMsg, 0)
	err := db.Model(&model.ChatMsg{}).Where("msg_id in ?", chatMsgIDs).Find(msgs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return msgs, nil
}

func (d *mysqlDal) GetMessageContent(db *gorm.DB, chatID,
	start int64, limit int) ([]model.ChatMsg, error) {
	msgs := make([]model.ChatMsg, 0)
	err := db.Model(&model.ChatMsg{}).Where("to = ? AND id > ?", chatID, start).Limit(limit).Find(msgs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return msgs, nil
}

func (d *mysqlDal) GetMessages(db *gorm.DB, receiver,
	start int64, limit int) ([]model.UserMsg, error) {
	msgs := make([]model.UserMsg, 0)
	err := db.Model(&model.UserMsg{}).
		Where("receiver = ? AND msg_id > ?", receiver, start).
		Limit(limit).Find(msgs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return msgs, nil
}

func (d *mysqlDal) AddMessageContent(db *gorm.DB, msg *model.ChatMsg) (int64, error) {
	return int64(msg.ID), db.Create(msg).Error
}

func (d *mysqlDal) AddMessage(db *gorm.DB, msg *model.UserMsg) error {
	return db.Create(msg).Error
}

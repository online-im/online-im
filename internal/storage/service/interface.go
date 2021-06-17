package service

import (
	"context"

	"github.com/online-im/online-im/internal/storage/model"
	"github.com/online-im/online-im/pkg/constant"
)

type PagerConfig struct {
	StartMsgID int64
	Limit      int
}

type Service interface {
	// SendMsgTo used to store messages to be sent, and returns the id of message content
	SendMsgTo(ctx context.Context, from, to int64, content string, chatType constant.PublishType) (int64, error)
	// ReadMsgs load all chat messages which received by given user
	ReadMsgs(ctx context.Context, userID int64,
		pageInfo *PagerConfig) (map[int64]model.ChatMsg, int64, error)
	// GetChatMsgs get messages in given chat
	GetChatMsgs(ctx context.Context, chatID int64, chatType constant.PublishType,
		pageInfo *PagerConfig) ([]model.ChatMsg, int64, error)
}

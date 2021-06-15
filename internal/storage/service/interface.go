package service

import (
	"context"

	"github.com/online-im/online-im/internal/storage/model"
)

type PagerConfig struct {
	StartMsgID int64
	Limit      int
}

type Service interface {
	// SendMsgTo used to store messages to be sent
	SendMsgTo(ctx context.Context, from, to int64, content string, chatType model.RoomType) (int64, error)
	// ReadMsgs load all chat messages which received by given user
	ReadMsgs(ctx context.Context, userID int64,
		pageInfo *PagerConfig) ([]model.ChatMsg, error)
	// GetChatMsgs can be used to get messages in given chat
	GetChatMsgs(ctx context.Context, userID int64, chatType model.RoomType,
		pageInfo *PagerConfig) ([]model.ChatMsg, error)
}

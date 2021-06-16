package service

import (
	"context"
	"sync"

	"github.com/glory-go/glory/log"
	"github.com/online-im/online-im/internal/storage/dal"
	"github.com/online-im/online-im/internal/storage/model"
	"github.com/online-im/online-im/pkg/constant"
	"gorm.io/gorm"
)

type serviceStruct struct {
	dalClient dal.Dal
}

func (s *serviceStruct) SendMsgTo(ctx context.Context, from, to int64, content string, chatType model.RoomType) (int64, error) {
	funcName := "SendMsgTo"
	// store message content first(no tx)
	msgContentID, err := s.dalClient.AddMessageContent(dal.DB, &model.ChatMsg{
		Content:  content,
		From:     from,
		To:       to,
		RoomType: chatType,
	})
	if err != nil {
		log.CtxErrorf(ctx, "[%v] fail to save msg content, receiver %v of type %v with error %v",
			funcName, to, chatType, err)
		return -1, constant.CoreErrorCode_DB
	}
	// store message for all receiver(tx)
	if err := dal.DB.Transaction(func(tx *gorm.DB) error {
		userIDs := make([]int64, 0)
		switch chatType {
		case model.RoomTypeUser:
			// if receiver is user, store to given user
			userIDs = append(userIDs, to)
		case model.RoomTypeChatRoom:
			// TODO: get all chatroom users
		}
		if len(userIDs) == 0 {
			return nil
		}
		// insert message for every receiver to db
		var err1 error
		var wg sync.WaitGroup
		for _, userID := range userIDs {
			go func() {
				wg.Add(1)
				defer wg.Done()
				// get msg id
				msgID, err := s.dalClient.GetMsgID(dal.DB, userID)
				if err != nil {
					log.CtxErrorf(ctx, "[%v] fail to get msg id for user %v of type %v with error %v",
						funcName, userID, chatType, err)
					err1 = err
				}
				if err := s.dalClient.AddMessage(tx, &model.UserMsg{
					Receiver:  userID,
					MsgID:     msgID,
					ChatMsgID: msgContentID,
				}); err != nil {
					log.CtxErrorf(ctx, "[%v] fail to save msg for all receiver user, msg id %v, sender %v, receiver %v,  type %v with error %v",
						funcName, msgID, from, to, chatType, err)
					err1 = constant.CoreErrorCode_DB
				}
			}()
		}
		wg.Wait()
		return err1
	}); err != nil {
		return -1, err
	}
	return msgContentID, nil
}

func (s *serviceStruct) ReadMsgs(ctx context.Context, userID int64,
	pageInfo *PagerConfig) (map[int64]model.ChatMsg, int64, error) {
	funcName := "ReadMsgs"
	userMsgs, err := s.dalClient.GetMessages(dal.DB, userID, pageInfo.StartMsgID, pageInfo.Limit)
	if err != nil {
		log.CtxErrorf(ctx, "[%v] fail to get receive msg for user %v with error %v",
			funcName, userID, err)
		return nil, -1, constant.CoreErrorCode_DB
	}
	chatMsgIDs := make([]int64, 0)
	for _, msg := range userMsgs {
		chatMsgIDs = append(chatMsgIDs, msg.ChatMsgID)
	}
	if len(chatMsgIDs) == 0 {
		return map[int64]model.ChatMsg{}, pageInfo.StartMsgID, nil
	}
	// get message content
	msgs, err := s.dalClient.BatchGetMessageContent(dal.DB, chatMsgIDs)
	if err != nil {
		log.CtxErrorf(ctx, "[%v] fail to get msg content for user %v with error %v",
			funcName, userID, err)
		return nil, -1, constant.CoreErrorCode_DB
	}
	chatMsgMap := make(map[int64]model.ChatMsg)
	for _, msg := range msgs {
		chatMsgMap[int64(msg.ID)] = msg
	}
	// get return data
	msgMap := make(map[int64]model.ChatMsg)
	for _, msg := range userMsgs {
		chatMsg, ok := chatMsgMap[msg.ChatMsgID]
		if !ok {
			continue
		}
		msgMap[msg.MsgID] = chatMsg
	}
	return msgMap, userMsgs[len(userMsgs)-1].MsgID, nil
}

func (s *serviceStruct) GetChatMsgs(ctx context.Context, userID int64, chatType model.RoomType,
	pageInfo *PagerConfig) ([]model.ChatMsg, int64, error) {
	funcName := "GetChatMsgs"
	msgs, err := s.dalClient.GetMessageContent(dal.DB, userID, pageInfo.StartMsgID, pageInfo.Limit)
	if err != nil {
		log.CtxErrorf(ctx, "[%v] fail to get msg content for chat %v with error %v",
			funcName, userID, err)
		return msgs, -1, constant.CoreErrorCode_DB
	}
	return msgs, int64(msgs[len(msgs)-1].ID), nil
}

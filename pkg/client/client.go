package client

import (
	"github.com/glory-go/glory/log"
	"github.com/online-im/online-im/pkg/constant"
	"github.com/online-im/online-im/pkg/message"
	"golang.org/x/net/websocket"
)

type IMCoreClient struct {
	userID string
	ws     *websocket.Conn
}

func NewImCoreClient(targetAddr, userID string) (*IMCoreClient, error) {
	ws, err := websocket.Dial("ws://"+targetAddr+"/online?fromid="+userID, "", "http://"+targetAddr+"/")
	if err != nil {
		log.Errorf("websocket dial instance addr = %s failed!, with err = %s", targetAddr, err)
		return nil, err
	}

	return &IMCoreClient{
		userID: userID,
		ws:     ws,
	}, nil
}

func (i *IMCoreClient) SendMessage(data []byte, targetID string, pubType constant.PublishType) error {
	cmsg := message.CoreMessage{
		Type: constant.CoreMessageType_Message,
		MessagePayload: message.CoreMessagePayload{
			Data:        data,
			FromID:      i.userID,
			TargetID:    targetID,
			PublishType: pubType,
		},
	}
	return websocket.JSON.Send(i.ws, cmsg)
}

func (i *IMCoreClient) RecvMessage() (*message.CoreMessage, error) {
	recv := message.CoreMessage{}
	if err := websocket.JSON.Receive(i.ws, &recv); err != nil {
		return nil, err
	}
	return &recv, nil
}

func (i *IMCoreClient) Close() {
	cmsg := message.CoreMessage{
		Type: constant.CoreMessageType_Close,
	}
	websocket.JSON.Send(i.ws, cmsg)
	i.ws.Close()
}

func (i *IMCoreClient) CreateGroup(userID, groupID string, users []string) error {
	cmsg := message.CoreMessage{
		Type: constant.CoreMessageType_Manage,
		ManagePayload: message.ManagePayload{
			ManageType: constant.ManageType_CreateGroup,
			User:       userID,
			GroupID:    groupID,
			Users:      users,
		},
	}
	return websocket.JSON.Send(i.ws, cmsg)
}

func (i *IMCoreClient) DeleteGroup(groupID string) error {
	cmsg := message.CoreMessage{
		Type: constant.CoreMessageType_Manage,
		ManagePayload: message.ManagePayload{
			ManageType: constant.ManageType_DeleteGroup,
			User:       groupID,
		},
	}
	return websocket.JSON.Send(i.ws, cmsg)
}

func (i *IMCoreClient) AddUserToGroup(userID, groupID string) error {
	cmsg := message.CoreMessage{
		Type: constant.CoreMessageType_Manage,
		ManagePayload: message.ManagePayload{
			ManageType: constant.ManageType_AddUserToGroup,
			GroupID:    groupID,
			User:       userID,
		},
	}
	return websocket.JSON.Send(i.ws, cmsg)
}

func (i *IMCoreClient) RemoveUserFromGroup(userID, groupID string) error {
	cmsg := message.CoreMessage{
		Type: constant.CoreMessageType_Manage,
		ManagePayload: message.ManagePayload{
			ManageType: constant.ManageType_RemoveUserFromGroup,
			User:       userID,
			GroupID:    groupID,
		},
	}
	return websocket.JSON.Send(i.ws, cmsg)
}

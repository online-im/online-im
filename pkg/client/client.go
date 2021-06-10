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

func (i *IMCoreClient) SendMessage(data []byte, targetID string) error {
	cmsg := message.CoreMessage{
		Type: constant.CoreMessageType_Message,
		Payload: message.CoreMessagePayload{
			Data:     data,
			FromID:   i.userID,
			TargetID: targetID,
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
	i.ws.Close()
}

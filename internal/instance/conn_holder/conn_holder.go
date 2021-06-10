package conn_holder

import (
	"github.com/glory-go/glory/log"
	"github.com/gorilla/websocket"
	"github.com/online-im/online-im/pkg/constant"
	"github.com/online-im/online-im/pkg/message"
	perrors "github.com/pkg/errors"
	"sync"
)

var once *sync.Once
var wsConnHolder *WSConnHolder

func init() {
	once = &sync.Once{}
}

type WSConnHolder struct {
	connMap *sync.Map
}

func GetWSConnHolder() *WSConnHolder {
	once.Do(func() {
		wsConnHolder = &WSConnHolder{
			connMap: &sync.Map{},
		}
	})
	return wsConnHolder
}

func (ws *WSConnHolder) Add(userID string, conn *websocket.Conn) {
	ws.connMap.Store(userID, conn)
}

func (ws *WSConnHolder) Del(userID string) {
	ws.connMap.Delete(userID)
}

// send
func (ws *WSConnHolder) SendMessageWithData(fromID, targetID string, data []byte) error {
	msg := message.CoreMessage{
		Type: constant.CoreMessageType_Message,
		MessagePayload: message.CoreMessagePayload{
			Data:     data,
			FromID:   fromID,
			TargetID: targetID,
		},
	}

	val, ok := ws.connMap.Load(targetID)
	if !ok {
		log.Errorf("ws holder has no conn with userid = %s", targetID)
		return perrors.Errorf("ws holder has no conn with userid = %s", targetID)
	}
	conn, _ := val.(*websocket.Conn)
	err := websocket.WriteJSON(conn, msg)
	if err != nil {
		return perrors.Errorf("websocket.WriteJSON error = %v", err)
	}
	return nil
}

func (ws *WSConnHolder) SendMessageWithErr(targetID string, code constant.CoreErrorCode, msg string) error {
	cmsg := message.CoreMessage{
		Type: constant.CoreMessageType_ErrorMessage,
		Err: message.CoreErrorMessage{
			Code:    code,
			Message: msg,
		},
	}

	val, ok := ws.connMap.Load(targetID)
	if !ok {
		log.Errorf("ws holder has no conn with userid = %s", targetID)
		return perrors.Errorf("ws holder has no conn with userid = %s", targetID)
	}
	conn, _ := val.(*websocket.Conn)
	return websocket.WriteJSON(conn, cmsg)
}

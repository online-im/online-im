package conn_holder

import (
	"git.go-online.org.cn/Glory/glory/log"
	"github.com/goonline/online-im/pkg/message"
	"github.com/gorilla/websocket"
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
func (ws *WSConnHolder) Send(userID string, msg message.Message) error {
	val, ok := ws.connMap.Load(userID)
	if !ok {
		log.Errorf("ws holder has no conn with userid = %s", userID)
		return perrors.Errorf("ws holder has no conn with userid = %s", userID)
	}
	conn, _ := val.(*websocket.Conn)
	return websocket.WriteJSON(conn, msg)
}

package handler

import (
	gostNet "github.com/dubbogo/gost/net"
	"github.com/glory-go/glory/config"
	ghttp "github.com/glory-go/glory/http"
	"github.com/glory-go/glory/log"
	"github.com/gorilla/websocket"
	"github.com/online-im/online-im/internal/constant"
	"github.com/online-im/online-im/internal/instance/conn_holder"
	"github.com/online-im/online-im/internal/instance/ierror"
	"github.com/online-im/online-im/internal/instance/manager"
	"github.com/online-im/online-im/internal/instance/publisher"
	"github.com/online-im/online-im/internal/redis_client"
	pconst "github.com/online-im/online-im/pkg/constant"
	"github.com/online-im/online-im/pkg/message"
	"strconv"
)

var LocalIP string

func init() {
	var err error
	LocalIP, err = gostNet.GetLocalIP()
	if err != nil {
		panic(err)
	}
}

func OnlineHandler(c *ghttp.GRegisterWSController) {
	// register id to instance
	if err := c.R.ParseForm(); err != nil {
		panic(err)
	}
	fromID := c.R.Form.Get("fromid")
	log.Infof("get from id = %s", fromID)
	redis_client.GetIMRedisClientInstance().SetUserID2instanceIP(fromID, LocalIP+":"+strconv.Itoa(config.GlobalServerConf.ServiceConfigs[constant.GRPCProviderName].Port))
	conn_holder.GetWSConnHolder().Add(fromID, c.WSConn)
	defer func() {
		// remove userID from conn holder
		conn_holder.GetWSConnHolder().Del(fromID)

		// remove userID -> instance ip map from redis
		redis_client.GetIMRedisClientInstance().DelUserID2instanceIP(fromID)
	}()
	msg := &message.CoreMessage{}

	for {
		if err := c.WSConn.ReadJSON(msg); err != nil {
			// close this ws conn
			ierror.SendError(ierror.NewError(pconst.CoreErrorCode_Network, err), fromID)
			if websocket.CloseNormalClosure == err.(*websocket.CloseError).Code {
				return
			}
			log.Errorf("Read json err = %s", err)
			return
		}

		switch msg.Type {
		case pconst.CoreMessageType_Message:
			// publish the given core message to other users
			if err := publisher.IMPublisher.Publish(&msg.MessagePayload); err != nil {
				ierror.SendError(err, fromID)
			}
		case pconst.CoreMessageType_Close:
			return
		case pconst.CoreMessageType_Manage:
			if err := manager.IMManager.Manage(&msg.ManagePayload); err != nil {
				ierror.SendError(err, fromID)
			}
		}
		//if msg.Type != pconst.CoreMessageType_Message {
		//	ierror.SendError(ierror.NewError(pconst.CoreErrorCode_CoreMessageTypeUnsupported, errors.New("unsupported msgType = "+strconv.Itoa(int(msg.Type)))), fromID)
		//	continue
		//}

	}

}

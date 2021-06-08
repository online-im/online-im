package handler

import (
	"github.com/glory-go/glory/config"
	ghttp "github.com/glory-go/glory/http"
	"github.com/glory-go/glory/log"
	gostNet "github.com/dubbogo/gost/net"
	"github.com/goonline/online-im/internal/constant"
	"github.com/goonline/online-im/internal/instance/common"
	"github.com/goonline/online-im/internal/instance/conn_holder"
	"github.com/goonline/online-im/internal/instance/publisher"
	"github.com/goonline/online-im/internal/redis_client"
	"github.com/goonline/online-im/pkg/message"
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
		conn_holder.GetWSConnHolder().Del(fromID)
		redis_client.GetIMRedisClientInstance().DelUserID2instanceIP(fromID)
	}()
	msg := &message.Message{}
	for {
		if err := c.WSConn.ReadJSON(msg); err != nil {
			log.Errorf("Read json err = %s", err)
			return
		}
		if err := publisher.IMPublisher.Publish(&common.DataFromClient{
			FromID:      msg.FromID,
			Message:     msg.Data,
			PublishType: constant.PublishMessageType(msg.Type),
			TargetID:    msg.TargetID,
		}); err != nil {
			log.Error(err)
		}
	}

}

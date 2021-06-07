package handler

import (
	"git.go-online.org.cn/Glory/glory/config"
	ghttp "git.go-online.org.cn/Glory/glory/http"
	"git.go-online.org.cn/Glory/glory/log"
	gostNet "github.com/dubbogo/gost/net"
	"github.com/goonline/online-im/internal/constant"
	"github.com/goonline/online-im/internal/instance/common"
	"github.com/goonline/online-im/internal/instance/http/pkg"
	"github.com/goonline/online-im/internal/instance/publisher"
	"github.com/goonline/online-im/internal/redis_client"
	"strconv"
)

var localIP string

func init() {
	var err error
	localIP, err = gostNet.GetLocalIP()
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
	redis_client.GetIMRedisClientInstance().SetUserID2instanceIP(fromID, localIP+":"+strconv.Itoa(config.GlobalServerConf.ServiceConfigs[constant.GRPCProviderName].Port))
	msg := &pkg.Message{}

	for {
		if err := c.WSConn.ReadJSON(msg); err != nil {
			log.Errorf("Read json err = %s", err)
			return
		}
		if err := publisher.IMPublisher.Publish(&common.DataFromClient{
			FromID:      msg.FromID,
			Message:     msg.Message,
			PublishType: constant.PublishMessageType(msg.PublishType),
			TargetID:    msg.TargetID,
		}); err != nil {
			log.Error(err)
		}
	}

}

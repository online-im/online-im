package http

import (
	ghttp "github.com/glory-go/glory/http"
	gatewayConfig "github.com/goonline/online-im/internal/gateway/config"
	"github.com/goonline/online-im/internal/gateway/instance_selector"
	"github.com/goonline/online-im/internal/redis_client"
)

import (
	"crypto/rand" //真随机
	"math/big"
)

func GetRandom(max uint32) int64 {
	// 生成 20 个 [0, 100) 范围的真随机数。
	val, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return val.Int64()
}

func ConnHandler(c *ghttp.GRegisterController) error {
	// register id to instance
	rsp, _ := c.Rsp.(*ConnRsp)
	address := ""
	var err error
	if gatewayConfig.GlobalConfig.InstanceDiscType == "redis" {
		list, e := redis_client.GetIMRedisClientInstance().GetInstances()
		if e == nil {
			address = list[GetRandom(uint32(len(list)))] // random select instance
		}
		err = e
	} else if gatewayConfig.GlobalConfig.InstanceDiscType == "k8s" {
		address, err = instance_selector.GetInstanceSelector().Select() // round robin select instance
	}
	if err != nil {
		rsp.Ok = false
		return err
	}
	rsp.Address = address
	rsp.Ok = true
	return nil
}

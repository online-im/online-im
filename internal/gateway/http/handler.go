package http

import (
	ghttp "git.go-online.org.cn/Glory/glory/http"
	"git.go-online.org.cn/Glory/glory/log"
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
	//address, err := instance_selector.GetInstanceSelector().Select()
	list, err := redis_client.GetIMRedisClientInstance().GetInstances()
	address := list[GetRandom(3)]
	if err != nil {
		// try redis: for local debug, FIXME
		log.Errorf("gateway k8s selector select target instances ips filed with err = %s", err)
		list, err := redis_client.GetIMRedisClientInstance().GetInstances()
		if err != nil {
			log.Errorf("gateway redis selector select target instances ips failed with err = %s", err)
		}
		if len(list) == 0 {
			log.Errorf("gateway redis selector select target instances ips failed with instance ip list empty")
		}
		rsp.Ok = false
		return err
	}
	rsp.Address = address
	rsp.Ok = true
	return nil
}

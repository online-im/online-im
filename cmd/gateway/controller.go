package gateway

import (
	"git.go-online.org.cn/Glory/glory/common"
	"git.go-online.org.cn/Glory/glory/config"
	"git.go-online.org.cn/Glory/glory/plugin"
	_ "git.go-online.org.cn/Glory/glory/registry/k8s"
	"github.com/goonline/online-im/internal/constant"
	"github.com/goonline/online-im/internal/redis_client"
)

func main() {
	reg := plugin.GetRegistry(config.GlobalServerConf.RegistryConfig[constant.RegistryKey])
	eventCh, err := reg.Subscribe(constant.GRPCProviderName)
	if err != nil {
		panic(err)
	}

	imRedisClient := redis_client.GetIMRedisClientInstance()
	for v := range eventCh {
		addr := v.Addr.GetUrl()
		switch v.Opt {
		case common.RegistryAddEvent, common.RegistryUpdateEvent:
			imRedisClient.ADDInstance(addr)
		case common.RegistryDeleteEvent:

		}
	}

}

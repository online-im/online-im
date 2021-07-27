package main

import (
	"flag"
	gConfig "github.com/glory-go/glory/config"
	"github.com/glory-go/glory/glory"
	"github.com/glory-go/glory/log"
	"github.com/online-im/online-im/internal/constant"
	"github.com/online-im/online-im/internal/instance/config"
	"github.com/online-im/online-im/internal/instance/http"
	"github.com/online-im/online-im/internal/instance/http/handler"
	"github.com/online-im/online-im/internal/instance/manager"
	"github.com/online-im/online-im/internal/instance/publisher"
	"github.com/online-im/online-im/internal/instance/service"
	"github.com/online-im/online-im/internal/redis_client"
	"os"
	"strconv"

	_ "github.com/glory-go/glory/registry/k8s"
)

func main() {
	mode := "dev"
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}
	if !(mode == "dev" || mode == "pro") {
		log.Errorf("mod should be dev or pro, given is mode = %s", mode)
		return
	}

	imConfig := &config.Config{
		Mod: mode,
	}

	loadInstanceConfig(imConfig)

	if mode == "dev" {
		setDevGloryConfig(imConfig)
	} else {
		setProGloryConfig(imConfig)
	}

	if err := publisher.NewPublisherInstance(imConfig); err != nil {
		panic(err)
	}

	if err := manager.NewManagerInstance(imConfig); err != nil {
		panic(err)
	}

	gloryServer := glory.NewServer()

	// start httpHandler
	imHttpService := http.NewIMHTTPService(imConfig)
	imHttpService.Start(gloryServer)

	// start IMService
	imService := service.NewIMGRPCService(imConfig)
	imService.Start(gloryServer)

	if err := redis_client.GetIMRedisClientInstance().ADDInstance(handler.LocalIP + ":" + strconv.Itoa(gConfig.GlobalServerConf.ServiceConfigs[constant.HTTPProviderName].Port)); err != nil {
		panic(err)
	}
	gloryServer.Run()
}

func setDevGloryConfig(conf *config.Config) {
	gConfig.GlobalServerConf.OrgName = "glory"
	gConfig.GlobalServerConf.ServerName = "go-online"
	gConfig.GlobalServerConf.RedisConfig[constant.RedisKey] = &gConfig.RedisConfig{
		Host: conf.RedisHost,
		Port: conf.RedisPort,
	}

	gConfig.GlobalServerConf.RegistryConfig[constant.RegistryKey] = &gConfig.RegistryConfig{
		Service: "k8s",
	}

	gConfig.GlobalServerConf.ServiceConfigs[constant.GRPCProviderName] = &gConfig.ServiceConfig{
		Protocol: "grpc",
		Port:     8081,
	}
	gConfig.GlobalServerConf.ServiceConfigs[constant.HTTPProviderName] = &gConfig.ServiceConfig{
		Protocol: "http",
		Port:     8080,
	}
	gConfig.GlobalServerConf.ClientConfig[constant.GRPCClientName] = &gConfig.ClientConfig{
		Protocol:  "grpc",
		ServiceID: constant.IMInstanceServiceID,
	}
}

func setProGloryConfig(conf *config.Config) {
	gConfig.GlobalServerConf.OrgName = "glory"
	gConfig.GlobalServerConf.ServerName = "go-online"
	gConfig.GlobalServerConf.RedisConfig[constant.RedisKey] = &gConfig.RedisConfig{
		Host: "online-im-redis",
		Port: "6379",
	}

	gConfig.GlobalServerConf.RegistryConfig[constant.RegistryKey] = &gConfig.RegistryConfig{
		Service: "k8s",
	}

	gConfig.GlobalServerConf.ServiceConfigs[constant.GRPCProviderName] = &gConfig.ServiceConfig{
		Protocol: "grpc",
		Port:     8081,
	}
	gConfig.GlobalServerConf.ServiceConfigs[constant.HTTPProviderName] = &gConfig.ServiceConfig{
		Protocol: "http",
		Port:     8080,
	}
	gConfig.GlobalServerConf.ClientConfig[constant.GRPCClientName] = &gConfig.ClientConfig{
		Protocol:  "grpc",
		ServiceID: constant.IMInstanceServiceID,
	}
}

func loadInstanceConfig(conf *config.Config) {
	conf.RedisHost = redisHost
	conf.RedisPort = redisPort
}

var (
	redisHost string
	redisPort string
)

func init() {
	flag.StringVar(&redisHost, "rh", "localhost", "redis host, used by dev mod")
	flag.StringVar(&redisPort, "rp", "6379", "redis host, used by dev mod")
}

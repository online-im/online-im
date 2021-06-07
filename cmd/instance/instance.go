package main

import (
	"git.go-online.org.cn/Glory/glory/glory"
	"github.com/goonline/online-im/internal/instance/config"
	"github.com/goonline/online-im/internal/instance/http"
	"github.com/goonline/online-im/internal/instance/publisher"
	"github.com/goonline/online-im/internal/instance/service"

	_ "git.go-online.org.cn/Glory/glory/registry/k8s"
)

func main() {

	imConfig := &config.Config{}
	if err := publisher.NewPublisherInstance(imConfig); err != nil {
		panic(err)
	}

	gloryServer := glory.NewServer()

	// start httpHandler
	imHttpService := http.NewIMHTTPService(imConfig)
	imHttpService.Start(gloryServer)

	// start IMService
	imService := service.NewIMGRPCService(imConfig)
	imService.Start(gloryServer)

	gloryServer.Run()
}

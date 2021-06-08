package main

import (
	"github.com/glory-go/glory/glory"
	_ "github.com/glory-go/glory/registry/k8s"
	"github.com/goonline/online-im/internal/gateway/config"
	"github.com/goonline/online-im/internal/gateway/http"
)

func main() {

	config.GlobalConfig.InstanceDiscType = "k8s"

	gloryServer := glory.NewServer()

	// start httpHandler
	imHttpService := http.NewIMGatewayHTTPService()
	imHttpService.Start(gloryServer)

	gloryServer.Run()
}

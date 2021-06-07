package main

import (
	"git.go-online.org.cn/Glory/glory/glory"
	_ "git.go-online.org.cn/Glory/glory/registry/k8s"
	"github.com/goonline/online-im/internal/gateway/http"
)

func main() {
	gloryServer := glory.NewServer()

	// start httpHandler
	imHttpService := http.NewIMGatewayHTTPService()
	imHttpService.Start(gloryServer)

	gloryServer.Run()
}

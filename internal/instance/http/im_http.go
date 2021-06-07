package http

import (
	"git.go-online.org.cn/Glory/glory/server"
	"git.go-online.org.cn/Glory/glory/service"
	"github.com/goonline/online-im/internal/constant"
	"github.com/goonline/online-im/internal/instance/config"
	"github.com/goonline/online-im/internal/instance/http/handler"
)

type IMHTTPService struct {
}

func NewIMHTTPService(conf *config.Config) *IMHTTPService {
	return &IMHTTPService{}
}

func (s *IMHTTPService) Start(gloryServer server.GloryServer) {
	httpService := service.NewHttpService(constant.HTTPProviderName)
	httpService.RegisterWSRouter("/online", handler.OnlineHandler)
	gloryServer.RegisterService(httpService)
}

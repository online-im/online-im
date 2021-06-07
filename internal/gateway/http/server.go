package http

import (
	"git.go-online.org.cn/Glory/glory/server"
	"git.go-online.org.cn/Glory/glory/service"
	"github.com/goonline/online-im/internal/constant"
)

type IMGatewayHTTPService struct {
}

func NewIMGatewayHTTPService() *IMGatewayHTTPService {
	return &IMGatewayHTTPService{}
}

func (s *IMGatewayHTTPService) Start(gloryServer server.GloryServer) {
	httpService := service.NewHttpService(constant.HTTPGatewayProviderName)
	httpService.RegisterRouter("/connect", ConnHandler, nil, &ConnRsp{}, "GET")
	gloryServer.RegisterService(httpService)
}

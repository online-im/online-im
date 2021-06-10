package http

import (
	"github.com/glory-go/glory/server"
	"github.com/glory-go/glory/service"
	"github.com/online-im/online-im/internal/constant"
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

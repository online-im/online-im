package service

import (
	"git.go-online.org.cn/Glory/glory/service"
	"github.com/goonline/online-im/internal/constant"
	pb "github.com/goonline/online-im/internal/instance/api"
	"github.com/goonline/online-im/internal/instance/config"

	"git.go-online.org.cn/Glory/glory/server"
)

type IMGRPCService struct {
}

func NewIMGRPCService(config *config.Config) *IMGRPCService {
	return &IMGRPCService{}
}

func (s IMGRPCService) Start(server server.GloryServer) {
	gloryService := service.NewGrpcService(constant.GRPCProviderName)
	pb.RegisterIMServiceProviderServer(gloryService.GetGrpcServer(), &provider{})
	server.RegisterService(gloryService)
}

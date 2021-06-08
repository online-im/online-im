package publisher

import (
	ggrpc "github.com/glory-go/glory/grpc"
	"github.com/goonline/online-im/internal/constant"
	pb "github.com/goonline/online-im/internal/instance/api"
)

type IMGRPCClientCache struct {
	grpcClientMap map[string]pb.IMServiceProviderClient
}

func NewIMGRPCClientCache() *IMGRPCClientCache {
	return &IMGRPCClientCache{}
}

func (c *IMGRPCClientCache) GetIMGRPCClient(instanceIP string) (pb.IMServiceProviderClient, error) {
	client, ok := c.grpcClientMap[instanceIP]
	if ok {
		return client, nil
	}
	rawClient := ggrpc.NewGrpcClientWithDynamicAddr(constant.GRPCClientName, instanceIP)
	return pb.NewIMServiceProviderClient(rawClient.GetConn()), nil
}

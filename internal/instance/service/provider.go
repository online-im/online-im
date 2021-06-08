package service

import (
	"context"
	"github.com/glory-go/glory/log"
	_ "github.com/glory-go/glory/registry/nacos"
	pb "github.com/goonline/online-im/internal/instance/api"
	"github.com/goonline/online-im/internal/instance/conn_holder"
	"github.com/goonline/online-im/pkg/message"
)

// server is used to implement helloworld.GreeterServer.
type provider struct {
}

func (s *provider) PublishMessage(ctx context.Context, in *pb.PublishMessageRequest) (*pb.PublishMessageResponse, error) {
	log.Infof("Received: %v", in.GetData())
	if err := conn_holder.GetWSConnHolder().Send(in.ToID, message.Message{
		Type:     0, // todo
		FromID:   in.FromID,
		TargetID: in.ToID,
		Data:     in.Data,
	}); err != nil {
		log.Errorf("send message %s to id %d err = %s", in.Data, in.ToID, err)
	}
	return &pb.PublishMessageResponse{Message: "success", Code: 0}, nil
}

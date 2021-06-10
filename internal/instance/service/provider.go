package service

import (
	"context"
	"fmt"
	"github.com/glory-go/glory/log"
	_ "github.com/glory-go/glory/registry/nacos"
	pb "github.com/online-im/online-im/internal/instance/api"
	"github.com/online-im/online-im/internal/instance/conn_holder"
	"github.com/online-im/online-im/pkg/constant"
)

// server is used to implement helloworld.GreeterServer.
type provider struct {
}

func (s *provider) PublishMessage(ctx context.Context, in *pb.PublishMessageRequest) (*pb.PublishMessageResponse, error) {
	log.Infof("Received: %v", in.GetData())
	if err := conn_holder.GetWSConnHolder().SendMessageWithData(in.FromID, in.ToID, []byte(in.Data)); err != nil {
		log.Errorf("send message %s to id %s err = %s", in.Data, in.ToID, err)
		code := constant.CoreErrorCode_UserOffLine
		//if ie, ok := err.(*ierror.Error); ok {
		//	code = ie.Code
		//}
		return &pb.PublishMessageResponse{Message: fmt.Sprintf("send message %s to id %s err = %s", in.Data, in.ToID, err), Code: uint32(code)}, nil
	}
	return &pb.PublishMessageResponse{Message: "success", Code: 0}, nil
}

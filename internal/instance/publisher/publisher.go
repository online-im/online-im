package publisher

import (
	"context"
	"github.com/glory-go/glory/log"
	pb "github.com/online-im/online-im/internal/instance/api"
	"github.com/online-im/online-im/internal/instance/config"
	"github.com/online-im/online-im/internal/instance/ierror"
	"github.com/online-im/online-im/pkg/constant"
	"github.com/online-im/online-im/pkg/message"
	perrors "github.com/pkg/errors"
)

var IMPublisher *Publisher

type Publisher struct {
	cache      *Cache
	grpcClient *IMGRPCClientCache
}

func NewPublisherInstance(conf *config.Config) error {
	cache, err := NewCache(conf)
	if err != nil {
		return err
	}
	IMPublisher = &Publisher{
		cache:      cache,
		grpcClient: NewIMGRPCClientCache(),
	}
	return nil
}

func (p *Publisher) Publish(msg *message.CoreMessagePayload) error {
	var err error
	switch msg.PublishType {
	case constant.Publish2User:
		err = p.publish2User(msg)
	case constant.Publish2Group:
		err = p.publish2Group(msg)
		//todo extension
	}
	return err
}

func (p *Publisher) publish2User(msg *message.CoreMessagePayload) error {
	instanceIP, err := p.cache.GetUserID2InstanceIP(msg.TargetID)
	if err != nil {
		return err
	}
	client := p.grpcClient.GetIMGRPCClient(instanceIP)

	rsp, err := client.PublishMessage(context.Background(), &pb.PublishMessageRequest{
		FromID: msg.FromID,
		ToID:   msg.TargetID,
		Data:   string(msg.Data),
	})
	if err != nil {
		log.Errorf("GetUserID2InstanceIP error = %s", err)
		return ierror.NewError(constant.CoreErrorCode_PublishError, perrors.Errorf("Publish message error = %s", err))
	}
	if rsp.Code != 0 {
		return ierror.NewError(constant.CoreErrorCode(rsp.Code), perrors.Errorf(rsp.Message))
	}
	return nil
}

func (p *Publisher) publish2Group(msg *message.CoreMessagePayload) error {
	userIDs, err := p.cache.GetRoomID2UserIDs(msg.TargetID)
	if err != nil {
		return err
	}

	for _, v := range userIDs {
		go func(v string, msg message.CoreMessagePayload) {
			if err := p.publish2User(&message.CoreMessagePayload{
				Data:        msg.Data,
				FromID:      msg.FromID,
				TargetID:    v,
				PublishType: msg.PublishType,
			}); err != nil {
				log.Errorf("error publish to group err = %s", err)
			}
		}(v, *msg)
	}
	return nil
}

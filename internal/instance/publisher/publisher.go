package publisher

import (
	"context"
	"github.com/goonline/online-im/internal/constant"
	pb "github.com/goonline/online-im/internal/instance/api"
	"github.com/goonline/online-im/internal/instance/common"
	"github.com/goonline/online-im/internal/instance/config"
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

func (p *Publisher) Publish(msg *common.DataFromClient) error {
	var err error
	switch msg.PublishType {
	case constant.Publish2User:
		err = p.publish2User(msg)
	case constant.Publish2Room:
		err = p.publish2Room(msg)
	case constant.Publish2Global:
		//todo
	}
	return err
}

func (p *Publisher) publish2User(msg *common.DataFromClient) error {
	instanceIP, err := p.cache.GetUserID2InstanceIP(msg.TargetID)
	if err != nil {
		return err
	}
	client, err := p.grpcClient.GetIMGRPCClient(instanceIP)
	if err != nil {
		return err
	}
	rsp, err := client.PublishMessage(context.Background(), &pb.PublishMessageRequest{
		FromID: msg.FromID,
		ToID:   msg.TargetID,
		Data:   msg.Message,
	})
	if err != nil {
		return perrors.Errorf("Publish message error = %s", err)
	}
	if rsp.Code != 0 {
		return perrors.Errorf("Publish message rsp code = %d, msg = %s", err, rsp.Code, rsp.Message)
	}
	return nil
}

func (p *Publisher) publish2Room(msg *common.DataFromClient) error {
	userIDs, err := p.cache.GetRoomID2UserIDs(msg.TargetID)
	if err != nil {
		return err
	}

	for _, v := range userIDs {
		if err := p.publish2User(&common.DataFromClient{
			FromID:      msg.FromID,
			TargetID:    v,
			PublishType: constant.Publish2Room,
			Message:     msg.Message,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (p *Publisher) publish2Global(msg *common.DataFromClient) error {
	// todo
	return nil
}

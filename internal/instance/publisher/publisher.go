package publisher

import (
	"context"
	"strconv"

	"github.com/glory-go/glory/log"
	pb "github.com/online-im/online-im/internal/instance/api"
	"github.com/online-im/online-im/internal/instance/config"
	"github.com/online-im/online-im/internal/instance/ierror"
	"github.com/online-im/online-im/internal/storage/dal"
	"github.com/online-im/online-im/internal/storage/service"
	"github.com/online-im/online-im/pkg/constant"
	"github.com/online-im/online-im/pkg/message"
	perrors "github.com/pkg/errors"
)

var IMPublisher *Publisher

type Publisher struct {
	cache          *Cache
	grpcClient     *IMGRPCClientCache
	storageService service.Service
}

func NewPublisherInstance(conf *config.Config) error {
	cache, err := NewCache(conf)
	if err != nil {
		return err
	}
	dalClient := dal.NewMysqlDal()
	IMPublisher = &Publisher{
		cache:          cache,
		grpcClient:     NewIMGRPCClientCache(),
		storageService: service.NewService(dalClient),
	}
	return nil
}

func (p *Publisher) Publish(ctx context.Context, msg *message.CoreMessagePayload) error {
	// save message before send
	fromID, err := strconv.ParseInt(msg.FromID, 10, 64)
	if err != nil {
		return constant.CoreErrorCode_ParamError
	}
	targetID, err := strconv.ParseInt(msg.TargetID, 10, 64)
	if err != nil {
		return constant.CoreErrorCode_ParamError
	}
	_, err = p.storageService.SendMsgTo(ctx, fromID, targetID, string(msg.Data), msg.PublishType)
	if err != nil {
		return err
	}
	// publish message to user
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

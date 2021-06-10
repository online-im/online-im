package manager

import (
	"github.com/online-im/online-im/internal/instance/config"
	"github.com/online-im/online-im/internal/redis_client"
	"github.com/online-im/online-im/pkg/constant"
	"github.com/online-im/online-im/pkg/message"
)

var IMManager *Manager

type Manager struct {
}

func NewManagerInstance(conf *config.Config) error {
	IMManager = &Manager{}
	return nil
}

func (p *Manager) Manage(msg *message.ManagePayload) error {
	switch msg.ManageType {
	case constant.ManageType_CreateGroup:
		return redis_client.GetIMRedisClientInstance().CreateGroup(msg.User, msg.GroupID, msg.Users)
	case constant.ManageType_DeleteGroup:
		return redis_client.GetIMRedisClientInstance().DeleteGroup(msg.GroupID)
	case constant.ManageType_AddUserToGroup:
		return redis_client.GetIMRedisClientInstance().ADDUserID2Group(msg.User, msg.GroupID)
	case constant.ManageType_RemoveUserFromGroup:
		return redis_client.GetIMRedisClientInstance().RemoveUserIDFromGroup(msg.User, msg.GroupID)
	}
	return nil
}

//func (p *Publisher) publish2User(msg *message.CoreMessagePayload) error {
//	instanceIP, err := p.cache.GetUserID2InstanceIP(msg.TargetID)
//	if err != nil {
//		return err
//	}
//	client := p.grpcClient.GetIMGRPCClient(instanceIP)
//
//	rsp, err := client.PublishMessage(context.Background(), &pb.PublishMessageRequest{
//		FromID: msg.FromID,
//		ToID:   msg.TargetID,
//		Data:   string(msg.Data),
//	})
//	if err != nil {
//		log.Errorf("GetUserID2InstanceIP error = %s", err)
//		return ierror.NewError(constant.CoreErrorCode_PublishError, perrors.Errorf("Publish message error = %s", err))
//	}
//	if rsp.Code != 0 {
//		return ierror.NewError(constant.CoreErrorCode(rsp.Code), perrors.Errorf(rsp.Message))
//	}
//	return nil
//}
//
//func (p *Publisher) publish2Room(msg *message.CoreMessagePayload) error {
//	//userIDs, err := p.cache.GetRoomID2UserIDs(msg.TargetID)
//	//if err != nil {
//	//	return err
//	//}
//
//	//for _, v := range userIDs {
//	//	if err := p.publish2User(msg); err != nil {
//	//		return err
//	//	}
//	//}
//
//	return nil
//}
//
//func (p *Publisher) publish2Global(msg *message.CoreMessagePayload) error {
//	// todo
//	return nil
//}

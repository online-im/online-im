package redis_client

import (
	"github.com/glory-go/glory/log"
	gloryRedis "github.com/glory-go/glory/redis"
	"github.com/go-redis/redis"
	"github.com/online-im/online-im/internal/constant"
	"strings"
	"sync"
)

type IMRedisClient struct {
	rawClient *redis.Client
}

var imRedisClientInstance *IMRedisClient

var once sync.Once

func GetIMRedisClientInstance() *IMRedisClient {
	once.Do(func() {
		redisClient, err := gloryRedis.NewRedisClient(constant.RedisKey, 0)
		if err != nil {
			log.Errorf("GetIMRedisClientInstance err = %s", err)
		}
		imRedisClientInstance = &IMRedisClient{
			rawClient: redisClient,
		}
	})
	return imRedisClientInstance
}

func (imrc *IMRedisClient) GetUserID2instanceIP(uid string) (string, error) {
	return imrc.rawClient.Get(uid).Result()
}

func (imrc *IMRedisClient) GetRoomID2userIDs(rid string) ([]string, error) {
	res, err := imrc.rawClient.Exists(rid).Result()
	if err != nil {
		return nil, err
	}
	if res == 0 {
		return make([]string, 0), nil
	}
	val, err := imrc.rawClient.Get(rid).Result()
	if err != nil {
		return nil, err
	}
	return strings.Split(val, "|"), nil
}

func (imrc *IMRedisClient) SetUserID2instanceIP(uid, instanceIP string) {
	imrc.rawClient.Set(uid, instanceIP, 0)
}

func (imrc *IMRedisClient) DelUserID2instanceIP(uid string) {
	imrc.rawClient.Del(uid)
}

func (imrc *IMRedisClient) CreateGroup(uid, groupID string, userIDList []string) error {
	newUserList := append(userIDList, uid)
	if err := imrc.rawClient.Set(groupID, strings.Join(newUserList, "|"), 0).Err(); err != nil {
		log.Error("CreateGroup Set failed with err = ", err)
		return err
	}
	return nil
}

func (imrc *IMRedisClient) DeleteGroup(groupID string) error {
	err, ok := imrc.exist(groupID)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	if err := imrc.rawClient.Del(groupID).Err(); err != nil {
		log.Error("ADDUserID2RoomID Del roomid failed with err = ", err)
		return err
	}
	return nil
}

func (imrc *IMRedisClient) ADDUserID2Group(uid, groupID string) error {
	err, ok := imrc.exist(groupID)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	userList, err := imrc.GetRoomID2userIDs(groupID)
	if err != nil {
		log.Error("ADDUserID2RoomID failed with err = ", err)
		return err
	}
	for _, v := range userList {
		if v == uid {
			return nil
		}
	}
	userList = append(userList, uid)
	imrc.rawClient.Set(groupID, strings.Join(userList, "|"), 0)
	return nil
}

func (imrc *IMRedisClient) exist(key string) (error, bool) {
	res, err := imrc.rawClient.Exists(key).Result()
	if err != nil {
		return err, false
	}
	if res == 0 {
		return nil, false
	}
	return nil, true
}

func (imrc *IMRedisClient) RemoveUserIDFromGroup(uid, groupID string) error {
	err, ok := imrc.exist(groupID)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	userList, err := imrc.GetRoomID2userIDs(groupID)
	if err != nil {
		log.Error("ADDUserID2RoomID failed with err = ", err)
		return err
	}
	newUserList := make([]string, 0)
	for _, v := range userList {
		if v != uid {
			newUserList = append(newUserList, v)
		}
	}
	if len(newUserList) == 0 {
		return imrc.DeleteGroup(groupID)
	}
	if err := imrc.rawClient.Set(groupID, strings.Join(newUserList, "|"), 0).Err(); err != nil {
		log.Error("ADDUserID2RoomID Set failed with err = ", err)
		return err
	}
	return nil
}

func (imrc *IMRedisClient) ADDInstance(instanceAddr string) error {
	res, err := imrc.rawClient.Exists(constant.IMInstanceIPsKey).Result()
	if err != nil {
		return err
	}
	ipList := make([]string, 0)
	if res != 0 {
		val, err := imrc.rawClient.Get(constant.IMInstanceIPsKey).Result()
		if err != nil {
			return err
		}
		ipList = strings.Split(val, "|")
		for _, v := range ipList {
			if v == instanceAddr {
				return nil
			}
		}
	}
	ipList = append(ipList, instanceAddr)
	imrc.rawClient.Set(constant.IMInstanceIPsKey, strings.Join(ipList, "|"), 0)
	return nil
}

func (imrc *IMRedisClient) GetInstances() ([]string, error) {
	val, err := imrc.rawClient.Get(constant.IMInstanceIPsKey).Result()
	if err != nil {
		return nil, err
	}
	return strings.Split(val, "|"), nil
}

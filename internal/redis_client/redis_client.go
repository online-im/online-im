package redis_client

import (
	"git.go-online.org.cn/Glory/glory/log"
	gloryRedis "git.go-online.org.cn/Glory/glory/redis"
	"github.com/go-redis/redis"
	"github.com/goonline/online-im/internal/constant"
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

func (imrc *IMRedisClient) ADDUserID2RoomID(uid, roomID string) error {
	userList, err := imrc.GetRoomID2userIDs(roomID)
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
	imrc.rawClient.Set(roomID, strings.Join(userList, "|"), 0)
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

package publisher

import (
	"github.com/online-im/online-im/internal/instance/config"
	"github.com/online-im/online-im/internal/instance/ierror"
	"github.com/online-im/online-im/internal/redis_client"
	"github.com/online-im/online-im/pkg/constant"
)

type Cache struct {
	imRedisClient *redis_client.IMRedisClient
	uID2iIP       map[string]string
	rID2uIDs      map[string][]string
}

func NewCache(conf *config.Config) (*Cache, error) {
	return &Cache{
		imRedisClient: redis_client.GetIMRedisClientInstance(),
		uID2iIP:       make(map[string]string),
		rID2uIDs:      make(map[string][]string),
	}, nil
}

func (c *Cache) GetUserID2InstanceIP(userID string) (string, error) {
	iIP, _ := c.uID2iIP[userID]
	// todo using subscribe cache
	if true {
		iIP, err := c.imRedisClient.GetUserID2instanceIP(userID)
		if err != nil {
			return "", ierror.NewError(constant.CoreErrorCode_UserOffLine, err)
		}
		c.uID2iIP[userID] = iIP
		return iIP, nil
	}
	return iIP, nil
}

func (c *Cache) GetRoomID2UserIDs(roomID string) ([]string, error) {
	uIPs, ok := c.rID2uIDs[roomID]
	if !ok {
		uIPs, err := c.imRedisClient.GetRoomID2userIDs(roomID)
		if err != nil {
			return make([]string, 0), err
		}
		c.rID2uIDs[roomID] = uIPs
		return uIPs, nil
	}
	return uIPs, nil
}

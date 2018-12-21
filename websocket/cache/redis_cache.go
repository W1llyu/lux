package cache

import (
	"github.com/W1llyu/gdao/xredis"
)

type RedisCacheAdapter struct {}

const (
	HKEY = "LUX_REDIS_ROOMS"
)

func (c *RedisCacheAdapter) ClearAll() error {
	client := xredis.GetClient()
	defer client.Close()
	return client.Del(HKEY)
}

func (c *RedisCacheAdapter) IncreaseRoomMemberCount(room string, count int) error {
	client := xredis.GetClient()
	defer client.Close()
	val, err := client.Hincrby(HKEY, room, count)
	if err != nil {
		return err
	}

	if val <= 0 {
		return client.Hdel(HKEY, room)
	}

	return nil
}

func (c *RedisCacheAdapter) DecreaseRoomMemberCount(room string, count int) error {
	return c.IncreaseRoomMemberCount(room, -count)
}

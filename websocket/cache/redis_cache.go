package cache

import (
	"github.com/W1llyu/gdao/xredis"
	"github.com/W1llyu/lux/websocket/constant"
)

type RedisCacheAdapter struct {}

func (c *RedisCacheAdapter) ClearAll() error {
	client := xredis.GetClient()
	defer client.Close()
	return client.Del(constant.SOCKET_CACHE_KEY)
}

func (c *RedisCacheAdapter) IncreaseRoomMemberCount(room string, count int) error {
	client := xredis.GetClient()
	defer client.Close()
	val, err := client.Hincrby(constant.SOCKET_CACHE_KEY, room, count)
	if err != nil {
		return err
	}

	if val <= 0 {
		return client.Hdel(constant.SOCKET_CACHE_KEY, room)
	}

	return nil
}

func (c *RedisCacheAdapter) DecreaseRoomMemberCount(room string, count int) error {
	return c.IncreaseRoomMemberCount(room, -count)
}

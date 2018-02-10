package queue

import (
	"github.com/Lux-go/common/dao/xredis"
)

const (
	CHANNELKEY = "bet_topics"
)

type RedisQueue struct {}

func (q *RedisQueue) OnMessage(callback interface{}) {
	client := xredis.GetPubSubClient()
	defer client.Close()
	client.Subscribe(CHANNELKEY)
	client.Receive(callback)
}

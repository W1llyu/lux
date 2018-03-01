// Message Queue Implement In Redis Pub/Sub
package queue

import (
	"github.com/W1llyu/gdao/xredis"
)

type RedisQueue struct{}

func (q *RedisQueue) OnMessage(callback interface{}) {
	client := xredis.GetPubSubClient()
	defer client.Close()
	client.Subscribe(QUEUEKEY)
	client.Receive(callback)
}

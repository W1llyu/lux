/*
 * Redis Sub/Pub 消息队列
 */
package queue

import (
	"github.com/W1llyu/gdao/xredis"
)

type RedisQueue struct{
	Name string
}

func (q *RedisQueue) OnMessage(callback interface{}) {
	client := xredis.GetPubSubClient()
	defer client.Close()
	client.Subscribe(q.Name)
	client.Receive(callback)
}

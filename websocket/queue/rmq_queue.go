/*
 * RabbitMQ Consumer 消息队列
 */
package queue

import (
	"github.com/W1llyu/gdao/xrmq"
	"github.com/W1llyu/lux/websocket/constant"
)

type RmqQueue struct{
	Name string
}

func (q *RmqQueue) OnMessage(callback interface{}) {
	client := xrmq.GetClient()
	defer client.Close()
	ctx := xrmq.ExchangeCtx{
		Name:       constant.DEFAULT_KEY,
		Type:       "topic",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}
	client.Receive(ctx, q.Name, callback)
}

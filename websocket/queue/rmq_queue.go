/*
 * RabbitMQ Consumer 消息队列
 */
package queue

import "github.com/W1llyu/gdao/xrmq"

type RmqQueue struct{
	Name string
}

func (q *RmqQueue) OnMessage(callback interface{}) {
	client := xrmq.GetClient()
	defer client.Close()
	ctx := xrmq.ExchangeCtx{
		Name:       DEFAULT_KEY,
		Type:       "topic",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}
	ch := client.GetChannel(ctx)
	ch.Receive(q.Name, callback)
}

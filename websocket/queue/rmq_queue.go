// Message Queue Implement In RabbitMQ
package queue

import "github.com/W1llyu/gdao/xrmq"

type RmqQueue struct{}

func (q *RmqQueue) OnMessage(callback interface{}) {
	client := xrmq.GetClient()
	defer client.Close()
	ctx := xrmq.ExchangeCtx{
		Name:       "logs_direct",
		Type:       "direct",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}
	ch := client.GetChannel(ctx)
	ch.Receive(QUEUEKEY, callback)
}

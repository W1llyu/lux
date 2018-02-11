package queue

import "github.com/Lux-go/common/dao/xamqp"

type RmqQueue struct{}

func (q *RmqQueue) OnMessage(callback interface{}) {
	client := xamqp.GetClient()
	ctx := xamqp.ExchangeCtx{
		Name:       "logs_direct",
		Type:       "direct",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}
	ch := client.GetChannel(ctx)
	ch.Receive("", callback)
}

package xamqp

import "github.com/streadway/amqp"

type Channel struct {
	channel *amqp.Channel
	ExchangeCtx ExchangeCtx
}

type ExchangeCtx struct {
	Name string
	Type string
	Durable bool
	AutoDelete bool
	Internal bool
	NoWait bool
	Args amqp.Table
}

func NewDefaultExchangeCtx() ExchangeCtx {
	return ExchangeCtx{
		Name: "",
		Type: "fanout",
		Durable: true,
		AutoDelete: false,
		Internal: false,
		NoWait: false,
		Args: nil,
	}
}

func (ch *Channel) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	return ch.channel.Publish(exchange, key, mandatory, immediate, msg)
}
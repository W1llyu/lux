package scheduler

import (
	"encoding/json"
	"github.com/Lux-go/common/utils"
	"github.com/Lux-go/websocket/runtime"
)

type QueueScheduler struct {
	queue MsgQueue
}

type MsgQueue interface {
	OnMessage(callback interface{})
}

type Msg struct {
	Channel string
	Data    interface{}
}

func CreateQueueScheduler(queue MsgQueue) *QueueScheduler {
	return &QueueScheduler{queue}
}

func (qs QueueScheduler) Run() {
	qs.queue.OnMessage(onEvent)
}

func onEvent(channel, data string) {
	utils.Infof("Redis Queue Receive %s from %s", data, channel)
	msg := &Message{
		Event: "message",
	}
	err := json.Unmarshal([]byte(data), msg)
	if err == nil {
		runtime.Server().BroadcastTo(msg.Channel, msg.Event, msg.Data)
	}
}

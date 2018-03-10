/*
 * Queue Scheduler
 * 消息队列调度模块
 */
package scheduler

import (
	"encoding/json"
	"github.com/Lux-go/utils"
	"github.com/Lux-go/websocket/runtime"
)

type QueueScheduler struct {
	queue MsgQueue
}

// 消息队列
type MsgQueue interface {
	OnMessage(callback interface{})
}

func CreateQueueScheduler(queue MsgQueue) *QueueScheduler {
	return &QueueScheduler{queue}
}

func (qs QueueScheduler) Run() {
	qs.queue.OnMessage(onEvent)
}

// 获取消息队列消息时的回调
func onEvent(channel, data string) {
	utils.Infof("Queue Receive %s from %s", data, channel)
	msg := &Message{
		Event: "message",
	}
	err := json.Unmarshal([]byte(data), msg)
	if err == nil {
		runtime.Server().BroadcastTo(msg.Channel, msg.Event, msg.Data)
	}
}
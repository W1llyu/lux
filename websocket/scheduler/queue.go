/*
 * Queue Scheduler
 * 消息队列调度模块
 */
package scheduler

import (
	"encoding/json"
	"github.com/irelia_socket/websocket/utils"
	"github.com/irelia_socket/websocket/core"
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
		core.Server().BroadcastTo(msg.Channel, msg.Event, NewMessageBody(msg.Channel, msg.Data))
	}
}

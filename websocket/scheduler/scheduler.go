/*
 * websocket的主动推送调度模块
 */
package scheduler

import (
	"time"
	"github.com/W1llyu/lux/websocket/queue"
	"github.com/W1llyu/lux/websocket/constant"
)

// 调度器接口
type Scheduler interface {
	Run()
}

type Schedulers []Scheduler

/*
 调度消息结构体
 {
 	"channel": "",
    "event": "",
    "data": ""
 }
*/
type Message struct {
	Channel string      `json:"channel"`
	Event   string      `json:"event"`
	Data    interface{} `json:"data"`
}

/*
 发送消息结构体
 Client Received:
 {
 	"channel": "",
    "message_at": 1521020318769,
    "data": ""
 }
 */
type MessageBody struct {
	Channel string      `json:"channel"`
	Data    interface{} `json:"data"`
	MessageAt  int64   `json:"message_at"`
}

func NewMessageBody(channel string, data interface {}) MessageBody {
	return MessageBody{
		Channel: channel,
		Data: data,
		MessageAt: time.Now().UnixNano() / 1000000,
	}
}

func GetSchedulers() Schedulers {
	return []Scheduler{
		CreateQueueScheduler(&queue.RedisQueue{Name: constant.DEFAULT_KEY}),
		CreateQueueScheduler(&queue.RedisQueue{Name: constant.BET_TOPIC_KEY}),
		CreateQueueScheduler(&queue.RedisQueue{Name: constant.WAGER_BET_TOPIC_KEY}),
		CreateQueueScheduler(&queue.RmqQueue{Name: constant.RMQ_ROUTER}),
	}
}

func (s Schedulers) Run() {
	for _, scheduler := range s {
		go scheduler.Run()
	}
}

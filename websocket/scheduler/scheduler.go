/*
 * websocket的主动推送调度模块
 */
package scheduler

import "github.com/Lux-go/websocket/queue"

// 调度器接口
type Scheduler interface {
	Run()
}

type Schedulers []Scheduler

// 消息结构体
type Message struct {
	Channel string      `json:"channel"`
	Event   string      `json:"event"`
	Data    interface{} `json:"data"`
}

func GetSchedulers() Schedulers {
	return []Scheduler{
		CreateQueueScheduler(new(queue.RedisQueue)),
		//CreateQueueScheduler(new(queue.RmqQueue)),
	}
}

func (self Schedulers) Run() {
	for _, s := range self {
		go s.Run()
	}
}

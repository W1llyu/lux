package scheduler

import "github.com/Lux-go/websocket/queue"

type Scheduler interface {
	Run()
}

type Message struct {
	Channel string      `json:"channel"`
	Event   string      `json:"event"`
	Data    interface{} `json:"data"`
}

func GetSchedulers() []Scheduler {
	return []Scheduler{
		CreateQueueScheduler(new(queue.RedisQueue)),
	}
}

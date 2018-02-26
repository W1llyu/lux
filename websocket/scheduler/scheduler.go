package scheduler

import "github.com/Lux-go/websocket/queue"

type Scheduler interface {
	Run()
}

type Schedulers []Scheduler

type Message struct {
	Channel string      `json:"channel"`
	Event   string      `json:"event"`
	Data    interface{} `json:"data"`
}

func GetSchedulers() Schedulers {
	return []Scheduler{
		CreateQueueScheduler(new(queue.RedisQueue)),
		CreateQueueScheduler(new(queue.RmqQueue)),
	}
}

func (self Schedulers) Run() {
	for _, s := range self {
		go s.Run()
	}
}

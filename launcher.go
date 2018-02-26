package main

import (
	"github.com/Lux-go/websocket"
	"github.com/Lux-go/websocket/scheduler"
)

var (
	forever = make(chan bool)
)

func main() {
	scheduler.GetSchedulers().Run()
	go websocket.Run()
	<-forever
}

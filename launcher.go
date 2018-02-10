package main

import (
	"github.com/Lux-go/websocket"
	"github.com/Lux-go/websocket/scheduler"
)

var (
	forever = make(chan bool)
)

func main() {
	for _, s := range scheduler.GetSchedulers() {
		go s.Run()
	}
	go websocket.Run()
	<-forever
}
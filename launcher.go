/*
 * Lux 启动入口
 */
package main

import (
	"github.com/Lux-go/websocket"
	"github.com/Lux-go/websocket/scheduler"
	"github.com/Lux-go/httpserver"
)

var (
	forever = make(chan bool)
)

func main() {
	scheduler.GetSchedulers().Run()
	go websocket.Run()
	go httpserver.Run()
	<-forever
}

/*
 * Lux 启动入口
 */
package main

import (
	"github.com/irelia_socket/websocket"
	"github.com/irelia_socket/websocket/scheduler"
	"github.com/irelia_socket/httpserver"
	"github.com/W1llyu/gdao/config"
)

var (
	forever = make(chan bool)
)

func main() {
	initializeConfigPath()
	scheduler.GetSchedulers().Run()
	go websocket.Run()
	go httpserver.Run()
	<-forever
}

func initializeConfigPath () {
	config.SetConfPath("./config/gdao.toml")
}

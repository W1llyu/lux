package core

import (
	"log"
	"sync"
	"time"
	"github.com/W1llyu/go-socket.io"
	"github.com/W1llyu/lux/websocket/cache"
	"github.com/W1llyu/lux/config"
)

var (
	once   sync.Once
	cacheAdapter cache.CacheAdapter
	Sessions = newServerSessions()
	server *socketio.Server
)

func Server() *socketio.Server {
	once.Do(initServer)
	return server
}

/**
 * Initialize Server with events binding
 */
func initServer() {
	var err error
	cacheAdapter = &cache.RedisCacheAdapter{}
	server, err = socketio.NewServer(NewLuxBroadcast(cacheAdapter), nil)
	server.SetMaxConnection(65535)
    server.SetPingTimeout(30 * time.Second)
    server.SetPingInterval(10 * time.Second)
    server.SetSessionManager(Sessions)
	if config.GetConf().Websocket.Auth {
		server.SetAllowRequest(authRequest)
	}
	if err != nil {
		log.Fatal(err)
	}
	cacheAdapter.ClearAll()
}

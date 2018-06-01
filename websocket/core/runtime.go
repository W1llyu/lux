package core

import (
	"log"
	"sync"
	"github.com/W1llyu/go-socket.io"
	"github.com/irelia_socket/websocket/cache"
	"time"
)

var (
	once   sync.Once
	cacheAdapter cache.CacheAdapter
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
    server.SetSessionManager(newServerSessions())
	server.SetAllowRequest(authRequest)
	server.SetNewId(newSocketId)
	if err != nil {
		log.Fatal(err)
	}
	cacheAdapter.ClearAll()
}

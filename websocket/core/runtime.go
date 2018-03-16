package core

import (
	"log"
	"sync"
	"github.com/W1llyu/go-socket.io"
	"github.com/Lux-go/websocket/cache"
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
	if err != nil {
		log.Fatal(err)
	}
	cacheAdapter.ClearAll()
}

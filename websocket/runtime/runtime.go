package runtime

import (
	"log"
	"sync"
	"github.com/googollee/go-socket.io"
)

var (
	once   sync.Once
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
	server, err = socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
}

package websocket

import (
	"fmt"
	"log"
	"sync"
	"net/http"
	"github.com/googollee/go-socket.io"
	"github.com/Lux-go/common/config"
	"github.com/Lux-go/common/utils"
	"github.com/Lux-go/websocket/handler"
)

var (
	server *socketio.Server
	once sync.Once
)

/**
 * Run WebSocket Server
 */
func Run() {
	once.Do(initServer)

	port := config.GetConf().Websocket.Port
	http.HandleFunc("/socket.io/", serverHandler)
	utils.Info(fmt.Sprintf("Serving at localhost:%d ...", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func GetServer() *socketio.Server {
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
	handler.BindEvents(server)
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	server.ServeHTTP(w, r)
}

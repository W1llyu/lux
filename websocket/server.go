package websocket

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"github.com/Lux-go/config"
	"github.com/Lux-go/utils"
	"github.com/Lux-go/websocket/handler"
	"github.com/Lux-go/websocket/core"
)

var (
	once sync.Once
)

/**
 * Run WebSocket Server
 */
func Run() {
	once.Do(bindEvents)

	port := config.GetConf().Websocket.Port
	http.HandleFunc("/socket.io/", serverHandler)
	utils.Info(fmt.Sprintf("Websocket Serving at localhost:%d ...", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

/**
 * Initialize Server with events binding
 */
func bindEvents() {
	server := core.Server()
	handler.BindEvents(server)
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	core.Server().ServeHTTP(w, r)
}

package httpserver

import (
	"net/http"
	"fmt"
	"github.com/irelia_socket/config"
	"github.com/irelia_socket/httpserver/router"
	"github.com/irelia_socket/websocket/utils"
)

func Run() {
	port := config.GetConf().Http.Port
	utils.Info(fmt.Sprintf("HTTP Serving at localhost:%d ...", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), router.Routes())
}

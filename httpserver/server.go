package httpserver

import (
	"net/http"
	"fmt"
	"github.com/W1llyu/lux/config"
	"github.com/W1llyu/lux/httpserver/router"
	"github.com/W1llyu/lux/websocket/utils"
)

func Run() {
	port := config.GetConf().Http.Port
	utils.Info(fmt.Sprintf("HTTP Serving at localhost:%d ...", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), router.Routes())
}

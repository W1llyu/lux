package httpserver

import (
	"net/http"
	"fmt"
	"github.com/Lux-go/config"
	"github.com/Lux-go/httpserver/router"
	"github.com/Lux-go/utils"
)

func Run() {
	port := config.GetConf().Http.Port
	utils.Info(fmt.Sprintf("HTTP Serving at localhost:%d ...", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), router.Routes())
}

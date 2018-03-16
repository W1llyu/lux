package http

import (
	"net/http"
	"fmt"
	"github.com/Lux-go/config"
)

func Run() {
	port := config.GetConf().Http.Port
	http.ListenAndServe(fmt.Sprintf(":%d", port))
}

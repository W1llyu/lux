package router

import (
	"sync"
	"github.com/drone/routes"
)

var (
	once sync.Once
	mux *routes.RouteMux
)

func Routes() *routes.RouteMux {
	once.Do(initRoutes)
	return mux
}

func initRoutes() {
	mux = routes.New()
	initSocketRoutes()
}

package router

import (
	"github.com/Lux-go/httpserver/handler"
)

func initSocketRoutes() {
	mux.Get("/sockets/rooms/:room", handler.GetRoomMemberCount)
}

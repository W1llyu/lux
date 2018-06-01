package router

import (
	"github.com/irelia_socket/httpserver/handler"
)

func initSocketRoutes() {
	mux.Get("/sockets/rooms/:room", handler.GetRoomMemberCount)
	mux.Get("/sockets/stat", handler.GetSocketStat)
}

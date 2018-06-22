package router

import (
	"github.com/irelia_socket/httpserver/handler"
)

func initSocketRoutes() {
	mux.Get("/sockets/rooms/:room", handler.GetRoomMemberCount)
	mux.Get("/sockets/count", handler.GetSocketCount)
	mux.Get("/sockets/:socketId", handler.GetSocketDetail)
	mux.Get("/sockets/client/stat", handler.GetClientStat)
}

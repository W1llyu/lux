package handler

import (
	"github.com/googollee/go-socket.io"
)

type Event struct {
	Name string
	Handler interface{}
}

func BindEvents(server *socketio.Server) {
	server.On("connection", onConnection)
	server.On("error", onError)
}

func getEvents() []Event {
	return []Event{
		{
			Name: "subscribe",
			Handler: onSubscribe,
		},
		{
			Name: "unsubscribe",
			Handler: onUnSubscribe,
		},
		{
			Name: "unsubscribeAll",
			Handler: onUnSubscribeAll,
		},
	}
}

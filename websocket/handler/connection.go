package handler

import (
	"github.com/Lux-go/common/utils"
	"github.com/googollee/go-socket.io"
)

func onConnection(socket socketio.Socket) {
	utils.Infof("Socket[%s] connected", socket.Id())
	socket.On("disconnection", onDisconnection)
	for _, event := range getEvents() {
		socket.On(event.Name, event.Handler)
	}
}

func onError(socket socketio.Socket, err error) {
	utils.Error(err, "")
}

func onDisconnection(socket socketio.Socket) {
	utils.Infof("Socket[%s] disconnected", socket.Id())
}

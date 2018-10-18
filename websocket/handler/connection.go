package handler

import (
	"github.com/W1llyu/lux/websocket/utils"
	"github.com/W1llyu/go-socket.io"
)

func onConnection(socket socketio.Socket) {
	utils.Infof("Socket[%s] connected", socket.Id())
	socket.Join(socket.Id())
	socket.On("disconnection", onDisconnection)
	for _, event := range getEvents() {
		socket.On(event.Name, event.Handler)
	}
}

func onError(socket socketio.Socket, err error) {
	utils.Error(err, "")
}

func onDisconnection(socket socketio.Socket) {
	if len(socket.Rooms()) > 0 {
		onLeave(socket, JoinMessage{socket.Rooms()})
	}
	utils.Infof("Socket[%s] disconnected", socket.Id())
}

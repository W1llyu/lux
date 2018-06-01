package handler

import (
	"github.com/irelia_socket/utils"
	"github.com/W1llyu/go-socket.io"
	"regexp"
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
	for _, room := range socket.Rooms() {
		match, _ := regexp.MatchString("irelia:room:\\d+", room)
		if match {
			leaveIreRoom(socket, room)
		}
	}
	if len(socket.Rooms()) > 0 {
		onLeave(socket, JoinMessage{socket.Rooms()})
	}
	utils.Infof("Socket[%s] disconnected", socket.Id())
}

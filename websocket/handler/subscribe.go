package handler

import (
	"github.com/googollee/go-socket.io"
	"github.com/Lux-go/common/utils"
)

type SubscribeMessage struct {
	Channels []string
}

func onSubscribe(socket socketio.Socket, msg SubscribeMessage) {
	for _, channel := range msg.Channels {
		socket.Join(channel)
	}
	utils.Infof("Socket[%s] Subscribe %s", socket.Id(), msg.Channels)
}

func onUnSubscribe(socket socketio.Socket, msg SubscribeMessage) {
	for _, channel := range msg.Channels {
		socket.Leave(channel)
	}
	utils.Infof("Socket[%s] Unsubscribe %s", socket.Id(), msg.Channels)
}

func onUnSubscribeAll(socket socketio.Socket) {
	for _, channel := range socket.Rooms() {
		socket.Leave(channel)
	}
	utils.Infof("Socket[%s] Unsubscribe")
}
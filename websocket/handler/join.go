package handler

import (
	"github.com/Lux-go/utils"
	"github.com/Lux-go/websocket/runtime"
	"github.com/googollee/go-socket.io"
	"time"
)

const (
	NOTIFICATION = "notification"
	JOINTYPE     = 1
	LEAVETYPE    = -1
)

type JoinMessage struct {
	Rooms []string `json:"rooms"`
}

type NotifyMessage struct {
	Room        string `json:"room"`
	MemberCount int    `json:"member_count"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	NotifiedAt  int64  `json:"notified_at"`
}

func onJoin(socket socketio.Socket, msg JoinMessage) {
	for _, room := range msg.Rooms {
		socket.Join(room)
		socket.BroadcastTo(room, NOTIFICATION, createNotifyMessage(room, socket.Id(), JOINTYPE))
	}
	utils.Infof("Socket[%s] Join %s", socket.Id(), msg.Rooms)
}

func onLeave(socket socketio.Socket, msg JoinMessage) {
	for _, room := range msg.Rooms {
		socket.Leave(room)
		socket.BroadcastTo(room, NOTIFICATION, createNotifyMessage(room, socket.Id(), LEAVETYPE))
	}
	utils.Infof("Socket[%s] Leave %s", socket.Id(), msg.Rooms)
}

func createNotifyMessage(room, name string, eventType int) NotifyMessage {
	return NotifyMessage{
		Room:        room,
		MemberCount: len(runtime.Server().Sockets(room)),
		Name:        name,
		Type:        eventType,
		NotifiedAt:  time.Now().UnixNano() / 1000000,
	}
}

package handler

import (
	"time"
	"github.com/W1llyu/lux/websocket/utils"
	"github.com/W1llyu/lux/websocket/core"
	"github.com/W1llyu/go-socket.io"
)

const (
	NOTIFICATION = "lux:notification"
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

func onJoin(socket socketio.Socket, msg JoinMessage) interface {} {
	notifyMsgs := make(map[string]NotifyMessage)
	for _, room := range msg.Rooms {
		socket.Join(room)
		notifyMsg := createNotifyMessage(room, socket.Id(), JOINTYPE)
		socket.BroadcastTo(room, NOTIFICATION, notifyMsg)
		notifyMsgs[room] = notifyMsg
	}
	utils.Infof("Socket[%s] Join %s", socket.Id(), msg.Rooms)
	return notifyMsgs
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
		MemberCount: len(core.Server().Sockets(room)),
		Name:        name,
		Type:        eventType,
		NotifiedAt:  time.Now().UnixNano() / 1000000,
	}
}

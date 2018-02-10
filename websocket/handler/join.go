package handler

import (
	"github.com/googollee/go-socket.io"
)

const (
	NOTIFICATION = "notification"
)

type JoinMessage struct {
	Rooms []string `json:"rooms"`
}

type NotifyMessage struct {
	Room string `json:"room"`
	MemberCount int32 `json:"member_count"`
	NewMember string `json:"new_member"`
	NotifiedAt int64 `json:"notified_at"`
}

func onJoin(socket socketio.Socket, msg JoinMessage) {
	for _, room := range msg.Rooms {
		socket.Join(room)
		//socket.BroadcastTo(room, NOTIFICATION)
	}
}



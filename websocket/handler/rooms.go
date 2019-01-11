package handler

import (
	"time"
	"encoding/json"
	"github.com/W1llyu/lux/websocket/utils"
	"github.com/W1llyu/lux/websocket/core"
	"github.com/W1llyu/go-socket.io"
	"github.com/W1llyu/gdao/xrmq"
	"github.com/streadway/amqp"
	"github.com/W1llyu/lux/websocket/constant"
)

const (
	JOIN_ROOM_CALLBACK_KEY = "socket.joinRoom.callback"
	LEAVE_ROOM_CALLBACK_KEY = "socket.leaveRoom.callback"
)

type JoinMessage struct {
	Rooms []string `json:"rooms"`
}

func onJoin(socket socketio.Socket, msg JoinMessage) {
	for _, room := range msg.Rooms {
		socket.Join(room)
		go notifyCallback(JOIN_ROOM_CALLBACK_KEY, socket, room)
	}
	utils.Infof("Socket[%s] Join %s", socket.Id(), msg.Rooms)
}

func onLeave(socket socketio.Socket, msg JoinMessage) {
	for _, room := range msg.Rooms {
		socket.Leave(room)
		go notifyCallback(LEAVE_ROOM_CALLBACK_KEY, socket, room)
	}
	utils.Infof("Socket[%s] Leave %s", socket.Id(), msg.Rooms)
}

type callbackMessage struct {
	SocketId    string `json:"socket_id"`
	Credential  string `json:"credential"`
	MessageAt   int64  `json:"message_at"`
	RoomName    string `json:"room_name"`
	MemberCount int    `json:"member_count"`
}

func getChannel() *xrmq.Channel {
	return xrmq.GetClient().GetChannel(xrmq.ExchangeCtx{
		Name:       constant.DEFAULT_KEY,
		Type:       "topic",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	})
}

func notifyCallback(key string, socket socketio.Socket, roomName string) {
	msg := callbackMessage{
		SocketId: socket.Id(),
		Credential: socket.Request().URL.Query().Get("credential"),
		MessageAt: time.Now().UnixNano() / 1000000,
		RoomName: roomName,
		MemberCount: len(core.Server().Sockets(roomName)),
	}
	byt, err := json.Marshal(msg)
	if err != nil {
		return
	}
	err = getChannel().Publish(
		constant.DEFAULT_KEY,
		key,
		false,
		false,
		amqp.Publishing {
			ContentType: "text/plain",
			Body: []byte(string(byt)),
		})
	if err != nil {
		utils.Error(err, "")
	}
}

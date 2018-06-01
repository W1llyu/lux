package handler

import (
	"fmt"
	"time"
	"github.com/W1llyu/go-socket.io"
	"github.com/irelia_socket/websocket/core"
	"crypto/md5"
	"regexp"
	"strconv"
	"errors"
	"github.com/W1llyu/gdao/xredis"
	"encoding/json"
	"github.com/irelia_socket/utils"
)

type IreRoomMessage struct {
	RoomId int `json:"room_id"`
	UserType string `json:"user_type"`
	UserId int `json:"user_id"`
}

type IreAckMessage struct {
	MemberCount int `json:"member_count"`
}

type IreSidekiqMessage struct {
	Class string `json:"class"`
	Args []interface {} `json:"args"`
	Retry bool `json:"retry"`
	Queue string `json:"queue"`
	Jid string `json:"jid"`
	CreatedAt float32 `json:"created_at"`
	EnqueuedAt float32 `json:"enqueued_at"`
}

func onJoinIreRoom(socket socketio.Socket, msg IreRoomMessage) interface {} {
	room := channelName(msg.RoomId)
	onJoin(socket, JoinMessage{[]string{room}})
	pushSidekiqTask(msg.RoomId, msg.UserType, msg.UserId, JOINTYPE)
	return &IreAckMessage{len(core.Server().Sockets(room))}
}

func onLeaveIreRoom(socket socketio.Socket, msg IreRoomMessage) {
	room := channelName(msg.RoomId)
	onLeave(socket, JoinMessage{[]string{room}})
	pushSidekiqTask(msg.RoomId, msg.UserType, msg.UserId, LEAVETYPE)
}

func leaveIreRoom(socket socketio.Socket, channel string) {
	roomId, err := getRoomIdFromChannel(channel)
	if err != nil {
		return
	}
	onLeaveIreRoom(socket, IreRoomMessage{RoomId: roomId})
}

func channelName(roomId int) string {
	return fmt.Sprintf("irelia:room:%d", roomId)
}

func getRoomIdFromChannel(channel string) (int, error) {
	reg := regexp.MustCompile("irelia:room:(\\d+)")
	res := reg.FindSubmatch([]byte(channel))
	if len(res) != 2 {
		return 0, errors.New("illegal irelia channel")
	}
	roomId, err := strconv.Atoi(string(res[1]))
	if err != nil {
		return 0, err
	}
	return roomId, nil
}

func pushSidekiqTask(roomId int, userType string, userId int, eventType int) {
	sidekiqMsg := &IreSidekiqMessage{
		Class: "Rooms::MemberCountNotificationWorker",
		Args: []interface{} {roomId, userType, userId, eventType},
		Retry: true,
		Queue: "notification",
		Jid: generateJid(),
	}
	client := xredis.GetNamedClient("irelia")
	defer client.Close()
	jsonStr, err := json.Marshal(sidekiqMsg)
	if err == nil {
		client.Rpush("queue:notification", string(jsonStr))
		client.Sadd("queues", "notification")
		utils.Infof("Enqueue Sidekiq Task %s", sidekiqMsg.Jid)
	}
}

func generateJid() string {
	data := []byte(fmt.Sprintf("%s", time.Now().UnixNano()))
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}



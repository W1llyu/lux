package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/W1llyu/lux/websocket/core"
)

func GetRoomMemberCount(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	room := params.Get(":room")
	res := make(map[string]interface{})
	res["member_count"] = len(core.Server().Sockets(room))
	body, _ := json.Marshal(res)
	fmt.Fprint(w, string(body))
}

type socketDetail struct {
	RemoteAddr string
	Referer string
}

func GetSocketDetail(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	socketId := params.Get(":socketId")
	conn := core.Sessions.Get(socketId)
	res := make(map[string]interface{})
	if conn == nil {
		res["socket"] = nil
	} else {
		remoteIp := conn.Request().Header.Get("X-Real-IP")
		if remoteIp == "" {
			remoteIp = conn.Request().RemoteAddr
		}
		s := socketDetail{
			RemoteAddr: remoteIp,
			Referer: conn.Request().Referer(),
		}
		res["socket"] = s
	}
	body, _ := json.Marshal(res)
	fmt.Fprint(w, string(body))
}

func GetSocketCount(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]interface{})
	res["sockets"] = make(map[string]interface{})
	res["sockets"].(map[string]interface{})["count"] = core.Server().Count()
	body, _ := json.Marshal(res)
	fmt.Fprint(w, string(body))
}

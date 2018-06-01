package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/irelia_socket/websocket/core"
)

func GetRoomMemberCount(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	room := params.Get(":room")
	res := make(map[string]interface{})
	res["member_count"] = len(core.Server().Sockets(room))
	body, _ := json.Marshal(res)
	fmt.Fprint(w, string(body))
}

func GetSocketStat(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]interface{})
	res["sockets"] = make(map[string]interface{})
	res["sockets"].(map[string]interface{})["count"] = core.Server().Count()
	body, _ := json.Marshal(res)
	fmt.Fprint(w, string(body))
}
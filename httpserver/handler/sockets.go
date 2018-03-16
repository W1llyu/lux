package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/Lux-go/websocket/core"
)

func GetRoomMemberCount(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	room := params.Get(":room")
	res := make(map[string]interface{})
	res["member_count"] = len(core.Server().Sockets(room))
	body, _ := json.Marshal(res)
	fmt.Fprint(w, string(body))
}

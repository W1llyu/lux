package handler

import (
	"fmt"
	"net/http"
	"github.com/W1llyu/lux/websocket/core"
	"encoding/json"
)

type socketStat struct {
	ConnCount int
	SocketIds []string
}

func GetClientStat(w http.ResponseWriter, r *http.Request) {
	var stat socketStat
	for sid := range core.Sessions.Sessions() {
		stat.SocketIds = append(stat.SocketIds, sid)
		stat.ConnCount += 1
	}
	res := make(map[string]interface{})
	res["stat"] = stat
	body, _ := json.Marshal(res)
	fmt.Fprint(w, string(body))
}
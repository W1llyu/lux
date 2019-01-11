package handler

import (
	"net/http"
	"github.com/segmentio/ksuid"
	"encoding/json"
	"fmt"
	"time"
	"github.com/W1llyu/gdao/xredis"
	"github.com/W1llyu/lux/websocket/constant"
)

func GetAccessToken(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]interface{})
	token := ksuid.New().String()
	client := xredis.GetClient()
	defer client.Close()
	key := fmt.Sprintf("%s_%s", constant.SOCKET_SECURITY_KEY, token)
	err := client.Set(key, "1")
	if err == nil {
		client.Expire(key, 5 * time.Minute)
		res["access_token"] = token
	} else {
		res["access_token"] = nil
	}
	res["timestamp"] = time.Now().Unix()
	body, _ := json.Marshal(res)
	fmt.Fprint(w, string(body))
}

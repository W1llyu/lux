package core

import (
	"errors"
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"net/http"
	"fmt"
	"github.com/W1llyu/lux/websocket/constant"
	"github.com/W1llyu/gdao/xredis"
	"github.com/W1llyu/lux/websocket/utils"
	"strings"
	"strconv"
)

func authRequest(r *http.Request) error {
	timestamp, err := strconv.ParseInt(r.URL.Query().Get("timestamp"), 10, 64)
	secret := r.URL.Query().Get("secret")
	accessToken := r.URL.Query().Get("access_token")
	if err != nil || secret == "" || accessToken == "" {
		return errors.New("invalid request")
	}
	if checkAccessToken(accessToken) && checkSecret(timestamp, secret) {
		return nil
	} else {
		return errors.New("invalid request")
	}
}

func checkAccessToken(token string) bool {
	client := xredis.GetClient()
	defer client.Close()
	key := fmt.Sprintf("%s_%s", constant.SOCKET_SECURITY_KEY, token)
	ok, _ := client.Exists(key)
	if ok {
		client.Del(key)
	}
	return ok
}

func checkSecret(timestamp int64, secret string) bool {
	return strings.EqualFold(utils.Encrypt(timestamp), secret)
}

func authToken(clientType string, token string) bool {
	return true
}

func newSocketId(r *http.Request) string {
	return generateConnId(r.URL.Query().Get("token"))
}

func generateConnId(token string) string {
	hash := token
	buf := bytes.NewBuffer(nil)
	sum := md5.Sum([]byte(hash))
	encoder := base64.NewEncoder(base64.URLEncoding, buf)
	encoder.Write(sum[:])
	encoder.Close()
	return buf.String()[:20]
}

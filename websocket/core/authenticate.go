package core

import (
	"errors"
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"net/http"
	"net"
	"time"
	"github.com/irelia_socket/config"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)

var (
	httpClient = &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second * 30)    // 建立连接超时
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 30))   // 发送接收数据超时
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 30,
		},
	}
)

func authRequest(r *http.Request) error {
	if r.URL.Query().Get("token") == "" || r.URL.Query().Get("client") == "" {
		return errors.New("unknown request")
	}
	if !authToken(r.URL.Query().Get("client"), r.URL.Query().Get("token")) {
		return errors.New("request authenticate failed")
	}
	return nil
}

func authToken(clientType string, token string) bool {
	userType := ""
	switch clientType {
	case "ios":
		userType = "Guest"
	case "android":
		userType = "Guest"
	case "admin":
		userType = "AdminUser"
	default:
		return false
	}
	body := make(map[string]string)
	body["type"] = userType
	body["token"] = token
	bodyStr, _ := json.Marshal(body)
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/api/socket_auth", config.GetConf().Irelia.Host), bytes.NewBuffer(bodyStr))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return false
	}
	var response *http.Response
	response, err = httpClient.Do(request)
	if err != nil {
		log.Printf("[FAILED] Cannot get response from irelia, %s", err)
		return false
	}
	if response.StatusCode != 200 {
		return false
	}
	resBodyStr, _ := ioutil.ReadAll(response.Body)
	resBody := make(map[string]interface{})
	err = json.Unmarshal(resBodyStr, &resBody)
	if err != nil {
		return false
	}
	response.Body.Close()
	return resBody["ok"].(bool)
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

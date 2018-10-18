package core

import (
	"errors"
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"net/http"
	"net"
	"time"
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

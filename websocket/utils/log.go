package utils

import (
	"fmt"
	"log"
	"github.com/irelia_socket/config"
)

func WarnOnError(err error, msg string) {
	if !config.GetConf().Websocket.LogOpen {
		return
	}
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func Fatal(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func Info(msg string) {
	if !config.GetConf().Websocket.LogOpen {
		return
	}
	log.Printf("[INFO] %s", msg)
}

func Infof(format string, a ...interface{}) {
	if !config.GetConf().Websocket.LogOpen {
		return
	}
	log.Printf("[INFO] %s", fmt.Sprintf(format, a...))
}

func Error(err error, msg string) {
	if !config.GetConf().Websocket.LogOpen {
		return
	}
	log.Printf("[ERROR] %s %s", msg, err)
}

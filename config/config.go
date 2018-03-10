package config

import (
	"fmt"
	"os"
	"sync"
	"github.com/W1llyu/gdao/config"
)

type Config struct {
	Websocket *WebsocketConf
}

type WebsocketConf struct {
	Port int `toml:"port"`
}

var (
	cfg  *Config
	once sync.Once
)

func GetConf() *Config {
	once.Do(initConf)
	return cfg
}

func initConf() {
	gopath := os.Getenv("GOPATH")
	config.LoadConf(&cfg, fmt.Sprintf("%s/config/lux/config.toml", gopath))
}

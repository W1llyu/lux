package config

import (
	"sync"
	"github.com/W1llyu/gdao/config"
)

type Config struct {
	Websocket *WebsocketConf
	Http *HttpConf
	Irelia *IreliaConf
}

type WebsocketConf struct {
	Port int `toml:"port"`
}

type HttpConf struct {
	Port int `toml:"port"`
}

type IreliaConf struct {
	Host string `toml:"host"`
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
	config.LoadConf(&cfg, "./config//config.toml")
}

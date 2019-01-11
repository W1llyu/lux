package config

import (
	"sync"
	"github.com/W1llyu/gdao/config"
)

type Config struct {
	Websocket *WebsocketConf
	Http *HttpConf
}

type WebsocketConf struct {
	Port int `toml:"port"`
	LogOpen bool `toml:"logOpen"`
	Auth bool `toml:"auth"`
}

type HttpConf struct {
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
	config.LoadConf(&cfg, "./config//config.toml")
}

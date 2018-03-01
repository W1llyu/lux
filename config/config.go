package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Lux-go/utils"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Config struct {
	Websocket *WebsocketConf
}

type duration struct {
	time.Duration
}

type WebsocketConf struct {
	Port int `toml:"port"`
}

var (
	cfg  *Config
	once sync.Once
)

func GetConf() *Config {
	once.Do(loadConf)
	return cfg
}

func loadConf() {
	gopath := os.Getenv("GOPATH")
	filePath, err := filepath.Abs(fmt.Sprintf("%s/config/lux/config.toml", gopath))
	if err != nil {
		panic(err)
	}
	utils.Infof("parse toml file. filePath: %s\n", filePath)
	if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
		panic(err)
	}
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

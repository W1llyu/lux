package config

import (
	"fmt"
	"time"
	"sync"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"os"
)

type Config struct {
	Redis map[string]*RedisConf
	RabbitMQ map[string]*AmqpConf
}

type duration struct {
	time.Duration
}

type RedisConf struct {
	Addr string
	Database int
	MaxIdle int `toml:"max_idle"`
	MaxActive int `toml:"max_active"`
	IdleTimeout duration `toml:"idle_timeout"`
}

type AmqpConf struct {
	Addr string
}

var (
	cfg *Config
	once sync.Once
)

func GetConf() *Config {
	once.Do(loadConf)
	return cfg
}

func loadConf() {
	gopath := os.Getenv("GOPATH")
	filePath, err := filepath.Abs(fmt.Sprintf("%s/config/lux/config.toml", gopath))
	//filePath, err := filepath.Abs("./common/config/config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("parse toml file. filePath: %s\n", filePath)
	if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
		panic(err)
	}
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
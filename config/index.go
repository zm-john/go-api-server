package config

import (
	"fmt"
	ini "gopkg.in/ini.v1"
	"os"
)

// ServerConfig App
type ServerConfig struct {
	App   `ini:"app"`
	JWT   `ini:"jwt"`
	Mysql `ini:"mysql"`
}

var conf *ServerConfig

// Get get server config
func Get() ServerConfig {
	return *conf
}

func init() {
	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	confPath := fmt.Sprintf("%s/%s", wd, "config.ini")

	file, err := ini.Load(confPath)
	if err != nil {
		panic(err)
	}

	conf = new(ServerConfig)
	err = file.MapTo(conf)
	if err != nil {
		panic(err)
	}
}

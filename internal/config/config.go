package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Server struct {
		BindAddr string `toml:"bind_address"`
	}
	Database struct {
		URI  string `toml:"uri"`
		Name string `toml:"name"`
	}
	Shorten struct {
		Length int `toml:"length"`
	}
}

func NewConfig() (*Config, error) {
	var config Config

	_, err := toml.DecodeFile("./config/local.toml", &config)
	if err != nil {
		return nil, err
	}

	log.Print("Configuration read")
	return &config, nil
}

package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ID    int    `yaml:"id"`
	Token string `yaml:"token"`
}

func Read() (Config, error) {
	var c Config
	data, err := ioutil.ReadFile("config.yaml")
	if err == nil {
		yaml.Unmarshal(data, &c)
	}
	return c, err
}

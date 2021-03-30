package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name  string `yaml: name`
	ID    int    `yaml:"id"`
	Token string `yaml:"token"`
}

func parseYaml(data []byte) Config {
	var c Config
	yaml.Unmarshal(data, &c)
	return c
}

func Read() (Config, error) {
	data, err := ioutil.ReadFile("config.yaml")
	if err == nil {
		return parseYaml(data), nil
	}
	return Config{}, err
}

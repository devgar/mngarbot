package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var ConfigPath string = "config.yaml"

func init() {
	exe, err := os.Executable()
	if err == nil {
		ConfigPath = path.Join(path.Dir(exe), "config.yaml")
	} else {
		fmt.Println("Can't get executable folder")
	}
}

type Config struct {
	Name  string `yaml:"name"`
	ID    int    `yaml:"id"`
	Token string `yaml:"token"`
}

func parseYaml(data []byte) Config {
	var c Config
	yaml.Unmarshal(data, &c)
	return c
}

func Read() (Config, error) {
	data, err := ioutil.ReadFile(ConfigPath)
	if err == nil {
		return parseYaml(data), nil
	}
	return Config{}, err
}

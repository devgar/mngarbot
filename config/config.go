package config

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/yaml.v3"
)

var (
	config_path string = "config.yaml"
	config      Config = Config{}
	adminErr    error
	pkgName     string = "mngarbot"
)

func buildRoutes() []string {
	places := []string{}
	if exe, err := os.Executable(); err == nil {
		dir := path.Dir(exe)
		pkgName = path.Base(exe)
		places = append(places, path.Join(dir, "config.yaml"))
	}
	if cfgDir, err := os.UserConfigDir(); err == nil {
		places = append(places, path.Join(cfgDir, pkgName, "config.yaml"))
	}
	return places
}

type Config struct {
	Name  string `yaml:"name"`
	ID    int64  `yaml:"id"`
	Token string `yaml:"token"`
}

func parseYaml(data []byte) Config {
	var c Config
	yaml.Unmarshal(data, &c)
	return c
}

func read() (Config, error) {
	data, err := ioutil.ReadFile(config_path)
	if err == nil {
		return parseYaml(data), nil
	}
	return Config{}, err
}

func Get() Config {
	return config
}

func init() {
	TOKEN := os.Getenv("TOKEN")
	ADMIN, _ := strconv.ParseInt(os.Getenv("ADMIN"), 10, 64)
	for _, route := range buildRoutes() {
		if _, err := os.Stat(route); err == nil {
			config_path = route
			break
		}
	}
	config, _ = read()
	if config.Token == "" {
		config.Token = TOKEN
	}
	if config.ID == 0 {
		config.ID = ADMIN
	}
}

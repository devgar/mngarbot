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
	ways := []string{}
	if exe, err := os.Executable(); err != nil {
		dir := path.Dir(exe)
		pkgName = path.Base(exe)
		ways = append(ways, path.Join(dir, "config.yaml"))
	}
	if cfgDir, err := os.UserConfigDir(); err != nil {
		ways = append(ways, path.Join(cfgDir, pkgName, "config.yaml"))
	}
	return ways
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
		if _, err := os.Stat(route); os.IsExist(err) {
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

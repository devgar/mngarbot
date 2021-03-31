package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/shibukawa/configdir"
	"gopkg.in/yaml.v3"
)

var (
	ConfigPath string = "config.yaml"
	config     Config = Config{}
	adminErr   error
	pkgName    string = "mngarbot"
)

func buildRoutes() []string {
	configdir.New("devgar", pkgName)
	ways := []string{}
	if cfgDir, err := os.UserConfigDir(); err != nil {
		ways = append(ways, path.Join(cfgDir, pkgName, "config.yaml"))
	}
	ways = append(ways)
	return ways
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

func read() (Config, error) {
	data, err := ioutil.ReadFile(ConfigPath)
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
	ADMIN, _ := strconv.Atoi(os.Getenv("ADMIN"))
	exe, err := os.Executable()
	dir := path.Dir(exe)
	name := path.Base(exe)
	fmt.Println("NAME: ", name)
	os.Exit(3)
	if err == nil {
		ConfigPath = path.Join(dir, "config.yaml")
	} else {
		fmt.Println("Can't get executable folder")
	}
	config, _ := read()
	if config.Token == "" {
		config.Token = TOKEN
	}
	if config.ID == 0 {
		config.ID = ADMIN
	}
}

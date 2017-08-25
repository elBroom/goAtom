package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"log"

	yaml "gopkg.in/yaml.v2"
)

func GetYamlConfig(name string, config interface{}) error {
	dir := os.Getenv("GOPATH") + "/src/github.com/elBroom/goAtom/config"
	configPath := fmt.Sprintf("%s/%s.yml", dir, name)
	configContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("can't read config %q: %s", configPath, err)
	}

	if err = yaml.Unmarshal(configContent, config); err != nil {
		return fmt.Errorf("invalid yaml in config %q: %s", configPath, err)
	}

	return nil
}

type App struct {
	RequestWaitInQueueTimeout time.Duration `yaml:"request_wait_in_queue_timeout"`
	Workers                   int
	Port                      int
}

var app App
var RequestWaitInQueueTimeout time.Duration

func GetApp() App {
	return app
}

type Redis struct {
	Host     string
	Port     int
	Password string
	Database int
}

var redis Redis

func GetRedis() Redis {
	return redis
}

type Sql struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

var sql Sql

func GetSql() Sql {
	return sql
}

func init() {
	if err := GetYamlConfig("app", &app); err != nil {
		log.Fatalf("can't read app config: %s", err)
	}
	RequestWaitInQueueTimeout = time.Millisecond * app.RequestWaitInQueueTimeout

	if err := GetYamlConfig("sql", &sql); err != nil {
		log.Fatalf("can't read sql config: %s", err)
	}
	if err := GetYamlConfig("redis", &redis); err != nil {
		log.Fatalf("can't read redis config: %s", err)
	}
}

package config

import (
	"time"
)

type App struct {
	RequestWaitInQueueTimeout time.Duration `yaml:"request_wait_in_queue_timeout"`
	Workers int
	Port int
}
var app = App{RequestWaitInQueueTimeout: 100, Workers:20, Port:8080}
var RequestWaitInQueueTimeout = time.Millisecond * app.RequestWaitInQueueTimeout
func GetApp() App {
	return app
}

type Redis struct {
	Host string
	Port int
	Password string
	Database int
}
var redis = Redis{Host:"localhost", Port:6379, Database:0}
func GetRedis() Redis {
	return redis
}

type Sqlite struct {
	DatabasePath string `yaml:"database_path"`
}
var sqlite = Sqlite{DatabasePath:"./radish.db"}
func GetSqlite() Sqlite {
	return sqlite
}

// TODO: parse yaml file

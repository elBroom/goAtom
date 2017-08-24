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
var redis = Redis{Host:"elbroom.ru", Port:6379, Database:0}
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
var sqlite = Sql{Username:"go_atom_user", Password:"RvlAEHQFDNC1", Host:"elbroom.ru", Port:5432, Database:"go_atom"}
func GetSql() Sql {
	return sqlite
}

// TODO: parse yaml file

package db

import (
	"github.com/go-redis/redis"
)

var redis_connect  (*redis.Client)  = nil

func Redis_init() *redis.Client {
	if redis_connect != nil {
		return redis_connect
	}

	redis_connect := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//defer connect.Close() in main  =(
	return redis_connect
}
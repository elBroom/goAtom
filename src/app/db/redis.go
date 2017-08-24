package db

import (
	"app/config"
	"github.com/go-redis/redis"
	"strconv"
)

var redis_connect  (*redis.Client)  = nil

func Redis_init() *redis.Client {
	if redis_connect != nil {
		return redis_connect
	}

	cfg := config.GetRedis()
	redis_connect := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	//defer connect.Close() in main  =(
	return redis_connect
}
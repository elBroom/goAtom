package main

import (
	"./config"
	"./server"
	"./db"
	"log"
	"strconv"
)

func main() {
	sqlite_connect := db.Sqlite_connect()
	defer sqlite_connect.Close()
	//
	redis_connect := db.Redis_init()
	defer redis_connect.Close()

	cfg := config.GetApp()
	log.Printf("Start server: 127.0.0.1:%d\n", cfg.Port)
	log.Fatal(server.RunHTTPServer(":" + strconv.Itoa(cfg.Port)))
}

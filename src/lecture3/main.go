package main

import (
	"./server"
	"./db"
	"log"
)

func main() {
	sqlite_connect := db.Sqlite_connect()
	defer sqlite_connect.Close()

	redis_connect := db.Redis_init()
	defer redis_connect.Close()

	log.Println("Start server: 127.0.0.1:8080")
	log.Fatal(server.RunHTTPServer(":8080"))
}

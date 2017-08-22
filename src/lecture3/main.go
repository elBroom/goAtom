package main

import (
	"./server"
	"./db"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db_connect := db.Init()
	defer db_connect.Close()

	log.Fatal(server.RunHTTPServer(":8080"))
}

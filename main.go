package main

import (
	"log"
	"strconv"

	"github.com/elBroom/goAtom/app/config"
	"github.com/elBroom/goAtom/app/db"
	"github.com/elBroom/goAtom/app/server"
)

func main() {
	sql_connect := db.Sql_connect()
	defer sql_connect.Close()

	redis_connect := db.Redis_init()
	defer redis_connect.Close()

	cfg := config.GetApp()
	log.Printf("Start server: 127.0.0.1:%d\n", cfg.Port)
	log.Fatal(server.RunHTTPServer(":" + strconv.Itoa(cfg.Port)))
}

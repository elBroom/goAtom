package main

import (
	"./server"
	"log"
)

func main() {
	log.Fatal(server.RunHTTPServer(":8080"))
}

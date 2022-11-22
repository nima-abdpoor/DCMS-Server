package main

import (
	"DCMS/api"
	"log"
)

const (
	serverAddress = "0.0.0.0:8080"
)

func main() {
	server := api.NewServer()
	err := server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start the server... ", err.Error())
	}
}

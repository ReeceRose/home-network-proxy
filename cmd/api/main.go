package main

import (
	"log"

	"github.com/ReeceRose/home-network-proxy/internal/api/server"
	"github.com/ReeceRose/home-network-proxy/internal/database"
)

func main() {
	log.Default().Println("Home Network Proxy API")

	server := server.New()

	database, err := database.Instance()
	if err != nil {
		panic(err)
	}
	defer database.Disconnect()

	server.Start()
}

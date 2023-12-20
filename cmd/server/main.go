package main

import (
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client"
	"github.com/kevinrudde/gophercraft/internal/network/server"
	"log"
)

func main() {
	networkServer := server.NewServer(":25565")
	fmt.Println("Listening on :25565")

	client.InitializeClientPacketProcessors()

	log.Fatal(networkServer.Start())
}

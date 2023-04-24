package main

import (
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network/server"
	"log"
)

func main() {
	networkServer := server.NewServer(":25565")
	fmt.Println("Listening on :25565")

	log.Fatal(networkServer.Start())
}

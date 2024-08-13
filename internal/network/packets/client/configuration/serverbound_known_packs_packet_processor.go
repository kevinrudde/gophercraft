package configuration

import (
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"log"
)

func ProcessServerboundKnownPacksPacket(connection *networkplayer.PlayerConnection, packet *ServerboundKnownPacksPacket) error {
	log.Println("Received known packs from client")

	return nil
}

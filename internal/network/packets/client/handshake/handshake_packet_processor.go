package handshake

import (
	"github.com/kevinrudde/gophercraft/internal/network"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"log"
)

func ProcessHandshakePacket(connection *networkplayer.PlayerConnection, packet *HandshakePacket) error {
	connection.ConnectionState = network.ConnectionState(packet.NextState)

	log.Println(packet.ServerAddress)
	log.Println(packet.ProtocolVersion)
	log.Println(packet.ServerPort)

	return nil
}

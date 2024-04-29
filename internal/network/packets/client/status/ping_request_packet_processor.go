package status

import (
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/handshake"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func ProcessPingRequestPacket(connection *networkplayer.PlayerConnection, packet *PingRequestPacket) error {
	response := &handshake.PingResponsePacket{
		Payload: packet.Payload,
	}

	err := connection.SendPacket(response)
	if err != nil {
		return err
	}

	return connection.Close()
}

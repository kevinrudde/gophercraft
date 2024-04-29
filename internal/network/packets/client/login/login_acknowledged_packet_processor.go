package login

import (
	"github.com/kevinrudde/gophercraft/internal/network"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func ProcessLoginAcknowledgedPacket(connection *networkplayer.PlayerConnection, packet *LoginAcknowledgedPacket) error {
	connection.ConnectionState = network.Play

	return nil
}

package login

import (
	"fmt"

	"github.com/kevinrudde/gophercraft/internal/network"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func ProcessLoginAcknowledgedPacket(connection *networkplayer.PlayerConnection, packet *LoginAcknowledgedPacket) error {
	connection.ConnectionState = network.Configuration
	fmt.Println("Login acknowledged, switching to configuration state")

	return nil
}

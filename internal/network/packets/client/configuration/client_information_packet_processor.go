package configuration

import (
	"fmt"

	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func ProcessClientInformationPacket(connection *networkplayer.PlayerConnection, packet *ClientInformationPacket) error {
	fmt.Println("Processing configuration packet")
	fmt.Println(packet.Locale)

	return nil
}

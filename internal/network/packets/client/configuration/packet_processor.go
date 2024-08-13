package configuration

import (
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func CallProcessor(connection *networkplayer.PlayerConnection, packet common.ClientPacket) error {
	switch packet.(type) {
	case *ClientInformationPacket:
		return ProcessClientInformationPacket(connection, packet.(*ClientInformationPacket))
	case *ServerboundKnownPacksPacket:
		return ProcessServerboundKnownPacksPacket(connection, packet.(*ServerboundKnownPacksPacket))
	}

	return nil
}

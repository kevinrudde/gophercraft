package status

import (
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func CallProcessor(connection *networkplayer.PlayerConnection, packet common.ClientPacket) error {
	switch packet.(type) {
	case *StatusRequestPacket:
		return ProcessStatusRequestPacket(connection, packet.(*StatusRequestPacket))
	case *PingRequestPacket:
		return ProcessPingRequestPacket(connection, packet.(*PingRequestPacket))
	}

	return nil
}

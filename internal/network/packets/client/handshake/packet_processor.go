package handshake

import (
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func CallProcessor(connection *networkplayer.PlayerConnection, packet common.ClientPacket) error {
	switch packet.(type) {
	case *HandshakePacket:
		return ProcessHandshakePacket(connection, packet.(*HandshakePacket))
	}

	return nil
}

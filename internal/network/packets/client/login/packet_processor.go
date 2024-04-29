package login

import (
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func CallProcessor(connection *networkplayer.PlayerConnection, packet common.ClientPacket) error {
	switch packet.(type) {
	case *LoginStartPacket:
		return ProcessLoginStartPacket(connection, packet.(*LoginStartPacket))
	case *EncryptionResponsePacket:
		return ProcessEncryptionResponsePacket(connection, packet.(*EncryptionResponsePacket))
	case *LoginAcknowledgedPacket:
		return ProcessLoginAcknowledgedPacket(connection, packet.(*LoginAcknowledgedPacket))
	}

	return nil
}

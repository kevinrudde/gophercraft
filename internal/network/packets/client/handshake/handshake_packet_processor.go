package handshake

import (
	"errors"
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"reflect"
)

func ProcessHandshakePacket(connection *networkplayer.PlayerConnection, p common.ClientPacket) error {
	packet, ok := p.(*HandshakePacket)
	if !ok {
		return errors.New("expected HandshakePacket, but got " + reflect.TypeOf(p).String())
	}

	connection.ConnectionState = network.ConnectionState(packet.NextState)

	return nil
}

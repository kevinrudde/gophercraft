package status

import (
	"errors"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/handshake"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"reflect"
)

func ProcessPingRequestPacket(connection *networkplayer.PlayerConnection, p common.ClientPacket) error {
	packet, ok := p.(*PingRequestPacket)
	if !ok {
		return errors.New("expected PingRequestPacket, but got " + reflect.TypeOf(p).String())
	}

	response := &handshake.PingResponsePacket{
		Payload: packet.Payload,
	}

	err := connection.SendPacket(response)
	if err != nil {
		return err
	}

	return connection.Close()
}

package login

import (
	"errors"
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"reflect"
)

func ProcessLoginAcknowledgedPacket(connection *networkplayer.PlayerConnection, p common.ClientPacket) error {
	_, ok := p.(*LoginAcknowledgedPacket)
	if !ok {
		return errors.New("expected LoginAcknowledgedPacket, but got " + reflect.TypeOf(p).String())
	}

	connection.ConnectionState = network.Play

	return nil
}

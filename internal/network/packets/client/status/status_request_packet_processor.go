package status

import (
	"errors"
	"github.com/kevinrudde/gophercraft/internal/event"
	"github.com/kevinrudde/gophercraft/internal/event/server"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/handshake"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"reflect"
)

func ProcessStatusRequestPacket(connection *networkplayer.PlayerConnection, p common.ClientPacket) error {
	_, ok := p.(*StatusRequestPacket)
	if !ok {
		return errors.New("expected StatusRequestPacket, but got " + reflect.TypeOf(p).String())
	}

	serverListEvent := &server.ServerListPingEvent{}

	var err error

	err = event.Call(serverListEvent, func() {
		response := &handshake.StatusResponsePacket{
			Payload: serverListEvent.ResponseData.GetPingResponse(),
		}

		err = connection.SendPacket(response)
	})

	return err
}

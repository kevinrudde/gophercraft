package status

import (
	"github.com/kevinrudde/gophercraft/internal/event"
	"github.com/kevinrudde/gophercraft/internal/event/server"
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/handshake"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func ProcessStatusRequestPacket(connection *networkplayer.PlayerConnection, _ *StatusRequestPacket) error {
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

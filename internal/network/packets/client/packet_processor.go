package client

import (
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func ProcessPacket(connection *networkplayer.PlayerConnection, packetId int, body []byte) error {
	buffer := network.CreateBufferWithBuf(body)
	var packet common.ClientPacket
	var err error

	switch connection.ConnectionState {

	case network.Handshake:
		packet, err = CreateHandshakePacket(packetId, buffer)
		break
	case network.Status:
		packet, err = CreateStatusPacket(packetId, buffer)
		break
	case network.Login:
		packet, err = CreateLoginPacket(packetId, buffer)
		break
	case network.Configuration:
		packet, err = CreateConfigurationPacket(packetId, buffer)
		break
	case network.Play:
		packet, err = CreatePlayPacket(packetId, buffer)
		break
	default:
		err = fmt.Errorf("unknown state with id %d", connection.ConnectionState)
	}

	if err != nil {
		return err
	}

	if packet == nil {
		return nil
	}

	return CallProcessor(connection, packet)
}

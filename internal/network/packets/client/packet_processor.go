package client

import (
	"errors"
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"reflect"
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
	case network.Play:
		packet, err = CreatePlayPacket(packetId, buffer)
		break
	case network.Unknown:
	default:
		err = errors.New("unknown state")
	}

	if err != nil {
		return err
	}

	fmt.Println("Process packet: ", reflect.TypeOf(packet).String())
	processor, ok := PacketProcessors[reflect.TypeOf(packet).String()]
	if !ok {
		return errors.New(fmt.Sprintf("PacketId %d does not exists", packetId))
	}

	return processor(connection, packet)
}

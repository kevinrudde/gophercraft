package client

import (
	"errors"
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func ProcessPacket(connection networkplayer.PlayerConnection, packetId int, body []byte) error {
	buffer := network.CreateBufferWithBuf(body)
	var packet Packet
	var err error

	switch connection.ConnectionState {

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
		err = errors.New("Unknown state")
	}

	if err != nil {
		return err
	}

	fmt.Println("Process packet: ", packet)

	return nil
}

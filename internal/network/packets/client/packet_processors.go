package client

import (
	"errors"

	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/configuration"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/handshake"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/login"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/status"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
)

func CallProcessor(connection *networkplayer.PlayerConnection, packet common.ClientPacket) error {
	switch connection.ConnectionState {
	case network.Handshake:
		return handshake.CallProcessor(connection, packet)
	case network.Status:
		return status.CallProcessor(connection, packet)
	case network.Login:
		return login.CallProcessor(connection, packet)
	case network.Configuration:
		return configuration.CallProcessor(connection, packet)
	default:
		return errors.New("unknown state")
	}
}

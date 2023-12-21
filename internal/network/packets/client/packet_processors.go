package client

import (
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/handshake"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/status"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"reflect"
)

type ProcessorFunc func(connection *networkplayer.PlayerConnection, packet common.ClientPacket) error

var PacketProcessors = make(map[string]ProcessorFunc)

func RegisterProcessor(packetType string, processor ProcessorFunc) {
	PacketProcessors[packetType] = processor
}

func InitializeClientPacketProcessors() {
	// Handshake
	RegisterProcessor(reflect.TypeOf(&handshake.HandshakePacket{}).String(), handshake.ProcessHandshakePacket)

	// Status
	RegisterProcessor(reflect.TypeOf(&status.StatusRequestPacket{}).String(), status.ProcessStatusRequestPacket)
}

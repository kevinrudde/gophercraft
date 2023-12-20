package client

import (
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/status"
	"reflect"
)

type ProcessorFunc func(packet common.Packet) error

var PacketProcessors = make(map[string]ProcessorFunc)

func RegisterProcessor(packetType string, processor ProcessorFunc) {
	PacketProcessors[packetType] = processor
}

func InitializeClientPacketProcessors() {
	RegisterProcessor(reflect.TypeOf(&status.HandshakePacket{}).String(), status.ProcessHandshakePacket)
}

package client

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

var playPacketMap = map[int]Packet{}

func CreatePlayPacket(packetId int, buffer network.Buffer) (Packet, error) {
	packet := playPacketMap[packetId]
	err := packet.From(buffer)
	return packet, err
}

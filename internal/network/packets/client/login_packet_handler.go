package client

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

var loginPacketMap = map[int]Packet{}

func CreateLoginPacket(packetId int, buffer network.Buffer) (Packet, error) {
	packet := loginPacketMap[packetId]
	err := packet.From(buffer)
	return packet, err
}

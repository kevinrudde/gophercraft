package client

import (
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
)

var loginPacketMap = map[int]common.ClientPacket{}

func CreateLoginPacket(packetId int, buffer network.Buffer) (common.ClientPacket, error) {
	packet := loginPacketMap[packetId]
	err := packet.From(buffer)
	return packet, err
}

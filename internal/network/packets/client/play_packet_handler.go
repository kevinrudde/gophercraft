package client

import (
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
)

var playPacketMap = map[int]common.ClientPacket{}

func CreatePlayPacket(packetId int, buffer network.Buffer) (common.ClientPacket, error) {
	packet := playPacketMap[packetId]
	err := packet.From(buffer)
	return packet, err
}

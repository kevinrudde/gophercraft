package client

import (
	"errors"
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/status"
)

var statusPacketMap = map[int]common.ClientPacket{
	0x00: &status.StatusRequestPacket{},
	0x01: &status.PingRequestPacket{},
}

func CreateStatusPacket(packetId int, buffer network.Buffer) (common.ClientPacket, error) {
	packet, ok := statusPacketMap[packetId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("PacketId %d does not exists", packetId))
	}

	err := packet.From(buffer)
	return packet, err
}

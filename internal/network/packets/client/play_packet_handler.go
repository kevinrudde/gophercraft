package client

import (
	"errors"
	"fmt"

	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
)

var playPacketMap = map[int]common.ClientPacket{}

func CreatePlayPacket(packetId int, buffer network.Buffer) (common.ClientPacket, error) {
	packet, ok := playPacketMap[packetId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("PacketId %d does not exists", packetId))
	}

	err := packet.From(buffer)
	return packet, err
}

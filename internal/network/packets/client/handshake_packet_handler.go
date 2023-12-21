package client

import (
	"errors"
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/handshake"
)

var handshakePacketMap = map[int]common.ClientPacket{
	0x00: &handshake.HandshakePacket{},
}

func CreateHandshakePacket(packetId int, buffer network.Buffer) (common.ClientPacket, error) {
	packet, ok := handshakePacketMap[packetId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("PacketId %d does not exists", packetId))
	}

	err := packet.From(buffer)
	return packet, err
}

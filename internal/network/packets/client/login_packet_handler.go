package client

import (
	"errors"
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/login"
)

var loginPacketMap = map[int]common.ClientPacket{
	0x00: &login.LoginStartPacket{},
	0x01: &login.EncryptionResponsePacket{},
}

func CreateLoginPacket(packetId int, buffer network.Buffer) (common.ClientPacket, error) {
	packet, ok := loginPacketMap[packetId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("PacketId %d does not exists", packetId))
	}

	err := packet.From(buffer)
	return packet, err
}

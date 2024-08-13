package client

import (
	"log"

	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/configuration"
)

var configurationPacketMap = map[int]common.ClientPacket{
	0x00: &configuration.ClientInformationPacket{},
	0x07: &configuration.ServerboundKnownPacksPacket{},
}

func CreateConfigurationPacket(packetId int, buffer network.Buffer) (common.ClientPacket, error) {
	packet, ok := configurationPacketMap[packetId]
	if !ok {
		log.Printf("PacketId %d does not exists", packetId)
		return nil, nil
	}

	err := packet.From(buffer)
	return packet, err
}

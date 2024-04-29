package client

import (
	"errors"
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/login"
	"reflect"
)

var loginPacketMap = map[int]reflect.Type{
	0x00: reflect.TypeOf(&login.LoginStartPacket{}),
	0x01: reflect.TypeOf(&login.EncryptionResponsePacket{}),
	0x03: reflect.TypeOf(&login.LoginAcknowledgedPacket{}),
}

func CreateLoginPacket(packetId int, buffer network.Buffer) (common.ClientPacket, error) {
	packetType, ok := loginPacketMap[packetId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("PacketId %d does not exists", packetId))
	}

	packet := reflect.New(packetType).Interface().(common.ClientPacket)

	err := packet.From(buffer)
	return packet, err
}

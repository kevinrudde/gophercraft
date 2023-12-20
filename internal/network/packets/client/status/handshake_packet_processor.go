package status

import (
	"errors"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"log"
	"reflect"
)

func ProcessHandshakePacket(p common.Packet) error {
	packet, ok := p.(*HandshakePacket)
	if !ok {
		return errors.New("expected HandshakePacket, but got " + reflect.TypeOf(p).String())
	}

	log.Println(packet)

	return nil
}

package handshake

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

type PingResponsePacket struct {
	Payload int64
}

func (p *PingResponsePacket) PacketId() int {
	return 0x01
}

func (p *PingResponsePacket) Write(buffer network.Buffer) error {
	var err error

	err = buffer.WriteInt64(p.Payload)
	if err != nil {
		return err
	}

	return err
}

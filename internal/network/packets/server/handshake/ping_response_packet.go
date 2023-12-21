package handshake

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

type PingResponsePacket struct {
	Payload int64
}

func (p *PingResponsePacket) Write(buffer network.Buffer) error {
	var err error

	return err
}

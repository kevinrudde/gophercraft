package handshake

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

type StatusResponsePacket struct {
	Payload string
}

func (p *StatusResponsePacket) Write(buffer network.Buffer) error {
	var err error

	buffer.WriteString(p.Payload)
	
	return err
}

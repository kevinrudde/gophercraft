package status

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

type StatusRequestPacket struct {
}

func (p *StatusRequestPacket) From(buffer network.Buffer) error {
	return nil
}

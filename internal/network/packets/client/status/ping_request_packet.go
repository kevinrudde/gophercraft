package status

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

type PingRequestPacket struct {
	Payload int64
}

func (p *PingRequestPacket) From(buffer network.Buffer) error {
	var err error

	p.Payload, err = buffer.ReadInt64()
	if err != nil {
		return err
	}

	return err
}

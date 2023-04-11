package handshakepackets

import "github.com/kevinrudde/gophercraft/internal/network/packet"

type PingRequestPacket struct {
	Payload int64
}

func (p *PingRequestPacket) From(reader packet.Reader) error {
	var err error

	p.Payload, err = reader.Int64()
	if err != nil {
		return err
	}

	return err
}

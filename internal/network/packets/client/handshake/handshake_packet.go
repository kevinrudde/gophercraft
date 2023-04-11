package handshakepackets

import (
	"github.com/kevinrudde/gophercraft/internal/network/packet"
)

type HandshakePacket struct {
	ProtocolVersion int
	ServerAddress   string
	ServerPort      int16
	NextState       int
}

func (p *HandshakePacket) From(reader packet.Reader) error {
	var err error

	p.ProtocolVersion, err = reader.VarInt()
	if err != nil {
		return err
	}

	p.ServerAddress, err = reader.String()
	if err != nil {
		return err
	}

	p.ServerPort, err = reader.Int16()
	if err != nil {
		return err
	}

	p.NextState, err = reader.VarInt()
	if err != nil {
		return err
	}

	return err
}

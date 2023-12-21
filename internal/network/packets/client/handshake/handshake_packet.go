package handshake

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

type HandshakePacket struct {
	ProtocolVersion int
	ServerAddress   string
	ServerPort      int16
	NextState       int
}

func (p *HandshakePacket) From(buffer network.Buffer) error {
	var err error

	p.ProtocolVersion, err = buffer.ReadVarInt()
	if err != nil {
		return err
	}

	p.ServerAddress, err = buffer.ReadString()
	if err != nil {
		return err
	}

	p.ServerPort, err = buffer.ReadInt16()
	if err != nil {
		return err
	}

	p.NextState, err = buffer.ReadVarInt()
	if err != nil {
		return err
	}

	return err
}

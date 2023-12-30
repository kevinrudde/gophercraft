package login

import (
	"github.com/google/uuid"
	"github.com/kevinrudde/gophercraft/internal/network"
)

type LoginStartPacket struct {
	Name string
	UUID *uuid.UUID
}

func (p *LoginStartPacket) From(buffer network.Buffer) error {
	var err error

	p.Name, err = buffer.ReadString()
	if err != nil {
		return err
	}
	p.UUID, err = buffer.ReadUuid()
	if err != nil {
		return err
	}

	return nil
}

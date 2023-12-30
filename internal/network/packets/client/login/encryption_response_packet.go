package login

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

type EncryptionResponsePacket struct {
	SharedSecretLength int
	SharedSecret       []byte
	VerifyTokenLength  int
	VerifyToken        []byte
}

func (p *EncryptionResponsePacket) From(buffer network.Buffer) error {
	var err error

	p.SharedSecretLength, err = buffer.ReadVarInt()
	if err != nil {
		return err
	}
	p.SharedSecret, err = buffer.ReadBytes(p.SharedSecretLength)
	if err != nil {
		return err
	}
	p.VerifyTokenLength, err = buffer.ReadVarInt()
	if err != nil {
		return err
	}

	p.VerifyToken, err = buffer.ReadBytes(p.VerifyTokenLength)
	if err != nil {
		return err
	}

	return nil
}

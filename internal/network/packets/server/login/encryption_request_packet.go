package login

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

type EncryptionRequestPacket struct {
	ServerId          string
	PublicKeyLength   int
	PublicKey         []byte
	VerifyTokenLength int
	VerifyToken       []byte
}

func (p *EncryptionRequestPacket) PacketId() int {
	return 0x01
}

func (p *EncryptionRequestPacket) Write(buffer network.Buffer) error {
	buffer.WriteString(p.ServerId)
	buffer.WriteVarInt(p.PublicKeyLength)
	buffer.WriteBytes(p.PublicKey)
	buffer.WriteVarInt(p.VerifyTokenLength)
	buffer.WriteBytes(p.VerifyToken)

	return nil
}

package login

import (
	"github.com/google/uuid"
	"github.com/kevinrudde/gophercraft/internal/network"
)

// LoginSuccessPacket https://wiki.vg/Protocol#Login_Success
type LoginSuccessPacket struct {
	Uuid               *uuid.UUID
	Username           string
	NumberOfProperties int
}

func (p *LoginSuccessPacket) PacketId() int {
	return 0x02
}

func (p *LoginSuccessPacket) Write(buffer network.Buffer) error {
	buffer.WriteUuid(p.Uuid)
	buffer.WriteString(p.Username)
	buffer.WriteVarInt(p.NumberOfProperties)

	return nil
}

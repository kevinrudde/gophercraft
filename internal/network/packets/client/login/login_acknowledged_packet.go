package login

import "github.com/kevinrudde/gophercraft/internal/network"

type LoginAcknowledgedPacket struct {
}

func (s *LoginAcknowledgedPacket) From(buffer network.Buffer) error {
	return nil
}

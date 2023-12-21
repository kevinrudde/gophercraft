package common

import "github.com/kevinrudde/gophercraft/internal/network"

type ServerPacket interface {
	Write(buffer network.Buffer) error
}

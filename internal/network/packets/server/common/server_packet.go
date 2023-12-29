package common

import "github.com/kevinrudde/gophercraft/internal/network"

type ServerPacket interface {
	PacketId() int
	Write(buffer network.Buffer) error
}

package common

import "github.com/kevinrudde/gophercraft/internal/network"

type ClientPacket interface {
	From(buffer network.Buffer) error
}

package common

import "github.com/kevinrudde/gophercraft/internal/network"

type Packet interface {
	From(buffer network.Buffer) error
}

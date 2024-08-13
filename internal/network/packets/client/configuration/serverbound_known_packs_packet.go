package configuration

import (
	"github.com/kevinrudde/gophercraft/internal/network"
)

type ServerboundKnownPacksPacket struct {
	KnownPackCount int
	KnownPacks     []KnownPacks
}

type KnownPacks struct {
	Namespace string
	ID        string
	Version   string
}

func (p *ServerboundKnownPacksPacket) From(buffer network.Buffer) error {
	var err error

	p.KnownPackCount, err = buffer.ReadVarInt()
	if err != nil {
		return err
	}

	p.KnownPacks = make([]KnownPacks, p.KnownPackCount)

	for i := 0; i < p.KnownPackCount; i++ {
		p.KnownPacks[i].Namespace, err = buffer.ReadString()
		if err != nil {
			return err
		}

		p.KnownPacks[i].ID, err = buffer.ReadString()
		if err != nil {
			return err
		}

		p.KnownPacks[i].Version, err = buffer.ReadString()
		if err != nil {
			return err
		}
	}

	return err
}

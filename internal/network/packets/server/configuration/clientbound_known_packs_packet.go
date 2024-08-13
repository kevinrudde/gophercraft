package configuration

import "github.com/kevinrudde/gophercraft/internal/network"

// https://wiki.vg/Protocol#Clientbound_Known_Packs
type ClientboundKnownPacksPacket struct {
	KnownPacks []KnownPacks
}

type KnownPacks struct {
	Namespace string
	ID        string
	Version   string
}

func (p *ClientboundKnownPacksPacket) PacketId() int {
	return 0x0E
}

func (p *ClientboundKnownPacksPacket) Write(buffer network.Buffer) error {
	buffer.WriteVarInt(len(p.KnownPacks))

	for _, pack := range p.KnownPacks {
		buffer.WriteString(pack.Namespace)
		buffer.WriteString(pack.ID)
		buffer.WriteString(pack.Version)
	}

	return nil
}

package configuration

import "github.com/kevinrudde/gophercraft/internal/network"

// https://wiki.vg/Protocol#Registry_Data
type RegistryDataPacket struct {
	RegistryId string
	EntryCount int
}

type Entries struct {
	EntryId string
	HasData bool
	Data    string // TODO: Implement NBT
}

func (p *RegistryDataPacket) PacketId() int {
	return 0x0E
}

func (p *RegistryDataPacket) Write(buffer network.Buffer) error {
	return nil
}

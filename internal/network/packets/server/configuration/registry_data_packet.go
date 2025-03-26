package configuration

import (
	"bytes"

	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/pkg/nbt"
)

type RegistryDataPacket struct {
	RegistryId string
	Entries    map[string]*Entry
}

type Entry struct {
	Id   string
	Data nbt.Tag
}

func (p *RegistryDataPacket) PacketId() int {
	return 0x0E
}

func (p *RegistryDataPacket) Write(buffer network.Buffer) error {
	buffer.WriteString(p.RegistryId)
	buffer.WriteVarInt(len(p.Entries))

	// Write each registry
	for _, entry := range p.Entries {
		buffer.WriteString(entry.Id)

		if entry.Data != nil {
			buffer.WriteBool(true)
			err := buffer.WriteNBT(entry.Data)
			if err != nil {
				return err
			}
		} else {
			buffer.WriteBool(false)
		}
	}

	return nil
}

// NewRegistryDataPacket creates a new registry data packet
func NewRegistryDataPacket() *RegistryDataPacket {
	return &RegistryDataPacket{
		RegistryCount: 0,
		Registries:    make(map[string]*Registry),
	}
}

// AddRegistry adds a new registry to the packet
func (p *RegistryDataPacket) AddRegistry(registryId string) *Registry {
	registry := &Registry{
		RegistryId: registryId,
		Elements:   make(map[string]*RegistryElement),
	}
	p.Registries[registryId] = registry
	p.RegistryCount++
	return registry
}

// AddElement adds a new element to a registry
func (r *Registry) AddElement(elementId string, valueId int, data *nbt.CompoundTag) *RegistryElement {
	hasValue := data != nil
	element := &RegistryElement{
		Id:       valueId,
		Data:     data,
		HasValue: hasValue,
	}
	r.Elements[elementId] = element
	return element
}

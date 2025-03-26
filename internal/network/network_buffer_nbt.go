package network

import (
	"github.com/kevinrudde/gophercraft/pkg/nbt"
)

// ReadCompoundTag reads and returns a compound NBT tag from the buffer.
// This is a convenience method that returns nil if the tag is not a compound tag.
func (b *Buffer) ReadCompoundTag() (*nbt.CompoundTag, error) {
	tag, err := b.ReadNBT()
	if err != nil {
		return nil, err
	}

	if compoundTag, ok := tag.(*nbt.CompoundTag); ok {
		return compoundTag, nil
	}
	return nil, nil
}

// WriteCompoundTag writes a compound NBT tag to the buffer.
// This is a convenience method for writing compound tags which are commonly used in the protocol.
func (b *Buffer) WriteCompoundTag(tag *nbt.CompoundTag) error {
	return b.WriteNBT(tag)
}

// ReadListTag reads and returns a list NBT tag from the buffer.
// This is a convenience method that returns nil if the tag is not a list tag.
func (b *Buffer) ReadListTag() (*nbt.ListTag, error) {
	tag, err := b.ReadNBT()
	if err != nil {
		return nil, err
	}

	if listTag, ok := tag.(*nbt.ListTag); ok {
		return listTag, nil
	}
	return nil, nil
}

// WriteListTag writes a list NBT tag to the buffer.
func (b *Buffer) WriteListTag(tag *nbt.ListTag) error {
	return b.WriteNBT(tag)
}

// ReadEmptyNameCompound reads an NBT compound tag with an empty name.
// Minecraft often uses empty-named NBT tags in its protocol.
func (b *Buffer) ReadEmptyNameCompound() (*nbt.CompoundTag, error) {
	tag, err := b.ReadCompoundTag()
	if err != nil {
		return nil, err
	}

	// For protocol purposes, we often need to verify the tag has no name
	if tag.Name() != "" {
		return nil, nil
	}
	return tag, nil
}

// WriteEmptyNameCompound writes an NBT compound with an empty name.
// This is commonly used in the Minecraft protocol.
func (b *Buffer) WriteEmptyNameCompound(tag *nbt.CompoundTag) error {
	// Save the original name
	originalName := tag.Name()

	// Set empty name for protocol purposes
	tag.SetName("")

	// Write the tag
	err := b.WriteCompoundTag(tag)

	// Restore the original name
	tag.SetName(originalName)

	return err
}

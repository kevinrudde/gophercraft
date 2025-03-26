// Package nbt provides an implementation of Minecraft's Named Binary Tag format
package nbt

import (
	"errors"
	"fmt"
	"io"
)

// TagType represents the type of an NBT tag
type TagType byte

const (
	TagEnd       TagType = 0
	TagByte      TagType = 1
	TagShort     TagType = 2
	TagInt       TagType = 3
	TagLong      TagType = 4
	TagFloat     TagType = 5
	TagDouble    TagType = 6
	TagByteArray TagType = 7
	TagString    TagType = 8
	TagList      TagType = 9
	TagCompound  TagType = 10
	TagIntArray  TagType = 11
	TagLongArray TagType = 12
)

// Tag is the interface implemented by all NBT tag types
type Tag interface {
	Type() TagType
	Name() string
	SetName(name string)
	Read(r io.Reader) error
	Write(w io.Writer) error
	String() string
}

// BaseTag provides common fields and methods for NBT tags
type BaseTag struct {
	name string
}

// Name returns the tag's name
func (bt *BaseTag) Name() string {
	return bt.name
}

// SetName sets the tag's name
func (bt *BaseTag) SetName(name string) {
	bt.name = name
}

// ReadNamedTag reads a named NBT tag from the given reader
func ReadNamedTag(r io.Reader) (Tag, error) {
	var tagType TagType
	if err := readByte(r, (*byte)(&tagType)); err != nil {
		return nil, err
	}

	if tagType == TagEnd {
		return nil, errors.New("unexpected end tag")
	}

	// Read name
	name, err := readString(r)
	if err != nil {
		return nil, err
	}

	tag, err := newTag(tagType)
	if err != nil {
		return nil, err
	}

	tag.SetName(name)
	if err := tag.Read(r); err != nil {
		return nil, err
	}

	return tag, nil
}

// newTag creates a new tag of the specified type
func newTag(tagType TagType) (Tag, error) {
	switch tagType {
	case TagByte:
		return new(ByteTag), nil
	case TagShort:
		return new(ShortTag), nil
	case TagInt:
		return new(IntTag), nil
	case TagLong:
		return new(LongTag), nil
	case TagFloat:
		return new(FloatTag), nil
	case TagDouble:
		return new(DoubleTag), nil
	case TagByteArray:
		return new(ByteArrayTag), nil
	case TagString:
		return new(StringTag), nil
	case TagList:
		return new(ListTag), nil
	case TagCompound:
		return new(CompoundTag), nil
	case TagIntArray:
		return new(IntArrayTag), nil
	case TagLongArray:
		return new(LongArrayTag), nil
	default:
		return nil, fmt.Errorf("invalid tag type: %d", tagType)
	}
}

// WriteNamedTag writes a named NBT tag to the given writer
func WriteNamedTag(w io.Writer, tag Tag) error {
	// Write the tag type
	if err := writeByte(w, byte(tag.Type())); err != nil {
		return err
	}

	// Write the name
	if err := writeString(w, tag.Name()); err != nil {
		return err
	}

	// Write the tag payload
	if err := tag.Write(w); err != nil {
		return err
	}

	return nil
}

package nbt

import (
	"fmt"
	"io"
	"strings"
)

// ByteArrayTag represents a TAG_Byte_Array
type ByteArrayTag struct {
	BaseTag
	Value []byte
}

func (t *ByteArrayTag) Type() TagType {
	return TagByteArray
}

func (t *ByteArrayTag) Read(r io.Reader) error {
	length, err := readInt(r)
	if err != nil {
		return err
	}

	if length < 0 {
		return fmt.Errorf("negative length (%d) for byte array", length)
	}

	t.Value = make([]byte, length)
	_, err = io.ReadFull(r, t.Value)
	return err
}

func (t *ByteArrayTag) Write(w io.Writer) error {
	if err := writeInt(w, int32(len(t.Value))); err != nil {
		return err
	}
	_, err := w.Write(t.Value)
	return err
}

func (t *ByteArrayTag) String() string {
	return fmt.Sprintf("%s(%d): [%d bytes]", t.Name(), t.Type(), len(t.Value))
}

// IntArrayTag represents a TAG_Int_Array
type IntArrayTag struct {
	BaseTag
	Value []int32
}

func (t *IntArrayTag) Type() TagType {
	return TagIntArray
}

func (t *IntArrayTag) Read(r io.Reader) error {
	length, err := readInt(r)
	if err != nil {
		return err
	}

	if length < 0 {
		return fmt.Errorf("negative length (%d) for int array", length)
	}

	t.Value = make([]int32, length)
	for i := int32(0); i < length; i++ {
		t.Value[i], err = readInt(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *IntArrayTag) Write(w io.Writer) error {
	if err := writeInt(w, int32(len(t.Value))); err != nil {
		return err
	}
	for i := range t.Value {
		if err := writeInt(w, t.Value[i]); err != nil {
			return err
		}
	}
	return nil
}

func (t *IntArrayTag) String() string {
	return fmt.Sprintf("%s(%d): [%d ints]", t.Name(), t.Type(), len(t.Value))
}

// LongArrayTag represents a TAG_Long_Array
type LongArrayTag struct {
	BaseTag
	Value []int64
}

func (t *LongArrayTag) Type() TagType {
	return TagLongArray
}

func (t *LongArrayTag) Read(r io.Reader) error {
	length, err := readInt(r)
	if err != nil {
		return err
	}

	if length < 0 {
		return fmt.Errorf("negative length (%d) for long array", length)
	}

	t.Value = make([]int64, length)
	for i := int32(0); i < length; i++ {
		t.Value[i], err = readLong(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *LongArrayTag) Write(w io.Writer) error {
	if err := writeInt(w, int32(len(t.Value))); err != nil {
		return err
	}
	for i := range t.Value {
		if err := writeLong(w, t.Value[i]); err != nil {
			return err
		}
	}
	return nil
}

func (t *LongArrayTag) String() string {
	return fmt.Sprintf("%s(%d): [%d longs]", t.Name(), t.Type(), len(t.Value))
}

// CompoundTag represents a TAG_Compound
type CompoundTag struct {
	BaseTag
	Tags map[string]Tag
}

// NewCompoundTag creates a new compound tag with the given name
func NewCompoundTag(name string) *CompoundTag {
	tag := &CompoundTag{
		Tags: make(map[string]Tag),
	}
	tag.name = name
	return tag
}

func (t *CompoundTag) Type() TagType {
	return TagCompound
}

func (t *CompoundTag) Read(r io.Reader) error {
	t.Tags = make(map[string]Tag)

	for {
		var tagType TagType
		if err := readByte(r, (*byte)(&tagType)); err != nil {
			return err
		}

		if tagType == TagEnd {
			break
		}

		name, err := readString(r)
		if err != nil {
			return err
		}

		tag, err := newTag(tagType)
		if err != nil {
			return err
		}

		tag.SetName(name)
		if err := tag.Read(r); err != nil {
			return err
		}

		t.Tags[name] = tag
	}

	return nil
}

func (t *CompoundTag) Write(w io.Writer) error {
	for _, tag := range t.Tags {
		if err := writeByte(w, byte(tag.Type())); err != nil {
			return err
		}
		if err := writeString(w, tag.Name()); err != nil {
			return err
		}
		if err := tag.Write(w); err != nil {
			return err
		}
	}

	// Write end tag
	return writeByte(w, byte(TagEnd))
}

func (t *CompoundTag) GetTag(name string) Tag {
	return t.Tags[name]
}

func (t *CompoundTag) PutTag(tag Tag) {
	t.Tags[tag.Name()] = tag
}

func (t *CompoundTag) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s(%d): %d entries {\n", t.Name(), t.Type(), len(t.Tags)))
	for _, tag := range t.Tags {
		sb.WriteString("  " + tag.String() + "\n")
	}
	sb.WriteString("}")
	return sb.String()
}

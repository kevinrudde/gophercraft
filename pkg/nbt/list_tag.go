package nbt

import (
	"fmt"
	"io"
	"strings"
)

// ListTag represents a TAG_List
type ListTag struct {
	BaseTag
	ElementType TagType // Renamed from Type to ElementType to avoid conflict
	Value       []Tag
}

// NewListTag creates a new list tag with the given name and tag type
func NewListTag(name string, tagType TagType) *ListTag {
	tag := &ListTag{
		ElementType: tagType,
		Value:       make([]Tag, 0),
	}
	tag.name = name
	return tag
}

func (t *ListTag) TagType() TagType {
	return TagList
}

func (t *ListTag) Type() TagType {
	return TagList
}

func (t *ListTag) Read(r io.Reader) error {
	var err error
	var tagType TagType
	if err = readByte(r, (*byte)(&tagType)); err != nil {
		return err
	}
	t.ElementType = tagType

	var length int32
	if length, err = readInt(r); err != nil {
		return err
	}

	if length < 0 {
		return fmt.Errorf("negative length (%d) for list", length)
	}

	t.Value = make([]Tag, length)
	for i := int32(0); i < length; i++ {
		if t.ElementType == TagEnd && length > 0 {
			return fmt.Errorf("list with type TagEnd but non-zero length: %d", length)
		}

		tag, err := newTag(t.ElementType)
		if err != nil {
			return err
		}

		tag.SetName("") // List elements don't have names
		if err := tag.Read(r); err != nil {
			return err
		}

		t.Value[i] = tag
	}

	return nil
}

func (t *ListTag) Write(w io.Writer) error {
	if err := writeByte(w, byte(t.ElementType)); err != nil {
		return err
	}

	if err := writeInt(w, int32(len(t.Value))); err != nil {
		return err
	}

	for i := range t.Value {
		if err := t.Value[i].Write(w); err != nil {
			return err
		}
	}

	return nil
}

func (t *ListTag) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s(%d): %d entries of type %d {\n", t.Name(), t.Type(), len(t.Value), t.ElementType))
	for i, tag := range t.Value {
		sb.WriteString(fmt.Sprintf("  %d: %s\n", i, tag.String()))
	}
	sb.WriteString("}")
	return sb.String()
}

// Add adds a tag to the list
func (t *ListTag) Add(tag Tag) error {
	if tag.Type() != t.ElementType {
		return fmt.Errorf("cannot add tag of type %d to list of type %d", tag.Type(), t.ElementType)
	}

	t.Value = append(t.Value, tag)
	return nil
}

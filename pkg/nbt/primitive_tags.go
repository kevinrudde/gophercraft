package nbt

import (
	"fmt"
	"io"
	"strconv"
)

// ByteTag represents a TAG_Byte
type ByteTag struct {
	BaseTag
	Value byte
}

func (t *ByteTag) Type() TagType {
	return TagByte
}

func (t *ByteTag) Read(r io.Reader) error {
	return readByte(r, &t.Value)
}

func (t *ByteTag) Write(w io.Writer) error {
	return writeByte(w, t.Value)
}

func (t *ByteTag) String() string {
	return fmt.Sprintf("%s(%d): %d", t.Name(), t.Type(), t.Value)
}

// ShortTag represents a TAG_Short
type ShortTag struct {
	BaseTag
	Value int16
}

func (t *ShortTag) Type() TagType {
	return TagShort
}

func (t *ShortTag) Read(r io.Reader) error {
	var err error
	t.Value, err = readShort(r)
	return err
}

func (t *ShortTag) Write(w io.Writer) error {
	return writeShort(w, t.Value)
}

func (t *ShortTag) String() string {
	return fmt.Sprintf("%s(%d): %d", t.Name(), t.Type(), t.Value)
}

// IntTag represents a TAG_Int
type IntTag struct {
	BaseTag
	Value int32
}

func (t *IntTag) Type() TagType {
	return TagInt
}

func (t *IntTag) Read(r io.Reader) error {
	var err error
	t.Value, err = readInt(r)
	return err
}

func (t *IntTag) Write(w io.Writer) error {
	return writeInt(w, t.Value)
}

func (t *IntTag) String() string {
	return fmt.Sprintf("%s(%d): %d", t.Name(), t.Type(), t.Value)
}

// LongTag represents a TAG_Long
type LongTag struct {
	BaseTag
	Value int64
}

func (t *LongTag) Type() TagType {
	return TagLong
}

func (t *LongTag) Read(r io.Reader) error {
	var err error
	t.Value, err = readLong(r)
	return err
}

func (t *LongTag) Write(w io.Writer) error {
	return writeLong(w, t.Value)
}

func (t *LongTag) String() string {
	return fmt.Sprintf("%s(%d): %d", t.Name(), t.Type(), t.Value)
}

// FloatTag represents a TAG_Float
type FloatTag struct {
	BaseTag
	Value float32
}

func (t *FloatTag) Type() TagType {
	return TagFloat
}

func (t *FloatTag) Read(r io.Reader) error {
	var err error
	t.Value, err = readFloat(r)
	return err
}

func (t *FloatTag) Write(w io.Writer) error {
	return writeFloat(w, t.Value)
}

func (t *FloatTag) String() string {
	return fmt.Sprintf("%s(%d): %s", t.Name(), t.Type(), strconv.FormatFloat(float64(t.Value), 'f', -1, 32))
}

// DoubleTag represents a TAG_Double
type DoubleTag struct {
	BaseTag
	Value float64
}

func (t *DoubleTag) Type() TagType {
	return TagDouble
}

func (t *DoubleTag) Read(r io.Reader) error {
	var err error
	t.Value, err = readDouble(r)
	return err
}

func (t *DoubleTag) Write(w io.Writer) error {
	return writeDouble(w, t.Value)
}

func (t *DoubleTag) String() string {
	return fmt.Sprintf("%s(%d): %s", t.Name(), t.Type(), strconv.FormatFloat(t.Value, 'f', -1, 64))
}

// StringTag represents a TAG_String
type StringTag struct {
	BaseTag
	Value string
}

func (t *StringTag) Type() TagType {
	return TagString
}

func (t *StringTag) Read(r io.Reader) error {
	var err error
	t.Value, err = readString(r)
	return err
}

func (t *StringTag) Write(w io.Writer) error {
	return writeString(w, t.Value)
}

func (t *StringTag) String() string {
	return fmt.Sprintf("%s(%d): %s", t.Name(), t.Type(), t.Value)
}

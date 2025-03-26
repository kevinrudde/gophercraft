package nbt

import (
	"encoding/binary"
	"io"
	"math"
)

// Utility functions for reading data types

func readByte(r io.Reader, v *byte) error {
	var buf [1]byte
	if _, err := r.Read(buf[:]); err != nil {
		return err
	}
	*v = buf[0]
	return nil
}

func readShort(r io.Reader) (int16, error) {
	var buf [2]byte
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return 0, err
	}
	return int16(binary.BigEndian.Uint16(buf[:])), nil
}

func readInt(r io.Reader) (int32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return 0, err
	}
	return int32(binary.BigEndian.Uint32(buf[:])), nil
}

func readLong(r io.Reader) (int64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(buf[:])), nil
}

func readFloat(r io.Reader) (float32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return 0, err
	}
	return math.Float32frombits(binary.BigEndian.Uint32(buf[:])), nil
}

func readDouble(r io.Reader) (float64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.BigEndian.Uint64(buf[:])), nil
}

func readString(r io.Reader) (string, error) {
	length, err := readShort(r)
	if err != nil {
		return "", err
	}

	if length == 0 {
		return "", nil
	}

	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return "", err
	}

	return string(buf), nil
}

// Utility functions for writing data types

func writeByte(w io.Writer, v byte) error {
	_, err := w.Write([]byte{v})
	return err
}

func writeShort(w io.Writer, v int16) error {
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:], uint16(v))
	_, err := w.Write(buf[:])
	return err
}

func writeInt(w io.Writer, v int32) error {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], uint32(v))
	_, err := w.Write(buf[:])
	return err
}

func writeLong(w io.Writer, v int64) error {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], uint64(v))
	_, err := w.Write(buf[:])
	return err
}

func writeFloat(w io.Writer, v float32) error {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], math.Float32bits(v))
	_, err := w.Write(buf[:])
	return err
}

func writeDouble(w io.Writer, v float64) error {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(v))
	_, err := w.Write(buf[:])
	return err
}

func writeString(w io.Writer, v string) error {
	if err := writeShort(w, int16(len(v))); err != nil {
		return err
	}
	if len(v) > 0 {
		_, err := w.Write([]byte(v))
		return err
	}
	return nil
}

// NBT Helper functions

// NewByteTag creates a new byte tag with the given name and value
func NewByteTag(name string, value byte) *ByteTag {
	tag := &ByteTag{Value: value}
	tag.name = name
	return tag
}

// NewShortTag creates a new short tag with the given name and value
func NewShortTag(name string, value int16) *ShortTag {
	tag := &ShortTag{Value: value}
	tag.name = name
	return tag
}

// NewIntTag creates a new int tag with the given name and value
func NewIntTag(name string, value int32) *IntTag {
	tag := &IntTag{Value: value}
	tag.name = name
	return tag
}

// NewLongTag creates a new long tag with the given name and value
func NewLongTag(name string, value int64) *LongTag {
	tag := &LongTag{Value: value}
	tag.name = name
	return tag
}

// NewFloatTag creates a new float tag with the given name and value
func NewFloatTag(name string, value float32) *FloatTag {
	tag := &FloatTag{Value: value}
	tag.name = name
	return tag
}

// NewDoubleTag creates a new double tag with the given name and value
func NewDoubleTag(name string, value float64) *DoubleTag {
	tag := &DoubleTag{Value: value}
	tag.name = name
	return tag
}

// NewStringTag creates a new string tag with the given name and value
func NewStringTag(name string, value string) *StringTag {
	tag := &StringTag{Value: value}
	tag.name = name
	return tag
}

// NewByteArrayTag creates a new byte array tag with the given name and value
func NewByteArrayTag(name string, value []byte) *ByteArrayTag {
	tag := &ByteArrayTag{Value: value}
	tag.name = name
	return tag
}

// NewIntArrayTag creates a new int array tag with the given name and value
func NewIntArrayTag(name string, value []int32) *IntArrayTag {
	tag := &IntArrayTag{Value: value}
	tag.name = name
	return tag
}

// NewLongArrayTag creates a new long array tag with the given name and value
func NewLongArrayTag(name string, value []int64) *LongArrayTag {
	tag := &LongArrayTag{Value: value}
	tag.name = name
	return tag
}

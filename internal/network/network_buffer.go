package network

import (
	"encoding/binary"
	"errors"
)

type Buffer struct {
	buf        []byte
	ReadIndex  int
	WriteIndex int
}

const defaultCapacity = 1024

func CreateBufferWithInitialCapacity(initialCapacity int) Buffer {
	return Buffer{
		buf: make([]byte, initialCapacity),
	}
}

func CreateBufferWithBuf(buf []byte) Buffer {
	return Buffer{
		buf: buf,
	}
}

func CreateBuffer() Buffer {
	return Buffer{
		buf: make([]byte, defaultCapacity),
	}
}

func (b *Buffer) ReadBool() (bool, error) {
	data, err := b.readBytes(1)
	if err != nil {
		return false, err
	}

	if data[0] == 0x01 {
		return true, nil
	}
	return false, nil
}

func (b *Buffer) ReadByte() (byte, error) {
	data, err := b.readBytes(1)
	if err != nil {
		return 0, err
	}

	return data[0], nil
}

func (b *Buffer) ReadInt16() (int16, error) {
	data, err := b.readBytes(4)
	if err != nil {
		return 0, err
	}
	value := binary.BigEndian.Uint16(data)

	return int16(value), nil
}

func (b *Buffer) ReadUInt16() (uint16, error) {
	data, err := b.readBytes(4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(data), nil
}

func (b *Buffer) ReadInt32() (int32, error) {
	data, err := b.readBytes(4)
	if err != nil {
		return 0, err
	}
	value := binary.BigEndian.Uint32(data)

	return int32(value), nil
}

func (b *Buffer) ReadInt64() (int64, error) {
	data, err := b.readBytes(4)
	if err != nil {
		return 0, err
	}
	value := binary.BigEndian.Uint64(data)

	return int64(value), nil
}

func (b *Buffer) String() (string, error) {
	length, err := b.ReadVarInt()
	if err != nil {
		return "", err
	}
	data, err := b.readBytes(length)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (b *Buffer) ReadVarInt() (int, error) {
	var value int = 0
	var position int = 0
	var currentByte byte

	for {
		data, err := b.readBytes(1)
		if err != nil {
			return 0, err
		}
		currentByte = data[0]

		value |= int(currentByte&0x7F) << position

		if (currentByte & 0x80) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return 0, errors.New("VarInt is bigger than 32")
		}
	}
	return value, nil
}

func (b *Buffer) readBytes(count int) ([]byte, error) {
	if len(b.buf) >= b.ReadIndex+count {
		return nil, errors.New("buffer read is out of bounds")
	}
	b.ReadIndex += count
	return b.buf[b.ReadIndex : b.ReadIndex+count], nil
}

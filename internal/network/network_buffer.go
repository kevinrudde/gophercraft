package network

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type Buffer struct {
	buf        []byte
	ReadIndex  int
	WriteIndex int
}

const (
	SegmentBits = 0x7F
	ContinueBit = 0x80
)

func CreateBufferWithBuf(buf []byte) Buffer {
	return Buffer{
		buf:        buf,
		ReadIndex:  0,
		WriteIndex: 0,
	}
}

func CreateBuffer() Buffer {
	return Buffer{
		buf:        make([]byte, 0),
		ReadIndex:  0,
		WriteIndex: 0,
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

func (b *Buffer) WriteBool(value bool) {
	if value {
		b.writeByte(0x01)
	} else {
		b.writeByte(0x00)
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	data, err := b.readBytes(1)
	if err != nil {
		return 0, err
	}

	return data[0], nil
}

func (b *Buffer) WriteByte(value byte) {
	b.writeByte(value)
}

func (b *Buffer) ReadInt16() (int16, error) {
	data, err := b.readBytes(4)
	if err != nil {
		return 0, err
	}
	value := binary.BigEndian.Uint16(data)

	return int16(value), nil
}

func (b *Buffer) WriteInt16(value int16) {
	buf := make([]byte, 1)
	binary.BigEndian.PutUint16(buf, uint16(value))
	fmt.Println(buf)
	b.writeBytes(buf)

	//b.buf = binary.BigEndian.AppendUint16(b.buf, uint16(value))
}

func (b *Buffer) ReadUInt16() (uint16, error) {
	data, err := b.readBytes(4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(data), nil
}

func (b *Buffer) WriteUInt16(value uint16) {
	b.buf = binary.BigEndian.AppendUint16(b.buf, value)
}

func (b *Buffer) ReadInt32() (int32, error) {
	data, err := b.readBytes(4)
	if err != nil {
		return 0, err
	}
	value := binary.BigEndian.Uint32(data)

	return int32(value), nil
}

func (b *Buffer) WriteInt32(value uint32) {
	b.buf = binary.BigEndian.AppendUint32(b.buf, value)
}

func (b *Buffer) ReadInt64() (int64, error) {
	data, err := b.readBytes(4)
	if err != nil {
		return 0, err
	}
	value := binary.BigEndian.Uint64(data)

	return int64(value), nil
}

func (b *Buffer) WriteUInt32(value uint32) {
	b.buf = binary.BigEndian.AppendUint32(b.buf, value)
}

func (b *Buffer) ReadString() (string, error) {
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

func (b *Buffer) WriteString(value string) {
	length := len(value)
	b.WriteVarInt(length)
	b.writeBytes([]byte(value))
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

		value |= int(currentByte&SegmentBits) << position

		if (currentByte & ContinueBit) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return 0, errors.New("VarInt is bigger than 32")
		}
	}
	return value, nil
}

func (b *Buffer) WriteVarInt(value int) {
	for {
		if value & ^SegmentBits == 0 {
			b.writeByte(byte(value))
			return
		}
		b.writeByte(byte(value&SegmentBits | ContinueBit))
		value >>= 7
	}
}

func (b *Buffer) readBytes(count int) ([]byte, error) {
	if len(b.buf) > b.ReadIndex+count {
		return nil, errors.New("buffer read is out of bounds")
	}
	b.ReadIndex += count
	return b.buf[b.ReadIndex-1 : b.ReadIndex+count-1], nil
}

func (b *Buffer) writeByte(data byte) {
	b.buf = append(b.buf, data)
	b.WriteIndex += 1
}

func (b *Buffer) writeBytes(data []byte) {
	b.buf = append(b.buf, data...)
	b.WriteIndex += len(data)
}

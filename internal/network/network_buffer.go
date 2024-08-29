package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/google/uuid"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() any {
		return createBuffer()
	},
}

type Buffer struct {
	buf *bytes.Buffer
}

const (
	SegmentBits = 0x7F
	ContinueBit = 0x80
)

func GetBufferFromPool() Buffer {
	return bufferPool.Get().(Buffer)
}

func GetBufferFromPoolWithBuf(buf []byte) Buffer {
	buffer := bufferPool.Get().(Buffer)
	buffer.WriteBytes(buf)
	return buffer
}

func PutBufferToPool(b Buffer) {
	b.Reset()
	bufferPool.Put(b)
}

func createBuffer() Buffer {
	return Buffer{
		buf: new(bytes.Buffer),
	}
}

func (b *Buffer) Length() int {
	return b.buf.Len()
}

func (b *Buffer) Bytes() []byte {
	return b.buf.Bytes()
}

func (b *Buffer) ReadBool() (bool, error) {
	var data [1]byte
	_, err := b.buf.Read(data[:1])
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
		b.buf.WriteByte(0x01)
	} else {
		b.buf.WriteByte(0x00)
	}
}

func (b *Buffer) ReadBytes(length int) ([]byte, error) {
	data := make([]byte, length)
	_, err := b.buf.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (b *Buffer) WriteBytes(value []byte) {
	b.buf.Write(value)
}

func (b *Buffer) ReadByte() (byte, error) {
	var data [1]byte
	_, err := b.buf.Read(data[:1])
	if err != nil {
		return 0, err
	}

	return data[0], nil
}

func (b *Buffer) WriteByte(value byte) error {
	return b.buf.WriteByte(value)
}

func (b *Buffer) ReadInt16() (int16, error) {
	var data [2]byte
	_, err := b.buf.Read(data[:2])
	if err != nil {
		return 0, err
	}
	value := binary.BigEndian.Uint16(data[:2])

	return int16(value), nil
}

func (b *Buffer) WriteInt16(value int16) {
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:2], uint16(value))
	b.buf.Write(buf[:2])
}

func (b *Buffer) ReadUInt16() (uint16, error) {
	var data [2]byte
	_, err := b.buf.Read(data[:2])
	if err != nil {
		return 0, err
	}
	value := binary.BigEndian.Uint16(data[:2])

	return value, nil
}

func (b *Buffer) WriteUInt16(value uint16) error {
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:2], value)
	_, err := b.buf.Write(buf[:2])
	return err
}

func (b *Buffer) ReadInt32() (int32, error) {
	var data [4]byte
	_, err := b.buf.Read(data[:4])
	if err != nil {
		return 0, err
	}
	value := binary.BigEndian.Uint32(data[:4])

	return int32(value), nil
}

func (b *Buffer) WriteInt32(value int32) error {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], uint32(value))
	_, err := b.buf.Write(buf[:4])
	return err
}

func (b *Buffer) WriteInt64(value int64) error {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:8], uint64(value))
	_, err := b.buf.Write(buf[:8])
	return err
}

func (b *Buffer) ReadInt64() (int64, error) {
	var data [8]byte
	_, err := b.buf.Read(data[:8])
	if err != nil {
		return 0, err
	}
	value := binary.BigEndian.Uint64(data[:8])

	return int64(value), nil
}

func (b *Buffer) ReadUuid() (*uuid.UUID, error) {
	var data [16]byte
	_, err := b.buf.Read(data[:16])
	if err != nil {
		return nil, err
	}
	readUuid, err := uuid.FromBytes(data[:16])
	return &readUuid, err
}

func (b *Buffer) WriteUuid(uuid *uuid.UUID) {
	b.buf.Write(uuid[:])
}

func (b *Buffer) ReadString() (string, error) {
	length, err := b.ReadVarInt()
	if err != nil {
		return "", err
	}
	data := make([]byte, length)
	_, err = b.buf.Read(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (b *Buffer) WriteString(value string) {
	length := len(value)
	b.WriteVarInt(length)
	b.buf.Write([]byte(value))
}

func (b *Buffer) ReadVarInt() (int, error) {
	var value int = 0
	var position int = 0

	for {
		currentByte, err := b.buf.ReadByte()
		if err != nil {
			return 0, err
		}

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

func (b *Buffer) WriteVarInt(value int) int {
	previousLen := b.buf.Len()
	for {
		if value & ^SegmentBits == 0 {
			b.buf.WriteByte(byte(value))
			return b.buf.Len() - previousLen
		}
		b.buf.WriteByte(byte(value&SegmentBits | ContinueBit))
		value >>= 7
	}
}

func (b *Buffer) WriteBuf(buf *Buffer) {
	b.buf.Write(buf.Bytes())
}

func (b *Buffer) WriteStringSlice(slice []string) error {
	if len(slice) == 0 {
		return b.WriteByte(0)
	}

	b.WriteVarInt(len(slice))
	for _, value := range slice {
		b.WriteString(value)
	}
	return nil
}

func (b *Buffer) Reset() {
	b.buf.Reset()
}

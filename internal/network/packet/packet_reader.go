package packet

import (
	"encoding/binary"
	"errors"
)

type Reader struct {
	RawPacket *RawPacket
	ReadIndex int
}

func NewReader(rawPacket *RawPacket) Reader {
	return Reader{
		RawPacket: rawPacket,
		ReadIndex: 0,
	}
}

func (r *Reader) VarInt() (int, error) {
	var value int = 0
	var position int = 0
	var currentByte byte

	for {
		currentByte = r.RawPacket.Data[r.ReadIndex]
		r.ReadIndex++
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

func (r *Reader) String() (string, error) {
	length, err := r.VarInt()
	if err != nil {
		return "", err
	}
	stringData := r.RawPacket.Data[r.ReadIndex : r.ReadIndex+length]
	r.ReadIndex += length
	return string(stringData), nil
}

func (r *Reader) Int() (int, error) {
	buf := r.RawPacket.Data[r.ReadIndex : r.ReadIndex+4]
	value := binary.BigEndian.Uint32(buf)
	r.ReadIndex += 4

	return int(value), nil
}

func (r *Reader) Int16() (int16, error) {
	buf := r.RawPacket.Data[r.ReadIndex : r.ReadIndex+4]
	value := binary.BigEndian.Uint16(buf)
	r.ReadIndex += 2

	return int16(value), nil
}

func (r *Reader) Int32() (int32, error) {
	buf := r.RawPacket.Data[r.ReadIndex : r.ReadIndex+4]
	value := binary.BigEndian.Uint32(buf)
	r.ReadIndex += 2

	return int32(value), nil
}

func (r *Reader) Int64() (int64, error) {
	buf := r.RawPacket.Data[r.ReadIndex : r.ReadIndex+4]
	value := binary.BigEndian.Uint64(buf)
	r.ReadIndex += 2

	return int64(value), nil
}

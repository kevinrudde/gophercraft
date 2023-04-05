package util

import (
	"errors"
)

func VarInt(b []byte) (int64, int) {
	var value int64 = 0
	var bitOffset byte = 0
	var currIndx = 0
	var currentByte byte

	for {
		if bitOffset == 35 {
			return 0, 0
		}

		currentByte = b[currIndx]
		value |= int64(currentByte&0x7F) << uint(bitOffset)

		currIndx++
		bitOffset += 7

		if currentByte&0x80 != 0x80 {
			break
		}
	}

	return int64(value), currIndx
}

func ReadVarInt(buf []byte) (int32, int, error) {
	var value int32 = 0
	var position int = 0
	var currentByte byte
	index := 0

	for {
		currentByte = buf[index]
		index++
		value |= int32(currentByte&0x7F) << position

		if (currentByte & 0x80) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return 0, index, errors.New("VarInt is bigger than 32")
		}
	}
	return value, index, nil
}

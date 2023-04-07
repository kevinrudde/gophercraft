package util

import (
	"errors"
)

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

package util

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

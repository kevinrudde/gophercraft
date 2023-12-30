package encryption

import (
	"crypto/cipher"
)

type Connection struct {
	Decrypter cipher.Stream
	Encrypter cipher.Stream
}

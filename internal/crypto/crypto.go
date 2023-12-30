package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

var (
	PrivateKey   *rsa.PrivateKey
	PublicKeyLen int
	PublicKey    []byte
)

func Init() {
	var err error

	PrivateKey, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	PrivateKey.Precompute()

	PublicKey, err = x509.MarshalPKIXPublicKey(PrivateKey.Public())
	if err != nil {
		panic(err)
	}
	PublicKeyLen = len(PublicKey)
}

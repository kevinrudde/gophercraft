package login

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/subtle"
	"errors"
	"github.com/kevinrudde/gophercraft/internal/crypto"
	"github.com/kevinrudde/gophercraft/internal/network/encryption"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client/common"
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/login"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"reflect"
)

func ProcessEncryptionResponsePacket(connection *networkplayer.PlayerConnection, p common.ClientPacket) error {
	packet, ok := p.(*EncryptionResponsePacket)
	if !ok {
		return errors.New("expected EncryptionResponsePacket, but got " + reflect.TypeOf(p).String())
	}

	decryptedVerifyToken, err := rsa.DecryptPKCS1v15(rand.Reader, crypto.PrivateKey, packet.VerifyToken)
	if err != nil {
		return err
	}

	if subtle.ConstantTimeCompare(decryptedVerifyToken, connection.EncryptionDetails.VerifyToken) != 1 {
		return errors.New("verify token mismatch")
	}

	decryptedSharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, crypto.PrivateKey, packet.SharedSecret)
	if err != nil {
		return err
	}

	decrypter, encrypter, err := createEncryptedConnection(decryptedSharedSecret)
	if err != nil {
		return err
	}

	connection.EncryptionDetails.SharedSecret = packet.SharedSecret
	connection.EncryptedConnection = &encryption.Connection{
		Decrypter: decrypter,
		Encrypter: encrypter,
	}

	response := &login.LoginSuccessPacket{
		Uuid:               connection.Uuid,
		Username:           connection.Username,
		NumberOfProperties: 0,
	}

	return connection.SendPacket(response)
}

func createEncryptedConnection(sharedSecret []byte) (cipher.Stream, cipher.Stream, error) {
	cipher, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return nil, nil, err
	}

	decrypter := crypto.NewCFB8Decrypter(cipher, sharedSecret)
	encrypter := crypto.NewCFB8Encrypter(cipher, sharedSecret)

	return decrypter, encrypter, nil
}

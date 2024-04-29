package login

import (
	"crypto/rand"
	"github.com/kevinrudde/gophercraft/internal/crypto"
	"github.com/kevinrudde/gophercraft/internal/network/encryption"
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/login"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"log"
)

func ProcessLoginStartPacket(connection *networkplayer.PlayerConnection, packet *LoginStartPacket) error {
	log.Printf("Got login start from %s (%s)\n", packet.Name, packet.UUID)

	verifyToken := make([]byte, 4)
	_, err := rand.Read(verifyToken)
	if err != nil {
		return err
	}

	connection.EncryptionDetails = &encryption.Details{
		VerifyToken: verifyToken,
	}

	connection.Uuid = packet.UUID
	connection.Username = packet.Name

	response := &login.EncryptionRequestPacket{
		ServerId:          "",
		PublicKeyLength:   crypto.PublicKeyLen,
		PublicKey:         crypto.PublicKey,
		VerifyTokenLength: len(verifyToken),
		VerifyToken:       verifyToken,
	}

	return connection.SendPacket(response)
}

package networkplayer

import (
	"github.com/google/uuid"
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/encryption"
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/common"
	"net"
)

var PlayerConnections = make(map[*PlayerConnection]struct{})

type PlayerConnection struct {
	Conn                net.Conn
	ConnectionState     network.ConnectionState
	EncryptionDetails   *encryption.Details
	EncryptedConnection *encryption.Connection

	Uuid     *uuid.UUID
	Username string
}

func (s *PlayerConnection) SendPacket(p common.ServerPacket) error {
	buf := network.GetBufferFromPool()
	defer network.PutBufferToPool(buf)
	dataBuf := network.GetBufferFromPool()
	defer network.PutBufferToPool(dataBuf)
	packetIdBuf := network.GetBufferFromPool()
	defer network.PutBufferToPool(packetIdBuf)

	err := p.Write(dataBuf)
	if err != nil {
		return err
	}

	// TODO: optimize this to not create a new buffer for the packetId
	packetId := p.PacketId()
	packetIdLen := packetIdBuf.WriteVarInt(packetId)

	length := dataBuf.Length() + packetIdLen

	buf.WriteVarInt(length)
	buf.WriteBuf(&packetIdBuf)
	buf.WriteBuf(&dataBuf)

	if s.EncryptedConnection != nil {
		s.EncryptedConnection.Encrypter.XORKeyStream(buf.Bytes(), buf.Bytes())
	}

	_, err = s.Conn.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (s *PlayerConnection) Close() error {
	return s.Conn.Close()
}

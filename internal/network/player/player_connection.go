package networkplayer

import (
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/server/common"
	"net"
)

var PlayerConnections = make(map[*PlayerConnection]struct{})

type PlayerConnection struct {
	Conn            net.Conn
	ConnectionState network.ConnectionState
}

func (s *PlayerConnection) SendPacket(p common.ServerPacket) error {
	buf := network.CreateBuffer()

	err := p.Write(buf)
	if err != nil {
		return err
	}

	_, err = s.Conn.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

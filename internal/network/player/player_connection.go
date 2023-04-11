package networkplayer

import (
	"github.com/kevinrudde/gophercraft/internal/network"
	"net"
)

var PlayerConnections = make(map[PlayerConnection]struct{})

type PlayerConnection struct {
	Conn            net.Conn
	ConnectionState network.ConnectionState
}

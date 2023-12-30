package minecraft

import (
	"github.com/kevinrudde/gophercraft/internal/crypto"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client"
	"github.com/kevinrudde/gophercraft/internal/network/server"
)

// Server represents the Minecraft server.
// It holds every component needed to run the server.
type Server struct {
	networkServer *server.Server
}

func newServer() *Server {
	return &Server{}
}

// Instance is the singleton instance of the Minecraft server.
var Instance *Server = newServer()

// Init initializes the Minecraft server.
// This function should be called before Start.
func (s *Server) Init() {
	crypto.Init()
	s.networkServer = server.NewServer()

	client.InitializeClientPacketProcessors()
}

// Start starts the Minecraft server on the given address.
// The address should be in the format of "host:port", e.g. ":25565".
// This function will block until the server is stopped.
func (s *Server) Start(listenAddr string) error {
	return s.networkServer.Start(listenAddr)
}

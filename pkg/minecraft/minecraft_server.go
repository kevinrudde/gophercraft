package minecraft

import (
	"github.com/kevinrudde/gophercraft/internal/crypto"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client"
	"github.com/kevinrudde/gophercraft/internal/network/server"
)

type Server struct {
	networkServer *server.Server
}

var Instance *Server = New()

func New() *Server {
	return &Server{}
}

func (s *Server) Init() {
	crypto.Init()
	s.networkServer = server.NewServer()

	client.InitializeClientPacketProcessors()
}

func (s *Server) Start(listenAddr string) error {
	return s.networkServer.Start(listenAddr)
}

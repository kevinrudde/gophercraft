package main

import (
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packet"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	listener   net.Listener
	quitCh     chan struct{}
	msgCh      chan *packet.RawPacket
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitCh:     make(chan struct{}),
		msgCh:      make(chan *packet.RawPacket, 10),
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()
	s.listener = listener

	go s.acceptLoop()

	<-s.quitCh
	close(s.msgCh)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("New connection from:", conn.RemoteAddr().String())
		playerConnection := networkplayer.PlayerConnection{
			Conn:            conn,
			ConnectionState: network.Status,
		}
		networkplayer.PlayerConnections[playerConnection] = struct{}{}

		go s.readLoop(conn, playerConnection)
	}
}

func (s *Server) readLoop(conn net.Conn, connection networkplayer.PlayerConnection) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error:", err)
			break
		}

		msg := buf[:length]

		buffer := network.CreateBufferWithBuf(msg)
		packetLength, err := buffer.ReadVarInt()
		if err != nil {
			continue
		}
		packetId, err := buffer.ReadVarInt()
		if err != nil {
			continue
		}

		if packetLength < 0 {
			continue
		}

		err = client.ProcessPacket(connection, packetId, buffer.Bytes())
		if err != nil {
			return
		}
	}
}

func main() {
	server := NewServer(":25565")
	fmt.Println("Listening on :25565")

	log.Fatal(server.Start())
}

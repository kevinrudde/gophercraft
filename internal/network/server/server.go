package server

import (
	"errors"
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network"
	"github.com/kevinrudde/gophercraft/internal/network/packets/client"
	networkplayer "github.com/kevinrudde/gophercraft/internal/network/player"
	"io"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	listener   net.Listener
	quitCh     chan struct{}
}

func NewServer() *Server {
	return &Server{
		quitCh: make(chan struct{}),
	}
}

func (s *Server) Start(listenAddr string) error {
	s.listenAddr = listenAddr
	listener, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()
	s.listener = listener

	go s.acceptLoop()

	<-s.quitCh

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		playerConnection := &networkplayer.PlayerConnection{
			Conn:            conn,
			ConnectionState: network.Handshake,
		}
		networkplayer.PlayerConnections[playerConnection] = struct{}{}

		go s.readLoop(conn, playerConnection)
	}
}

func (s *Server) readLoop(conn net.Conn, connection *networkplayer.PlayerConnection) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed) {
				break
			}

			fmt.Println("got an error in readLoop:", err)
			break
		}

		msg := buf[:length]

		if connection.EncryptedConnection != nil {
			connection.EncryptedConnection.Decrypter.XORKeyStream(msg, msg)
		}

		buffer := network.GetBufferFromPoolWithBuf(msg)
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
			log.Println("got an error while processing packet", err)
			return
		}
	}
}

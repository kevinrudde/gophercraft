package main

import (
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/network/packet"
	packets "github.com/kevinrudde/gophercraft/internal/network/packets/client/handshake"
	"github.com/kevinrudde/gophercraft/internal/util"
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

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error:", err)
			break
		}

		msg := buf[:length]

		packetLength, n, err := util.ReadVarInt(msg)
		packetId, m, err := util.ReadVarInt(msg[n:])

		if int(packetLength)-m <= 0 {
			continue
		}

		s.msgCh <- &packet.RawPacket{
			DataLength: length - (n + m),
			Length:     length,
			PacketId:   packetId,
			Data:       msg[n+m:],
		}
	}
}

func (s *Server) ProcessRawPackets() {
	for rawPacket := range s.msgCh {
		reader := packet.NewReader(rawPacket)

		// TODO: make this more magic
		var receivedPacket packets.HandshakePacket

		switch rawPacket.PacketId {
		case 0:
			receivedPacket = packets.HandshakePacket{}
			err := receivedPacket.From(reader)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(receivedPacket)

			break
		}
	}
}

func main() {
	server := NewServer(":25565")
	go server.ProcessRawPackets()
	fmt.Println("Listening on :25565")

	log.Fatal(server.Start())
}

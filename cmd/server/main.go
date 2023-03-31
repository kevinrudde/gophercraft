package main

import (
	"fmt"
	"github.com/kevinrudde/gophercraft/internal/util"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	listener   net.Listener
	quitCh     chan struct{}
	msgCh      chan []byte
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitCh:     make(chan struct{}),
		msgCh:      make(chan []byte, 10),
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
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}

		s.msgCh <- buf[:n]
	}
}

func main() {
	server := NewServer(":25565")
	fmt.Println("Listening on :25565")

	go func() {
		for msg := range server.msgCh {
			length, n := util.VarInt(msg)
			packetId, n := util.VarInt(msg[:n])
			fmt.Println("packet length:", length)
			fmt.Println("packet id:", packetId)
		}
	}()

	log.Fatal(server.Start())
}

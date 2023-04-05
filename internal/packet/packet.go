package packet

import "fmt"

type RawPacket struct {
	Length   int
	PacketId int32
	Data     []byte
}

type HandshakePacket struct {
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint32
	NextState       int32
}

func (p *HandshakePacket) FromRawPacket(rawPacket *RawPacket) {

}

func ReadPacket(rawPacket *RawPacket) {
	switch rawPacket.PacketId {
	case 0x00:
		fmt.Println("handshake")
		packet := HandshakePacket{}
		packet.FromRawPacket(rawPacket)
		fmt.Println(packet)
		break
	}
}

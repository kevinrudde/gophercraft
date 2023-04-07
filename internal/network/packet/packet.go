package packet

type RawPacket struct {
	DataLength int
	Length     int
	PacketId   int32
	Data       []byte
}

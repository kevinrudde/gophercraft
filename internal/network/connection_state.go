package network

type ConnectionState int

const (
	Handshake ConnectionState = iota
	Status
	Login
	Configuration
	Play
)

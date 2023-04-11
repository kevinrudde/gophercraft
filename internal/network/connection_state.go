package network

type ConnectionState int

const (
	Unknown ConnectionState = 0
	Status  ConnectionState = 1
	Login   ConnectionState = 2
	Play    ConnectionState = 3
)

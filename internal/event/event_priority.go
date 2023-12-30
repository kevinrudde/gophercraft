package event

type Priority int

const (
	Monitor Priority = -1
	Lowest  Priority = 0
	Low     Priority = 1
	Normal  Priority = 2
	High    Priority = 3
	Highest Priority = 4
)

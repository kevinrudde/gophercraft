package event

type Priority int

const (
	Monitor Priority = 0
	Lowest  Priority = 1
	Low     Priority = 2
	Normal  Priority = 3
	High    Priority = 4
	Highest Priority = 5
)

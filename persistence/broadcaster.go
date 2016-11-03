package persistence

type broadcaster interface {
	Broadcast(data []byte)
}
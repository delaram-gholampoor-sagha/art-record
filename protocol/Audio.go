package protocol

type Audio struct {
	ID         uint64
	QuiddityID [16]byte
	ObjectID   [16]byte
	Status     uint8
}

type AudioServices interface {
	
}

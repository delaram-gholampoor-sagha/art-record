package protocol

type Video struct {
	ID         uint64
	QuiddityID [16]byte
	ObjectID   [16]byte
}

type VideoServices interface {
}

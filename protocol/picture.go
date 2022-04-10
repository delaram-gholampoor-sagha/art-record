package protocol

type Picture struct {
	ID        [16]byte
	RelatedID [16]byte
	ObjectID  [16]byte
}

type PictureServices interface {
	// rating ?
	FindPictureByQuiddityID(quiddityID uint64, offset uint64, limit uint64) (id []uint64)
}

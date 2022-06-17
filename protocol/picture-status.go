package protocol

type PictureStatus interface {
	ObjectID() [16]byte     // object domain
	Status() Picture_Status //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type PictureStatus_StorageServices interface {
	Save(ps PictureStatus) (err protocol.Error)

	Count(objectID [16]byte) (numbers uint64, err protocol.Error)
	Get(objectID [16]byte, versionOffset uint64) (ps PictureStatus, err protocol.Error)
	Last(objectID [16]byte) (ps PictureStatus, numbers uint64, err protocol.Error)

	// FilterByStatus(status Picture_Status, offset, limit uint64) (objectIDs [][16]byte, numbers uint64, err protocol.Error)
	// protocol.EventTarget
}

type Picture_Status Quiddity_Status

const (
// Picture_Status_ Picture_Status = (Quiddity_Status_FreeFlag << iota)
)

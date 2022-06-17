package protocol

type CategoryStatus interface {
	CategoryID() [16]byte    // category domain
	Status() Category_Status //
	Time() protocol.Time     // save time
	RequestID() [16]byte     // user-request domain
}

type CategoryStatus_StorageServices interface {
	Save(gs CategoryStatus) protocol.Error

	Count(categoryID [16]byte) (numbers uint64, err protocol.Error)
	Get(categoryID [16]byte, versionOffset uint64) (gs CategoryStatus, err protocol.Error)
	Last(categoryID [16]byte) (gs CategoryStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status Category_Status, offset, limit uint64) (categoryIDs [][16]byte, err protocol.Error)
	// protocol.EventTarget
}

type Category_Status Quiddity_Status

const (
// Category_Status_ = Category_Status(Quiddity_Status_FreeFlag << iota)
)

package protocol

type ContentStatus interface {
	ContentID() [16]byte    // content domain
	Status() Content_Status //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ContentStatus_StorageServices interface {
	Save(cs ContentStatus) protocol.Error

	Count(contentID [16]byte) (numbers uint64, err protocol.Error)
	Get(contentID [16]byte, versionOffset uint64) (cs ContentStatus, err protocol.Error)
	Last(contentID [16]byte) (cs ContentStatus, numbers uint64, err protocol.Error)

	// FilterByStatus(status Content_Status, offset, limit uint64) (contentIDs [][16]byte, numbers uint64, err protocol.Error)
	// protocol.EventTarget
}

type Content_Status Quiddity_Status

const (
// Content_Status_ = Content_Status(Quiddity_Status_FreeFlag << iota)
)

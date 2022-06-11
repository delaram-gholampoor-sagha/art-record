package protocol


type ContentStatus interface {
	ContentID() [16]byte    // content domain
	Status() Content_Status //
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

type ContentStatus_StorageServices interface {
	Save(cs ContentStatus) protocol.Error

	Count(contentID [16]byte) (numbers uint64, err protocol.Error)
	Get(contentID [16]byte, versionOffset uint64) (cs ContentStatus, err protocol.Error)
	Last(contentID [16]byte) (cs ContentStatus, numbers uint64, err protocol.Error)
}

type Content_Status uint16

const (
	Content_Status_Unset Content_Status = iota
	Content_Status_Created
	Content_Status_Locked
	Content_Status_Hidden
	Content_Status_Deleted
	Content_Status_Blocked
)

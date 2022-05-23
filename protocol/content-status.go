

package protocol

import (
	"time"
)

type ContentStatus interface {
	ContentID() [16]byte          // content domain
	Status() ContentStatus_Status //
	Time() time.Time              // Save time
	RequestID() [16]byte          // user-request domain
}

type ContentStatus_StorageServices interface {
	Save(cs ContentStatus) error

	Count(contentID [16]byte) (length uint64, err error)
	Get(contentID [16]byte, versionOffset uint64) (cs ContentStatus, err error)
	Last(contentID [16]byte) (cs ContentStatus, length uint64, err error)

	
}

type ContentStatus_Status uint16

const (
	ContentStatus_Status_Unset ContentStatus_Status = iota
	ContentStatus_Status_Created
	ContentStatus_Status_Locked
	ContentStatus_Status_Hidden
	ContentStatus_Status_Deleted
	ContentStatus_Status_Blocked
)

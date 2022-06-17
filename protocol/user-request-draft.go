package protocol

// Store user requests as drafts for users to send later

// RequestID of this record use to prove how other records created as a chain to other records or standalone data.
// Prove how this record created in the chain of other records.
type RequestDraft interface {
	ID() [16]byte // RequestDraftID
	UserID() [16]byte
	UserConnectionID() [16]byte // Store to remember which ConnectionID/AppInstanceID set||chanaged this record.
	PageID() uint64
	ServiceID() uint64
	CompressID() uint64  // gzip
	Time() protocol.Time // save time
	Request() []byte     // Decode by service request structure and compress type.
}

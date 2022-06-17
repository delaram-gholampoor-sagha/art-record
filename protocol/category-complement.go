package protocol

type CategoryComplement interface {
	CategoryID() [16]byte   // category domain
	Priority() int32        // Complement priority
	ComplementID() [16]byte // category domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

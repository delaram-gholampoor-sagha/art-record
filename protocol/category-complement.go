package protocol

type CategoryComplement interface {
	CategoryID() [16]byte   // category domain
	Priority() int32         // Complement Priority
	ComplementID() [16]byte // category domain
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

package protocol

type CategorySubstitute interface {
	CategoryID() [16]byte   // category domain
	Priority() int32        // Substitute priority
	SubstituteID() [16]byte // category domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

package protocol

type ContentSubstitute interface {
	ContentID() [16]byte    // content domain
	Priority() int32        // substitute priority
	SubstituteID() [16]byte // content domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

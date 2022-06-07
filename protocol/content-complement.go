package protocol

type ContentComplement interface {
	ContentID() [16]byte    // content domain
	Priority() int32        // complement priority
	ComplementID() [16]byte // content domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

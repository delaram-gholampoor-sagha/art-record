package protocol

type ProductComplement interface {
	ProductID() [16]byte    // product domain
	Priority() uint64       // Complement priority
	ComplementID() [16]byte // product domain
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

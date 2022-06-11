package protocol

type ProductSubstitute interface {
	ProductID() [16]byte    // product domain
	Priority() uint64       // Substitute priority
	SubstituteID() [16]byte // product domain
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

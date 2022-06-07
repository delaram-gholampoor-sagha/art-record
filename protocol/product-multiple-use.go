package protocol


type ProductMultipleUse interface {
	ProductID() [16]byte        // product domain
	Maximum() uint64            // Maximum time this service can be served to buyer
	Timeout() protocol.Duration //
	Time() protocol.Time        // Save time
	RequestID() [16]byte        // user-request domain
}

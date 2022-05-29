package protocol


type ProductServiceInterval interface {
	productID()  [16]byte
	Min() uint64
	Max()   uint64
	Expected()  uint64 
 	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}


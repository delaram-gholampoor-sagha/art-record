package protocol

type ProductContent interface {
	ProductID() [16]byte // product domain
	ContentID() [16]byte // content domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

package protocol


type ProductTag interface {
	Tag() string         //
	ProductID() [16]byte // product domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

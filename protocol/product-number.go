package protocol


type ProductNumber interface {
	ProductID() [16]byte // product domain
	Minimum() uint64     //
	Maximum() uint64     // Maximum time can be served to buyers e.g. cinema tickets, panic room, ...
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

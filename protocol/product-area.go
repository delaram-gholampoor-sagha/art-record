package protocol

// ProductArea is proper or appropriate location indicate specific area not specific location e.g. Taxi, ...
type ProductArea interface {
	ProductID() [16]byte // product domain
	AreaID() [16]byte    // area-status domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

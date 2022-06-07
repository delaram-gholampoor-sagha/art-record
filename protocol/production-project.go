package protocol


type ProductionProject interface {
	// A lot number is an identification number assigned to a particular quantity or lot of material from a single manufacturer.
	// https://en.wikipedia.org/wiki/Lot_number
	LOT() [16]byte          // quiddity domain.
	ProductionID() [16]byte // production domain
	StartDate() uint64
	EndDate() uint64
	Number() uint64
	Status() uint8
}
	
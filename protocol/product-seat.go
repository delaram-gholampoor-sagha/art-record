package protocol

type ProductSeat interface {
	ProductID() [16]byte    // product domain
	TotalSeat() uint64      // max seat number
	Type() ProductSeat_Type //
	MapID() [16]byte        // object domain. to show background and rows and seats
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

type ProductSeat_Type uint8

const (
	ProductSeat_Type_Unset           ProductSeat_Type = 0
	ProductSeat_Type_ChooseSeatByOrg ProductSeat_Type = (1 << iota)
)

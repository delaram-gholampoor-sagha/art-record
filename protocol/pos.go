package protocol

// POS indicate the domain record data fields.
// POS or point-of-sale
type POS interface {
	PosID() [16]byte     // quiddity domain
	Type() POS_Type      //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type POS_StorageServices interface {
	Save(pos POS) protocol.Error

	Count(posID [16]byte) (numbers uint64, err protocol.Error)
	Get(posID [16]byte, versionOffset uint64) (pos POS, err protocol.Error)
	Last(posID [16]byte) (pos POS, numbers uint64, err protocol.Error)
}

type POS_Type uint64

// Common flags
const (
	POS_Type_Unset POS_Type = 0
	POS_Type_Cash  POS_Type = (1 << iota)
	POS_Type_Location
	POS_Type_Provider
	POS_Type_Product
	// POS_Type_
)

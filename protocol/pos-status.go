package protocol

// PosStatus indicate the domain record data fields.
type PosStatus interface {
	PosID() [16]byte     // pos domain
	Status() Pos_Status  //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type PosStatus_StorageServices interface {
	Save(ps PosStatus) protocol.Error

	Count(posID [16]byte) (numbers uint64, err protocol.Error)
	Get(posID [16]byte, versionOffset uint64) (ps PosStatus, err protocol.Error)
	Last(posID [16]byte) (ps PosStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status Pos_Status, offset, limit uint64) (posIDs [][16]byte, numbers uint64, err protocol.Error)

	protocol.EventTarget
}

type Pos_Status Quiddity_Status

const (
	Pos_Status_CashLimitInactivated = Pos_Status(Quiddity_Status_FreeFlag << iota)
)

package protocol


type PosStatus interface {
	PosID() [16]byte     // quiddity domain
	Status() Pos_Status  //
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type Pos_Status uint8

const (
	Pos_Status_Unset                Pos_Status = 0
	Pos_Status_PermanentInactivated Pos_Status = (1 << iota)
	Pos_Status_TemporaryInactivated
	Pos_Status_CashLimitInactivated
	Pos_Status_Blocked
)

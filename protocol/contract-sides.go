package protocol


type ContractSides interface {
	ContractID() [16]byte       // contract domain
	SideType() ContractSides_ST //
	UserID() [16]byte           // user-status domain. Can be any user type e.g. staff
	Time() protocol.Time        // Save time
	RequestID() [16]byte        // user-request domain
}

type ContractSides_ST uint8

const (
	ContractSides_ST_Unset ContractSides_ST = iota
	ContractSides_ST_First
	ContractSides_ST_Second
	ContractSides_ST_Third
	ContractSides_ST_Witness
	ContractSides_ST_Arbitrator
)

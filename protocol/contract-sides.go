package protocol

type ContractSides interface {
	ContractID() [16]byte       // contract domain
	SideType() ContractSides_ST //
	UserID() [16]byte           // user domain. Can be any user type e.g. staff
	UserSignature() []byte      //
	Time() protocol.Time        // save time
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
	ContractSides_ST_Expert
)

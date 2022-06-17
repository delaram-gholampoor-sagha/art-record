package protocol

type ContractStatus interface {
	ContractID() [16]byte    // contract domain
	Status() Contract_Status //
	Time() protocol.Time     // save time
	RequestID() [16]byte     // user-request domain
}

type Contract_Status Quiddity_Status

const (
	Contract_Status_Negotiation = Contract_Status(Quiddity_Status_FreeFlag << iota)
	Contract_Status_PartlySigned
	Contract_Status_Signed
	Contract_Status_LegalAction
	Contract_Status_Done
	Contract_Status_Canceled
)

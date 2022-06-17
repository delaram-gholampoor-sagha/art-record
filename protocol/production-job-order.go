package protocol

type ProductionJobOrder interface {
	// A lot number is an identification number assigned to a particular quantity or lot of material from a single manufacturer.
	// https://en.wikipedia.org/wiki/Lot_number
	LOT() [16]byte          // quiddity domain.
	ProductionID() [16]byte // production domain
	StartDate() uint64
	EndDate() uint64
	Number() uint64
	Status() uint8
}

type ProductionJobOrder_Status uint8

const (
	ProductionJobOrder_Status_Unset ProductionJobOrder_Status = iota
	ProductionJobOrder_Status_Rework
	ProductionJobOrder_Status_Rejection
	ProductionJobOrder_Status_Approved
)

package protocol

// a payment made to a professional person or to a professional or public body in exchange for advice or services.
type Fee struct {
	ID         [16]byte
	QuiddityID [16]byte
	Amount     uint64
	Status     uint8
}

type feeServices interface {
	FindFeeByQuiddittyID()
}

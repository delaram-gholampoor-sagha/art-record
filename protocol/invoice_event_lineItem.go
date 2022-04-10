package protocol

type Invoice_Line_Item struct {
	ID         [16]byte
	QuiddityID [16]byte
	Quantity   uint64
	Status     uint8
}

type Invoice_Line_Item_Service interface {
}

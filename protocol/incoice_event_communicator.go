package protocol

type Invoice_Communicator struct {
	ID [16]byte
	// can we mix both with a field like SenderAccountID ? I mean the communicators involved in this account ?
	SellerID    [16]byte
	BuyerID     [16]byte
	Status      uint8
	// ??
	RequestedID [16]byte
}

type Invoice_Communicator_service struct {
}

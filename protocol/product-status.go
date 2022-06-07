package protocol

type ProductStatus interface {
	ProductID() [16]byte    // product domain
	Status() Product_Status // overall
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

type Product_Status uint8

const (
	Product_Status_Unset Product_Status = iota
	Product_Status_Register
	Product_Status_Inactive

	Product_Status_InvoiceSell
	Product_Status_OrderSell

	// ActiveWithApprove
	// ActiveWithoutApprove
	// ActiveWithOrderer
)

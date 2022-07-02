package protocol

// InvoiceSides indicate the domain record data fields.
type InvoiceSides interface {
	InvoiceID() [16]byte  // invoice domain
	SellerID() [16]byte   // user domain
	CustomerID() [16]byte // user domain
	AgentID() [16]byte    // user domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type InvoiceSides_StorageServices interface {
	Save(is InvoiceSides) (numbers uint64, err protocol.Error)

	Get(invoiceID [16]byte) (is InvoiceSides, err protocol.Error)

	FindBySeller(sellerID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
	FindByCustomer(customerID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
	FindByAgent(agentID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
}

type InvoiceSides_Type uint8

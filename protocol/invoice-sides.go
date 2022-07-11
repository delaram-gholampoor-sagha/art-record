package protocol

import "../libgo/protocol"

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

type (
	InvoiceSides_Service_Register_Request interface {
		InvoiceID() [16]byte 
		SellerID() [16]byte   
		CustomerID() [16]byte 
		AgentID() [16]byte      
	}
	InvoiceSides_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	
	}
	
	InvoiceSides_Service_Get_Request interface { 
		InvoiceID() [16]byte 
	
	}
	InvoiceSides_Service_Get_Response interface {
		InvoiceSides
		Nv() protocol.NumberOfVersion
	}
	
	InvoiceSides_Service_FindBySeller_Request interface { 
		SellerID() [16]byte
		Offset() uint64 
		Limit() uint64
	}
	InvoiceSides_Service_FindBySeller_Response interface {
		InvoiceIDs() [][16]byte 
		Nv() protocol.NumberOfVersion
	
	}
	
	InvoiceSides_Service_FindByCustomer_Request interface { 
		CustomerID() uint64
		Offset() uint64 
		Limit() uint64
	}
	InvoiceSides_Service_FindByCustomer_Response interface {
		InvoiceIDs() [][16]byte 
		Nv() protocol.NumberOfVersion
	}
	//	FindByAgent(agentID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
	
	InvoiceSides_Service_FindByAgent_Request interface { 
		AgentID() [16]byte
		Offset() uint64 
		Limit() uint64
	}
	InvoiceSides_Service_FindByAgent_Response interface {
		InvoiceIDs() [][16]byte 
		Nv() protocol.NumberOfVersion
	}
)


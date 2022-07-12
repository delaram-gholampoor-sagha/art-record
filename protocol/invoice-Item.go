package protocol

import "../libgo/protocol"

// InvoiceItem indicate the domain record data fields.
// each InvoiceItem is immutable record and so use version mechanism to chain InvoiceItem in InvoiceID group.
// Ship service can add as an item to invoice
// Each item must check for reserve validity and make action in end of validity
type InvoiceItem interface {
	InvoiceID() [16]byte // invoice domain
	ProductID() [16]byte // product domain
	Quantity() uint64    // decimal >> 1.5 >> 1.2 >> 10 >> 0
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type InvoiceItem_StorageServices interface {
	Save(ii InvoiceItem) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(invoiceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (ii InvoiceItem, err protocol.Error)

	FindByProductID(productID uint64, offset, limit uint64) (invoiceIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
		InvoiceItem_Service_Register_Request interface {
		InvoiceID() [16]byte 
		ProductID() [16]byte 
		Quantity() uint64    
	}
	InvoiceItem_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	
	}

)

type (
		InvoiceItem_Service_Count_Request interface { 
		InvoiceID() [16]byte 
	
	}
	InvoiceItem_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)



type (
		InvoiceItem_Service_Get_Request interface { 
		InvoiceID() [16]byte
		VersionOffset() uint64
	}
	InvoiceItem_Service_Get_Response interface {
		InvoiceItem
	}
	
)



type (
	InvoiceItem_Service_FindByProductID_Request interface { 
		ProductID() uint64
		Offset() uint64
		Limit() uint64
	}
	InvoiceItem_Service_FindByProductID_Response interface {
		InvoiceIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	
	}
	
)
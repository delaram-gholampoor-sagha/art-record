package protocol

import "../libgo/protocol"

// InvoicePOS indicate the domain record data fields.
type InvoicePOS interface {
	InvoiceID() [16]byte // invoice domain
	PosID() [16]byte     // pos domain
	StaffID() [16]byte   // staff domain. or OrdererID that is creator of the invoice e.g. restaurant, drug prescription, sales agent,
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type InvoicePOS_StorageServices interface {
	Save(ip InvoicePOS) (numbers uint64, err protocol.Error)

	Get(invoiceID [16]byte) (ip InvoicePOS, err protocol.Error)

	FindByPos(posID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
	FindByStaff(staffID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
}

type (
	InvoicePOS_Service_Register_Request interface {
		InvoiceID() [16]byte 
		PosID() [16]byte     
		StaffID() [16]byte   
	}
	InvoicePOS_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	
	}
	
	InvoicePOS_Service_Get_Request interface { 
		InvoiceID() [16]byte 
	
	}
	InvoicePOS_Service_Get_Response interface {
		InvoicePOS
	}
	
	InvoicePOS_Service_FindByPos_Request interface { 
		PosID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	InvoicePOS_Service_FindByPos_Response interface {
		InvoiceIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
	InvoicePOS_Service_FindByStaff_Request interface { 
		StaffID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	InvoicePOS_Service_FindByStaff_Response interface {
		InvoiceIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
)
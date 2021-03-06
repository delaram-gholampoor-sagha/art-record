/* For license and copyright information please see LEGAL file in repository */

package art

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
	Save(ip InvoicePOS) (nv protocol.NumberOfVersion, err protocol.Error)

	Get(invoiceID [16]byte) (ip InvoicePOS, nv protocol.NumberOfVersion ,err protocol.Error)

	FindByPos(posID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	FindByStaff(staffID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	InvoicePOS_Service_Register_Request interface {
		InvoiceID() [16]byte 
		PosID() [16]byte     
		StaffID() [16]byte   
	}

	InvoicePOS_Service_Register_Response = protocol.NumberOfVersion

)

type (
	InvoicePOS_Service_Get_Request interface { 
		InvoiceID() [16]byte 
	}

	InvoicePOS_Service_Get_Response1 = InvoicePOS
	InvoicePOS_Service_Get_Response2 = protocol.NumberOfVersion
	
)



type (
	InvoicePOS_Service_FindByPos_Request interface { 
		PosID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	InvoicePOS_Service_FindByPos_Response1 interface {
		InvoiceIDs() [][16]byte
	}

	InvoicePOS_Service_FindByPos_Response2 = protocol.NumberOfVersion
	
)

type (
	InvoicePOS_Service_FindByStaff_Request interface { 
		StaffID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	InvoicePOS_Service_FindByStaff_Response1 interface {
		InvoiceIDs() [][16]byte
	}

	InvoicePOS_Service_FindByStaff_Response2 = protocol.NumberOfVersion
)
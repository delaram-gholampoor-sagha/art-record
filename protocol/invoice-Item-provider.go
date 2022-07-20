/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// InvoiceItemProvider indicate the domain record data fields.
type InvoiceItemProvider interface {
	InvoiceID() [16]byte // invoice domain
	ProductID() [16]byte // product domain
	StaffID() [16]byte   // staff domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type InvoiceItemProvider_StorageServices interface {
	Save(iip InvoiceItemProvider) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(invoiceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iip InvoiceItemProvider, nv protocol.NumberOfVersion, err protocol.Error)

	FindByStaff(staffID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	InvoiceItemProvider_Service_Register_Request interface {
		InvoiceID() [16]byte 
		ProductID() [16]byte 
		StaffID() [16]byte   
	}

	InvoiceItemProvider_Service_Register_Response = protocol.NumberOfVersion

)


type (
	InvoiceItemProvider_Service_Count_Request interface { 
		InvoiceID() [16]byte 
	}

	InvoiceItemProvider_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	InvoiceItemProvider_Service_Get_Request interface { 
		InvoiceID() [16]byte
		versionOffset() uint64
	}


	InvoiceItemProvider_Service_Get_Response1 = InvoiceItemProvider
	InvoiceItemProvider_Service_Get_Response2 = protocol.NumberOfVersion
	
)



type (
	InvoiceItemProvider_Service_FindByStaff_Request interface { 
		StaffId() [16]byte
		Offset() [16]byte
		Limit() uint64
	
	}
	InvoiceItemProvider_Service_FindByStaff_Response1 interface {
		InvoiceIDs() [][16]byte
	}

	InvoiceItemProvider_Service_FindByStaff_Response2 = protocol.NumberOfVersion
)
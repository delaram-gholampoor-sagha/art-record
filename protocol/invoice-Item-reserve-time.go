/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// InvoiceItemReserveTime indicate the domain record data fields.
type InvoiceItemReserveTime interface {
	InvoiceID() [16]byte  // invoice domain
	ProductID() [16]byte  // product domain
	Start() protocol.Time //
	End() protocol.Time   //
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type InvoiceItemReserveTime_StorageServices interface {
	Save(itr InvoiceItemReserveTime) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(invoiceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (itr InvoiceItemReserveTime, nv protocol.NumberOfVersion, err protocol.Error)
}



type (
	InvoiceItemReserveTime_Service_Register_Request interface {
		InvoiceID() [16]byte 
		ProductID() [16]byte 
		Start() protocol.Time 
		End() protocol.Time   
	}
	InvoiceItemReserveTime_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
	InvoiceItemReserveTime_Service_Count_Request interface { 
		InvoiceID() [16]byte 
	}
	InvoiceItemReserveTime_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)



type (
	InvoiceItemReserveTime_Service_Get_Request interface { 
		InvoiceID() [16]byte
		VersionOffset() uint64
	}
	InvoiceItemReserveTime_Service_Get_Response interface {
		InvoiceItemReserveTime
		NumberOfVersion() protocol.NumberOfVersion
	}
)
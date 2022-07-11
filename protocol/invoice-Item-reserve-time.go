package protocol

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
	Save(itr InvoiceItemReserveTime) (numbers uint64, err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (itr InvoiceItemReserveTime, numbers uint64, err protocol.Error)
}

type (
	InvoiceItemReserveTime_Service_Register_Request interface {
		InvoiceID() [16]byte 
		ProductID() [16]byte 
		Start() protocol.Time 
		End() protocol.Time   
	}
	InvoiceItemReserveTime_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	InvoiceItemReserveTime_Service_Count_Request interface { 
		InvoiceID() [16]byte 
	}
	InvoiceItemReserveTime_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	InvoiceItemReserveTime_Service_Get_Request interface { 
		InvoiceID() [16]byte
		VersionOffset() uint64
	}
	InvoiceItemReserveTime_Service_Get_Response interface {
		InvoiceItemReserveTime
		Nv() protocol.NumberOfVersion
	}
)
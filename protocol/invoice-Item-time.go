package protocol

import "../libgo/protocol"

// InvoiceItemTime indicate the domain record data fields.
type InvoiceItemTime interface {
	InvoiceID() [16]byte        // invoice domain
	ProductID() [16]byte        // product domain
	Type() InvoiceItemTime_Type //
	Time() protocol.Time        // save time
	RequestID() [16]byte        // user-request domain
}

type InvoiceItemTime_StorageServices interface {
	Save(iit InvoiceItemTime) (numbers uint64, err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iit InvoiceItemTime, err protocol.Error)
}

type (
	InvoiceItemTime_Service_Register_Request interface {
		InvoiceID() [16]byte 
		ProductID() [16]byte 
		Type() InvoiceItemTime_Type
	}
	InvoiceItemTime_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	
	}
	
	InvoiceItemTime_Service_Count_Request interface { 
		InvoiceID() [16]byte 
	
	}
	InvoiceItemTime_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	InvoiceItemTime_Service_Get_Request interface { 
		InvoiceID() [16]byte
		VersionOffset() uint64
	
	
	}
	InvoiceItemTime_Service_Get_Response interface {
		InvoiceItemTime
	}
	
)

type InvoiceItemTime_Type uint8

const (
	InvoiceItemTime_Type_Unset InvoiceItemTime_Type = iota
	InvoiceItemTime_Type_Enter
	InvoiceItemTime_Type_Exit
	InvoiceItemTime_Type_ManualEnter
	InvoiceItemTime_Type_ManualExit
	InvoiceItemTime_Type_Start
	InvoiceItemTime_Type_End
)

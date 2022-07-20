/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// InvoiceItemSeat indicate the domain record data fields.
type InvoiceItemSeat interface {
	InvoiceID() [16]byte            // invoice domain
	ProductID() [16]byte            // product domain
	SeatID() uint64                 // seat number
	Status() InvoiceItemSeat_Status //
	Time() protocol.Time            // save time
	RequestID() [16]byte            // user-request domain
}

type InvoiceItemSeat_StorageServices interface {
	Save(iis InvoiceItemSeat) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(invoiceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iis InvoiceItemSeat, nv protocol.NumberOfVersion, err protocol.Error)
}

type InvoiceItemSeat_Status Quiddity_Status

const (
	InvoiceItemSeat_Status_Voided = InvoiceItemSeat_Status(Quiddity_Status_FreeFlag << iota)
)


type (
	InvoiceItemSeat_Service_Register_Request interface {
		InvoiceID() [16]byte           
		ProductID() [16]byte           
		SeatID() uint64                
		Status() InvoiceItemSeat_Status      
	}

	InvoiceItemSeat_Service_Register_Response = protocol.NumberOfVersion
)

type (
	InvoiceItemSeat_Service_Count_Request interface { 
		InvoiceID() [16]byte 
	}

	InvoiceItemSeat_Service_Count_Response = protocol.NumberOfVersion
)

type (
	InvoiceItemSeat_Service_Get_Request interface { 
		InvoiceID() [16]byte
		versionOffset() uint64
	}
	
	InvoiceItemSeat_Service_Get_Response1 = InvoiceItemSeat
	InvoiceItemSeat_Service_Get_Response2 = protocol.NumberOfVersion
)



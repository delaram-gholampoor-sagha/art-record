/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// InvoiceDiscount indicate the domain record data fields.
type InvoiceDiscount interface {
	InvoiceID() [16]byte                // invoice domain
	DiscountID() [16]byte               // discount domain
	Discounted() protocol.AmountOfMoney //
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}

type InvoiceDiscount_StorageServices interface {
	Save(iv InvoiceDiscount) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(invoiceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iv InvoiceDiscount, nv protocol.NumberOfVersion, err protocol.Error)

	FindByDiscountID(discountID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	InvoiceDiscount_Service_Register_Request interface {
		InvoiceID() [16]byte               
		DiscountID() [16]byte               
		Discounted() protocol.AmountOfMoney 
	}

	InvoiceDiscount_Service_Register_Response = protocol.NumberOfVersion

)

type (
	InvoiceDiscount_Service_Count_Request interface { 
		InvoiceID() [16]byte 
	}

	InvoiceDiscount_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	InvoiceDiscount_Service_Get_Request interface { 
		InvoiceID() [16]byte
		versionOffset() uint64
	}

	InvoiceDiscount_Service_Get_Response1 = 	InvoiceDiscount
	InvoiceDiscount_Service_Get_Response2 = protocol.NumberOfVersion
	
)


type (
	InvoiceDiscount_Service_FindByDiscountID_Request interface { 
		DiscountID() [16]byte
		Offset() [16]byte
		Limit() uint64
	
	}
	InvoiceDiscount_Service_FindByDiscountID_Response1 interface {
		InvoiceIDs() [][16]byte
	}

	InvoiceDiscount_Service_FindByDiscountID_Response2 = protocol.NumberOfVersion
)
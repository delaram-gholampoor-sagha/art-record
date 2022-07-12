package protocol

import "../libgo/protocol"

// InvoiceVoucherAnalyticUtcDay indicate the domain record data fields.
type InvoiceVoucherAnalyticUtcDay interface {
	Day() utc.DayElapsed                    //
	Redeemed() uint64                       // Total
	RedeemedAmount() protocol.AmountOfMoney // Total
	Time() protocol.Time                    // save time
}


type InvoiceVoucherAnalyticUtcDay_StorageServices interface {
	Save(iv InvoiceVoucherAnalyticUtcDay) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(invoiceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iv InvoiceVoucherAnalyticUtcDay, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	InvoiceVoucherAnalyticUtcDay_Service_Register_Request interface {
		InvoiceID() [16]byte               
		VoucherID() [16]byte               
		Discounted() protocol.AmountOfMoney 
	}
	InvoiceVoucherAnalyticUtcDay_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	
	}
	
	InvoiceVoucherAnalyticUtcDay_Service_Count_Request interface {
		InvoiceID() [16]byte
		
	
	}
	InvoiceVoucherAnalyticUtcDay_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
	
	InvoiceVoucherAnalyticUtcDay_Service_Get_Request interface {
		InvoiceID() [16]byte    
		VersionOffset() uint64
	}
	InvoiceVoucherAnalyticUtcDay_Service_Get_Response interface {
		InvoiceVoucherAnalyticUtcDay
		NumberOfVersion() protocol.NumberOfVersion
	}

)
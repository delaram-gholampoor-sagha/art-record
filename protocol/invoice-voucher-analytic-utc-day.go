package protocol


// InvoiceVoucherAnalyticUtcDay indicate the domain record data fields.
type InvoiceVoucherAnalyticUtcDay interface {
	Day() utc.DayElapsed                    //
	Redeemed() uint64                       // Total
	RedeemedAmount() protocol.AmountOfMoney // Total
	Time() protocol.Time                    // save time
}


type InvoiceVoucherAnalyticUtcDay_StorageServices interface {
	Save(iv InvoiceVoucherAnalyticUtcDay) (numbers uint64, err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iv InvoiceVoucherAnalyticUtcDay, numbers uint64, err protocol.Error)


}


type (
	InvoiceVoucherAnalyticUtcDay_Service_Register_Request interface {
		InvoiceID() [16]byte               
		VoucherID() [16]byte               
		Discounted() protocol.AmountOfMoney 
	}
	InvoiceVoucherAnalyticUtcDay_Service_Register_Response interface {
		Numbers() uint64
	
	}
	
	InvoiceVoucherAnalyticUtcDay_Service_Count_Request interface {
		InvoiceID() [16]byte
		
	
	}
	InvoiceVoucherAnalyticUtcDay_Service_Count_Response interface {
		Numbers() uint64
	}
	
	InvoiceVoucherAnalyticUtcDay_Service_Get_Request interface {
		InvoiceID() [16]byte    
		VersionOffset() uint64
	}
	InvoiceVoucherAnalyticUtcDay_Service_Get_Response interface {
		InvoiceVoucherAnalyticUtcDay
		Numbers() uint64
	}

)
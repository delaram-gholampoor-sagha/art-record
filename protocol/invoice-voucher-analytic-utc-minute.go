package protocol

// InvoiceVoucherAnalyticUtcMinute indicate the domain record data fields.
type InvoiceVoucherAnalyticUtcMinute interface {
	Minute() utc.MinuteElapsed              //
	Redeemed() uint64                       // Total
	RedeemedAmount() protocol.AmountOfMoney // Total
	Time() protocol.Time                    // save time
}

type InvoiceVoucherAnalyticUtcMinute_StorageServices interface {
	Save(iva InvoiceVoucherAnalyticUtcMinute) (numbers uint64, err protocol.Error)

	Count() (numbers uint64, err protocol.Error)
	Get(versionOffset uint64) (iva InvoiceVoucherAnalyticUtcMinute, numbers uint64, err protocol.Error)
}

type (
	InvoiceVoucherAnalyticUtcMinute_Service_Register_Request interface {
		Minute() utc.MinuteElapsed              
		Redeemed() uint64                       
		RedeemedAmount() protocol.AmountOfMoney 
	}
	InvoiceVoucherAnalyticUtcMinute_Service_Register_Response interface {
		Numbers() uint64
	
	}
	
	InvoiceVoucherAnalyticUtcMinute_Service_Count_Request interface {
		Numbers() uint64
	
	}
	InvoiceVoucherAnalyticUtcMinute_Service_Count_Response interface {
		
	}
	
	InvoiceVoucherAnalyticUtcMinute_Service_Get_Request interface {  
		VersionOffset() uint64
	}
	InvoiceVoucherAnalyticUtcMinute_Service_Get_Response interface {
		InvoiceVoucherAnalyticUtcMinute
		Numbers() uint64
	}
	
)
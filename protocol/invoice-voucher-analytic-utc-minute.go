package protocol

import "../libgo/protocol"

// InvoiceVoucherAnalyticUtcMinute indicate the domain record data fields.
type InvoiceVoucherAnalyticUtcMinute interface {
	Minute() utc.MinuteElapsed              //
	Redeemed() uint64                       // Total
	RedeemedAmount() protocol.AmountOfMoney // Total
	Time() protocol.Time                    // save time
}

type InvoiceVoucherAnalyticUtcMinute_StorageServices interface {
	Save(iva InvoiceVoucherAnalyticUtcMinute) (nv protocol.NumberOfVersion, err protocol.Error)

	Count() (nv protocol.NumberOfVersion, err protocol.Error)
	Get(versionOffset uint64) (iva InvoiceVoucherAnalyticUtcMinute, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
		InvoiceVoucherAnalyticUtcMinute_Service_Register_Request interface {
		Minute() utc.MinuteElapsed              
		Redeemed() uint64                       
		RedeemedAmount() protocol.AmountOfMoney 
	}
	InvoiceVoucherAnalyticUtcMinute_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	
	}

)

type (
		InvoiceVoucherAnalyticUtcMinute_Service_Count_Request interface {
		NumberOfVersion() protocol.NumberOfVersion
	
	}
	InvoiceVoucherAnalyticUtcMinute_Service_Count_Response interface {
		
	}
	
)


type (
	
	InvoiceVoucherAnalyticUtcMinute_Service_Get_Request interface {  
		VersionOffset() uint64
	}
	InvoiceVoucherAnalyticUtcMinute_Service_Get_Response interface {
		InvoiceVoucherAnalyticUtcMinute
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)
package protocol

import "../libgo/protocol"

// VoucherAnalyticUtcMinute indicate the domain record data fields.
type VoucherAnalyticUtcMinute interface {
	Minute() utc.MinuteElapsed            //
	Issued() uint64                       // Total
	IssuedAmount() protocol.AmountOfMoney // Total
	Time() protocol.Time                  // save time
}

type VoucherAnalyticUtcMinute_StorageServices interface {
	Save(va VoucherAnalyticUtcMinute) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (va VoucherAnalyticUtcMinute, numbers uint64, err protocol.Error)
}

type (
	VoucherAnalyticUtcMinute_Service_Register_Request interface{
		Minute() utc.MinuteElapsed           
  	Issued() uint64                       
	  IssuedAmount() protocol.AmountOfMoney
	}
	
	VoucherAnalyticUtcMinute_Service_Register_Response interface{
		Nv() protocol.NumberOfVersion
	}
	
	VoucherAnalyticUtcMinute_Service_Count_Request interface{
		VoucherID() [16]byte
	}
	
	VoucherAnalyticUtcMinute_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	VoucherAnalyticUtcMinute_Service_Get_Request interface{
		VoucherID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherAnalyticUtcMinute_Service_Get_Response interface{
		VoucherAnalyticUtcMinute
		Nv() protocol.NumberOfVersion
	}
)

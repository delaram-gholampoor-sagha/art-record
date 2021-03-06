/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// VoucherAnalyticUtcMinute indicate the domain record data fields.
type VoucherAnalyticUtcMinute interface {
	Minute() utc.MinuteElapsed            //
	Issued() uint64                       // Total
	IssuedAmount() protocol.AmountOfMoney // Total
	Time() protocol.Time                  // save time
}

type VoucherAnalyticUtcMinute_StorageServices interface {
	Save(va VoucherAnalyticUtcMinute) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (va VoucherAnalyticUtcMinute, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	VoucherAnalyticUtcMinute_Service_Register_Request interface{
		Minute() utc.MinuteElapsed           
  	Issued() uint64                       
	  IssuedAmount() protocol.AmountOfMoney
	}
	
	VoucherAnalyticUtcMinute_Service_Register_Response = protocol.NumberOfVersion
)

type (
	VoucherAnalyticUtcMinute_Service_Count_Request interface{
		VoucherID() [16]byte
	}
	
	VoucherAnalyticUtcMinute_Service_Count_Response = protocol.NumberOfVersion
)



type (
	VoucherAnalyticUtcMinute_Service_Get_Request interface{
		VoucherID() [16]byte
		versionOffset() uint64
	}
	
	VoucherAnalyticUtcMinute_Service_Get_Response1 = VoucherAnalyticUtcMinute
	VoucherAnalyticUtcMinute_Service_Get_Response2 = protocol.NumberOfVersion
)

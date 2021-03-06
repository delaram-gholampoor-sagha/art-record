/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// VoucherAnalyticUtcDay indicate the domain record data fields.
type VoucherAnalyticUtcDay interface {
	Day() utc.DayElapsed                  //
	Issued() uint64                       // Total
	IssuedAmount() protocol.AmountOfMoney // Total
	Time() protocol.Time                  // save time
}

type VoucherAnalyticUtcDay_StorageServices interface {
	Save(va VoucherAnalyticUtcDay) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (va VoucherAnalyticUtcDay, nv protocol.NumberOfVersion, err protocol.Error)
}


type (

	VoucherAnalyticUtcDay_Service_Register_Request interface{
		Day() utc.DayElapsed                  
	  Issued() uint64                       
	  IssuedAmount() protocol.AmountOfMoney 
	}
	
	VoucherAnalyticUtcDay_Service_Register_Response = protocol.NumberOfVersion

)


type (
	VoucherAnalyticUtcDay_Service_Count_Request interface{
		VoucherID() [16]byte
	}
	
	VoucherAnalyticUtcDay_Service_Count_Response = protocol.NumberOfVersion

)


type (
	VoucherAnalyticUtcDay_Service_Get_Request interface{
		VoucherID() [16]byte
		versionOffset() uint64
	}
	
	VoucherAnalyticUtcDay_Service_Get_Response1 = VoucherAnalyticUtcDay
	VoucherAnalyticUtcDay_Service_Get_Response2 = protocol.NumberOfVersion
	
)
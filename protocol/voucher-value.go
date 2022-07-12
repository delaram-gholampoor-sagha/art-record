package protocol

import "../libgo/protocol"

// VoucherValue indicate the domain record data fields.
type VoucherValue interface {
	VoucherID() [16]byte           // voucher domain
	Each() uint8                   // Each time use
	Percentage() math.PerMyriad    // Max discount percentage per use
	Price() protocol.AmountOfMoney // Max discount price per use
	Time() protocol.Time           // save time
	RequestID() [16]byte           // user-request domain
}

type VoucherValue_StorageServices interface {
	Save(vv VoucherValue) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vv VoucherValue, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	VoucherValue_Service_Register_Request interface {
		VoucherID() [16]byte           
		Each() uint8                   
		Percentage() math.PerMyriad    
		Price() protocol.AmountOfMoney 
	}
	
	VoucherValue_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)


type (

	
	VoucherValue_Service_Count_Request interface {
		VoucherID() [16]byte
	}
	
	VoucherValue_Service_Count_Response interface {
	 NumberOfVersion() protocol.NumberOfVersion 
	}
	
)

type (
	VoucherValue_Service_Get_Request interface {
		voucherID() [16]byte
		versionOffset() uint64
	}
	
	VoucherValue_Service_Get_Response interface {
		 VoucherValue
		 NumberOfVersion() protocol.NumberOfVersion
	}

)
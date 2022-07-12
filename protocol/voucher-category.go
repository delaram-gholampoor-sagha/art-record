package protocol

import "../libgo/protocol"

// VoucherCategory indicate the domain record data fields.
// Use to restrict the voucher on specific category.
type VoucherCategory interface {
	VoucherID() [16]byte  // voucher domain
	Each() uint8          // Each time use
	CategoryID() [16]byte // category domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type VoucherCategory_StorageServices interface {
	Save(vc VoucherCategory) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vc VoucherCategory, nv protocol.NumberOfVersion, err protocol.Error)
}



type (

	VoucherCategory_Service_Register_Request interface{
	   VoucherID() [16]byte  
	   Each() uint8          
	   CategoryID() [16]byte 
	}
	
	VoucherCategory_Service_Register_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
	VoucherCategory_Service_Count_Request interface{
		VoucherID() [16]byte
	}
	
	VoucherCategory_Service_Count_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
	VoucherCategory_Service_Get_Request interface{
		VoucherID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherCategory_Service_Get_Response interface{
		VoucherCategory
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)
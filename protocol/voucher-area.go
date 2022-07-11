package protocol

import "../libgo/protocol"

// VoucherArea indicate the domain record data fields.
// VoucherArea is proper or appropriate location indicate specific area not specific location
// It's rule depend on products in invoices e.g. source or destination of taxi as a service(product).
type VoucherArea interface {
	VoucherID() [16]byte // voucher domain
	AreaID() [16]byte    // area domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type VoucherArea_StorageServices interface {
	Save(va VoucherArea) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (va VoucherArea, numbers uint64, err protocol.Error)

	FindByArea(AreaID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, numbers uint64, err protocol.Error)
}

type (
	VoucherArea_Service_Register_Request interface{
		VoucherID() [16]byte 
	  AreaID() [16]byte    
	}
	
	VoucherArea_Service_Register_Response interface{
		Nv() protocol.NumberOfVersion
	}
	
	VoucherArea_Service_Count_Request interface{
		VoucherID() [16]byte
	}
	
	VoucherArea_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	VoucherArea_Service_Get_Request interface{
		VoucherID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherArea_Service_Get_Response interface{
		VoucherArea
		Nv() protocol.NumberOfVersion
	}
	
	
	VoucherArea_Service_FindByArea_Request interface{
		AreaID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	VoucherArea_Service_FindByArea_Response interface{
		VoucherIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
)
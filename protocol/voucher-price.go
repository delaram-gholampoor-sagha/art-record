package protocol

import "time"

type VoucherPrice interface {
	VoucherID() [16]byte // voucher domain
	Each() uint8         // Each time use
	Price() int64        // Max discount price per use
	Time() time.Time     // Save time
	RequestID() [16]byte // user-request domain
}

type VoucherPrice_StorageServices interface {
	Save(vp VoucherPrice) error

	Count(id [16]byte) (length uint64, err error)
	Get(id [16]byte, versionOffset uint64) (vp VoucherPrice, err error)
	
}

type (
	VoucherPrice_Service_Register_Request interface{
		VoucherID() [16]byte 
	  Each() uint8        
	  Price() int64       
	}

	VoucherPrice_Service_Register_Response interface{
      Numbers() uint64      
	}
	

	VoucherPrice_Service_Count_Request interface{
		ID() [16]byte
	}
	
	VoucherPrice_Service_Count_Response interface{
		Length() uint64
	}
	VoucherPrice_Service_Get_Request interface{
		ID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherPrice_Service_Get_Response interface{
		VoucherPrice
	}
	
)
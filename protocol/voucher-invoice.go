package protocol

import "../libgo/protocol"

// VoucherInvoice indicate the domain record data fields.
type VoucherInvoice interface {
	VoucherID() [16]byte // voucher domain
	Each() uint8         // Each time use
	MinPrice() int64     // Minimum invoice price
	MinAmount() uint64   // Minimum invoice product numbers
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type VoucherInvoice_StorageServices interface {
	Save(vi VoucherInvoice) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vi VoucherInvoice, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	VoucherInvoice_Service_Register_Request interface{
		VoucherID() [16]byte 
  	Each() uint8        
	  MinPrice() int64     
	  MinAmount() uint64   
	}
	
	VoucherInvoice_Service_Register_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}
	
	VoucherInvoice_Service_Count_Request interface{
		VoucherID() [16]byte
	}
	
	VoucherInvoice_Service_Count_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}
	VoucherInvoice_Service_Get_Request interface{
		VoucherID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherInvoice_Service_Get_Response interface{
		VoucherInvoice
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)
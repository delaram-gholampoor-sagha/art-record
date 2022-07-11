package protocol

import "../libgo/protocol"

// VoucherProduct restrict use voucher on specific product
type VoucherProduct interface {
	VoucherID() [16]byte // voucher domain
	Each() uint8         // Each time use
	ProductID() [16]byte // product domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type VoucherProduct_StorageServices interface {
	Save(vp VoucherProduct) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vp VoucherProduct, numbers uint64, err protocol.Error)

	FindByProduct(productID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, numbers uint64, err protocol.Error)
}


type(
	VoucherProduct_Service_Register_Request interface{
		VoucherID() [16]byte 
 	  Each() uint8         
	  ProductID() [16]byte 
	}
	
	VoucherProduct_Service_Register_Response interface{
		Nv() protocol.NumberOfVersion
	}
	
	VoucherProduct_Service_Count_Request interface{
		VoucherID() [16]byte
	}
	
	VoucherProduct_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	VoucherProduct_Service_Get_Request interface{
		VoucherID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherProduct_Service_Get_Response interface{
		VoucherProduct
		Nv() protocol.NumberOfVersion
	}
	
	
	VoucherProduct_Service_FindByProduct_Request interface{
		ProductID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	VoucherProduct_Service_FindByProduct_Response interface{
		VoucherIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
)
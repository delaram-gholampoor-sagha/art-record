/* For license and copyright information please see LEGAL file in repository */

package art

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
	Save(vp VoucherProduct) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vp VoucherProduct, nv protocol.NumberOfVersion, err protocol.Error)

	FindByProduct(productID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	VoucherProduct_Service_Register_Request interface{
		VoucherID() [16]byte 
 	  Each() uint8         
	  ProductID() [16]byte 
	}
	
	VoucherProduct_Service_Register_Response = protocol.NumberOfVersion
)


type (
	VoucherProduct_Service_Count_Request interface{
		VoucherID() [16]byte
	}
	
	VoucherProduct_Service_Count_Response = protocol.NumberOfVersion

)

type (

		VoucherProduct_Service_Get_Request interface{
		VoucherID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherProduct_Service_Get_Response1 = 	VoucherProduct
	VoucherProduct_Service_Get_Response2 = protocol.NumberOfVersion
	
)

type (
	VoucherProduct_Service_FindByProduct_Request interface{
		ProductID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	VoucherProduct_Service_FindByProduct_Response interface{
		VoucherIDs() [][16]byte
	}

	VoucherProduct_Service_FindByProduct_Response2 = protocol.NumberOfVersion
)
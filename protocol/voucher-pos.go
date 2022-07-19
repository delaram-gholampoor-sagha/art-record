/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// VoucherPOS indicate the domain record data fields.
type VoucherPOS interface {
	VoucherID() [16]byte // voucher domain
	PosID() [16]byte     // pos domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type VoucherPOS_StorageServices interface {
	Save(vp VoucherPOS) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vp VoucherPOS, nv protocol.NumberOfVersion, err protocol.Error)

	FindByPOS(posID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}



type (


		VoucherPOS_Service_Register_Request interface {
			VoucherID() [16]byte 
    	PosID() [16]byte    
	  }
	   
		VoucherPOS_Service_Register_Response interface {
			NumberOfVersion() protocol.NumberOfVersion
		}
)

type (
		VoucherPOS_Service_Count_Request interface {
			VoucherID() [16]byte
		}
		
		VoucherPOS_Service_Count_Response interface {
			NumberOfVersion() protocol.NumberOfVersion
		}

)


type (
			VoucherPOS_Service_Get_Request interface {
			VoucherID() [16]byte
			VersionOffset() uint64
		
		}
		
		VoucherPOS_Service_Get_Response interface {
			VoucherPOS
			NumberOfVersion() protocol.NumberOfVersion
		}

)

type (
		VoucherPOS_Service_FindByPOS_Request interface {
			PosID() [16]byte
			OffSet() uint64
			Limit() uint64
		}
		
		VoucherPOS_Service_FindByPOS_Response interface {
			VoucherIDs() [][16]byte
			NumberOfVersion() protocol.NumberOfVersion
		}
)
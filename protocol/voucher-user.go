package protocol

import "../libgo/protocol"

// VoucherUser indicate the domain record data fields.
// Use to restrict the voucher usage by specific users.
type VoucherUser interface {
	VoucherID() [16]byte // voucher domain
	UserID() [16]byte    // user domain.
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type VoucherUser_StorageServices interface {
	Save(vu VoucherUser) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vu VoucherUser, nv protocol.NumberOfVersion, err protocol.Error)

	FindByUser(userID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	VoucherUser_Service_Register_Request interface{
		VoucherID() [16]byte 
		UserID() [16]byte    
	}

	VoucherUser_Service_Register_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}
)

type (
	VoucherUser_Service_Count_Request interface{
		VoucherID() [16]byte
	}

	VoucherUser_Service_Count_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (

		VoucherUser_Service_Get_Request interface{
		VoucherID() [16]byte
		VersionOffset() uint64
	}

	VoucherUser_Service_Get_Response interface{
		VoucherUser
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (

		VoucherUser_Service_FindByUser_Request interface{
		UserID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	VoucherUser_Service_FindByUser_Response interface{
		VoucherIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
)
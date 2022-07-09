package protocol


// VoucherUserType indicate the domain record data fields.
// Use to restrict the voucher usage by specific user type.
type VoucherUserType interface {
	VoucherID() [16]byte         // voucher domain
	UserType() protocol.UserType //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type VoucherUserType_StorageServices interface {
	Save(vut VoucherUserType) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vut VoucherUserType, numbers uint64, err protocol.Error)

	FilterByUserType(userType protocol.UserType, offset, limit uint64) (voucherIDs [][16]byte, numbers uint64, err protocol.Error)
}


type (
	VoucherUserType_Service_Register_Request interface{
			VoucherID() [16]byte         
			UserType() protocol.UserType 
	}

	VoucherUserType_Service_Register_Response interface{
		Numbers() uint64
	}

	VoucherUserType_Service_Count_Request interface{
		VoucherID() [16]byte
	}

	VoucherUserType_Service_Count_Response interface{
		Numbers() uint64
	}
	
	VoucherUserType_Service_Get_Request interface{
		VoucherID() [16]byte
		VersionOffset() uint64
	}

	VoucherUserType_Service_Get_Response interface{
		VoucherUserType
		Numbers() uint64
	}

	VoucherUserType_Service_FilterByUserType_Request interface{
		UserID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	VoucherUserType_Service_FilterByUserType_Response interface{
		VoucherIDs() [][16]byte
		Numbers() uint64
	}
)



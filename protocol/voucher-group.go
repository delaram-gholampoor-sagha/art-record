package protocol

import "../libgo/protocol"

// VoucherGroup indicate the domain record data fields.
type VoucherGroup interface {
	VoucherID() [16]byte // voucher domain
	GroupID() [16]byte   // group domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type VoucherGroup_StorageServices interface {
	Save(vg VoucherGroup) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vg VoucherGroup, nv protocol.NumberOfVersion, err protocol.Error)

	FindByGroup(groupID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	VoucherGroup_Service_Register_Request interface{
		VoucherID() [16]byte 
	  GroupID() [16]byte   
	}
	
	VoucherGroup_Service_Register_Response interface{
		Nv() protocol.NumberOfVersion
	}
	
	VoucherGroup_Service_Count_Request interface{
		VoucherID() [16]byte
	}
	
	VoucherGroup_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	VoucherGroup_Service_Get_Request interface{
		VoucherID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherGroup_Service_Get_Response interface{
		VoucherGroup
		Nv() protocol.NumberOfVersion
	}
	
	
	VoucherGroup_Service_FindByGroup_Request interface{
		GroupID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	VoucherGroup_Service_FindByGroup_Response interface{
		VoucherIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
)
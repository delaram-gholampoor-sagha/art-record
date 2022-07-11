package protocol

import "../libgo/protocol"

type VoucherLocation interface {
	VoucherID() [16]byte          // voucher domain
	BuildingLocationID() [16]byte // building-location domain
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type VoucherLocation_StorageServices interface {
	Save(vl VoucherLocation) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vl VoucherLocation, nv protocol.NumberOfVersion, err protocol.Error)

	FindByBuildingLocation(buildingLocationID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	VoucherLocation_Service_Register_Request interface {
		VoucherID() [16]byte
		BuildingLocationID() [16]byte
	}

	VoucherLocation_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}

	VoucherLocation_Service_Count_Request interface {
		VoucherID() [16]byte
	}

	VoucherLocation_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	VoucherLocation_Service_Get_Request interface {
		VoucherID() [16]byte
		VersionOffset() uint64
	}

	VoucherLocation_Service_Get_Response interface {
		VoucherLocation
		Nv() protocol.NumberOfVersion
	}

	VoucherLocation_Service_FindByBuildingLocation_Request interface {
		BuildingLocationID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	VoucherLocation_Service_FindByBuildingLocation_Response interface {
		VoucherIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
)
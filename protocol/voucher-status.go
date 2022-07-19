/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// VoucherStatus indicate the domain record data fields.
type VoucherStatus interface {
	VoucherID() [16]byte    //
	Status() Voucher_Status //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type VoucherStatus_StorageServices interface {
	Save(vs VoucherStatus) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vs VoucherStatus, nv protocol.NumberOfVersion, err protocol.Error)

	FilterByStatus(status Voucher_Status, offset, limit uint64) (voucherIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type Voucher_Status Quiddity_Status

const (
	Voucher_Status_AutoGenerated = Voucher_Status(Quiddity_Status_FreeFlag << iota)
	Voucher_Status_FirstInvoice

	Voucher_Status_PartlyUsed
	Voucher_Status_FullyUsed

	Voucher_Status_Expired
	Voucher_Status_Revoked
)


type (
	VoucherStatus_Service_Register_Request interface {
		VoucherID() [16]byte    
  	Status() Voucher_Status 
	}
	
	VoucherStatus_Service_Register_Response =   protocol.NumberOfVersion

)

type (
	VoucherStatus_Service_Count_Request interface {
		VoucherID() [16]byte
	}
	
	VoucherStatus_Service_Count_Response1 = protocol.NumberOfVersion

)


type (
	VoucherStatus_Service_Get_Request interface {
		VoucherID() [16]byte 
		VersionOffset() uint64
	}
	
	VoucherStatus_Service_Get_Response1 = protocol.NumberOfVersion
	VoucherStatus_Service_Get_Response2 = VoucherStatus
	
)


type (
	VoucherStatus_Service_FilterByStatus_Request interface {
		Voucher_Status() Quiddity_Status
		offset() uint64
		limit() uint64
	}
	
	VoucherStatus_Service_FilterByStatus_Response1 interface {
		VoucherIDs() [][16]byte
	}

	VoucherStatus_Service_FilterByStatus_Response2 = protocol.NumberOfVersion
)
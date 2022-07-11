package protocol

import "../libgo/protocol"

// VoucherDuration indicate the domain record data fields.
type VoucherDuration interface {
	VoucherID() [16]byte          // voucher domain
	Each() uint8                  // Each time use
	Epoch() VoucherDuration_Epoch //
	Duration() protocol.Duration  // from Epoch()
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type VoucherDuration_StorageServices interface {
	Save(vd VoucherDuration) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vd VoucherDuration, numbers uint64, err protocol.Error)
}

type VoucherDuration_Epoch uint8

const (
	VoucherDuration_Epoch_Unset VoucherDuration_Epoch = iota
	VoucherDuration_Epoch_IssueDate
	VoucherDuration_Epoch_LastUse
)

type (
	VoucherDuration_Service_Register_Request interface{
		VoucherID() [16]byte          
	  Each() uint8                  
	  Epoch() VoucherDuration_Epoch 
  	Duration() protocol.Duration  
	}
	
	VoucherDuration_Service_Register_Response interface{
		Nv() protocol.NumberOfVersion
	}
	
	VoucherDuration_Service_Count_Request interface{
		VoucherID() [16]byte
	}
	
	VoucherDuration_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	VoucherDuration_Service_Get_Request interface{
		VoucherID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherDuration_Service_Get_Response interface{
		VoucherDuration
		Nv() protocol.NumberOfVersion
	}
	
)
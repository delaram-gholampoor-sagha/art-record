package protocol

import (
	"../libgo/protocol"
)

// PosFinancialAccount or pos operator is to supply a particular service.
type PosFinancialAccount interface {
	PosID() [16]byte     // pos domain
	AccountID() [16]byte // financial-account domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}


type PosFinancialAccount_StorageServices interface {
	Save(ps PosFinancialAccount) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(posID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(posID [16]byte, vo protocol.VersionOffset) (ps PosFinancialAccount, nv protocol.NumberOfVersion, err protocol.Error)
	
}


type (
	PosFinancialAccount_Service_Register_Request interface {
		PosID() [16]byte
		AccountID() [16]byte
	}
	PosFinancialAccount_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	PosFinancialAccount_Service_Count_Request interface {
		PosID() [16]byte
	
	}
	PosFinancialAccount_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	
	}
	PosFinancialAccount_Service_Get_Request interface {
		PosID() [16]byte    
		Vo() protocol.VersionOffset
	}
	PosFinancialAccount_Service_Get_Response interface {
		PosFinancialAccount
		Nv() protocol.NumberOfVersion
	
	}
)
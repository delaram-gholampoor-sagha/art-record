package protocol

import "../libgo/protocol"
type FinancialAccountIBAN interface {
	AccountID() [16]byte // financial-account domain
	IBAN() iso.IBAN      //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type FinancialAccountIBAN_StorageServices interface {
	Save(fai FinancialAccountIBAN) (err protocol.Error)

	Count(accountID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fai FinancialAccountIBAN, err protocol.Error)
	

	FindByIBAN(iban iso.IBAN) (accountID [16]byte, err protocol.Error)
}

type (
	FinancialAccountIBAN_Service_Register_Request interface {
		AccountID() [16]byte
		IBAN() iso.IBAN
	}
	
	FinancialAccountIBAN_Service_Register_Response interface {
		Numbers() uint32
	}
	
	FinancialAccountIBAN_Service_Get_Request interface {
		AccountID() [16]byte
		VersionOffset() uint64
	}
	FinancialAccountIBAN_Service_Get_Response interface {
		FinancialAccountIBAN
	}
	
	FinancialAccountIBAN_Service_Count_Request interface {
		AccountID() [16]byte
	
	}
	
	FinancialAccountIBAN_Service_Count_Response interface {
		Numbers() uint32
	}
	
	FinancialAccountIBAN_Service_GetLast_Request interface {
		AccountID() [16]byte
	}
	FinancialAccountIBAN_Service_GetLast_Response interface {
		FinancialAccountIBAN
	}
	
	FinancialAccountIBAN_Service_FindByIBAN_Request interface {
		Iban() iso.IBAN
	}

	FinancialAccountIBAN_Service_FindByIBAN_Response interface {
		AccountID() [16]byte
	}
	
)
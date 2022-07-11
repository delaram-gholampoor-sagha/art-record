package protocol

import "../libgo/protocol"
type FinancialAccountCard interface {
	AccountID() [16]byte       // financial-account domain
	CardNumber() iso.Card      //
	ExpireDate() protocol.Time //
	CVC() uint16               //
	Time() protocol.Time       // save time
	RequestID() [16]byte       // user-request domain
}

type FinancialAccountCard_StorageServices interface {
	Save(fai FinancialAccountCard) (err protocol.Error)

	Count(accountID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fai FinancialAccountCard, err protocol.Error)
	

	FindByCardNumber(cardNumber iso.Card) (accountID [16]byte, err protocol.Error)
}

type (
	FinancialAccountCard_Service_Register_Request interface {
		AccountID() [16]byte
		CardNumber() iso.Card
		ExpireDate() protocol.Time
		CVC() uint16
	}
	
	FinancialAccountCard_Service_Register_Response interface {
		Numbers() uint32
	}
	
	FinancialAccountCard_Service_GetLast_Request interface {
		AccountID() [16]byte
	}
	FinancialAccountCard_Service_GetLast_Response interface {
		FinancialAccountCard
	}
	
	
	FinancialAccountCard_Service_Get_Request interface {
		AccountID() [16]byte
		VersionOffset() uint64
	}
	FinancialAccountCard_Service_Get_Response interface {
		FinancialAccountCard
	}
	
	
	FinancialAccountCard_Service_Count_Request interface {
		AccountID() [16]byte
	
	}
	FinancialAccountCard_Service_Count_Response interface {
		Numbers() uint32
	}
	
	FinancialAccountCard_Service_FindByCardNumber_Request interface {
		CardNumber() iso.Card
	
	}
	FinancialAccountCard_Service_FindByCardNumber_Response interface {
		AccountID() [16]byte
	}
)
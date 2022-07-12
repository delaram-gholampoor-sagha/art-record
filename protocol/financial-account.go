package protocol

import "../libgo/protocol"
type FinancialAccount interface {
	AccountID() [16]byte                // quiddity domain?
	Currency() [16]byte                 // financial-currency
	UserID() [16]byte                   // user domain
	MoneySettlementReference() [16]byte // user domain. can be any org or society user.
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}


type FinancialAccount_StorageServices interface {
	Save(fa FinancialAccount) (err protocol.Error)

	Count(accountID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fa FinancialAccount, err protocol.Error)
	

	FindByUserID(userID [16]byte, offset, limit uint64) (accountIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	ListUserCurrencies(userID [16]byte, offset, limit uint64) (currency [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListUserSettlementReferences(userID [16]byte, offset, limit uint64) (moneySettlementReferences [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	FinancialAccount_Service_Register_Request interface {          
		Currency() [16]byte                 
		UserID() [16]byte                   
		MoneySettlementReference() [16]byte      
	}
	
	
	FinancialAccount_Service_Register_Response interface {
		AccountID() [16]byte
		 Numbers() uint32
	}
	
	
	FinancialAccount_Service_Count_Request interface {
		ReferenceID() [16]byte
	
	}
	FinancialAccount_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
	
	
	FinancialAccount_Service_Get_Request interface {
		ReferenceID() [16]byte
		VersionOffset() uint64
	}
	
	FinancialAccount_Service_Get_Response interface {
		FinancialAccount
	}
	
	
	FinancialAccount_Service_FindByUserID_Request interface {
		UserID() [16]byte
		Offset() uint64 
		Limit() uint64
	}
	
	FinancialAccount_Service_FindByUserID_Response interface {
		AccountIDs() [][16]byte
		 NumberOfVersion() protocol.NumberOfVersion
	}
	
	FinancialAccount_Service_ListUserCurrencies_Request interface {
		UserID() [16]byte
		Offset() uint64 
		Limit() uint64
	}
	
	FinancialAccount_Service_ListUserCurrencies_Response interface {
		Currency() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
	
	FinancialAccount_Service_ListUserSettlementReferences_Request interface {
		UserID() [16]byte
		Offset() uint64 
		Limit() uint64
	}
	
	FinancialAccount_Service_ListUserSettlementReferences_Response interface {
		MoneySettlementReferences() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
)
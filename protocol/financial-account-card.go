/* For license and copyright information please see LEGAL file in repository */

package art

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
	Get(accountID [16]byte, versionOffset uint64) (fai FinancialAccountCard,nv protocol.NumberOfVersion , err protocol.Error)
	

	FindByCardNumber(cardNumber iso.Card) (accountID [16]byte,nv protocol.NumberOfVersion ,err protocol.Error)
}

type (
	FinancialAccountCard_Service_Register_Request interface {
		AccountID() [16]byte
		CardNumber() iso.Card
		ExpireDate() protocol.Time
		CVC() uint16
	}
	
	FinancialAccountCard_Service_Register_Response = protocol.NumberOfVersion

)



type (
	FinancialAccountCard_Service_Get_Request interface {
		AccountID() [16]byte
		versionOffset() uint64
	}

	FinancialAccountCard_Service_Get_Response1 = FinancialAccountCard
	FinancialAccountCard_Service_Get_Response2 = protocol.NumberOfVersion
	
)


type (
	FinancialAccountCard_Service_Count_Request interface {
		AccountID() [16]byte
	}

	FinancialAccountCard_Service_Count_Response = protocol.NumberOfVersion
)



type (	
	FinancialAccountCard_Service_FindByCardNumber_Request interface {
		CardNumber() iso.Card
	
	}
	FinancialAccountCard_Service_FindByCardNumber_Response1 interface {
		AccountID() [16]byte
	}

	FinancialAccountCard_Service_FindByCardNumber_Response2 = protocol.NumberOfVersion
)
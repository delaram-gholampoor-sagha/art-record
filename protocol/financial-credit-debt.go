/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// https://en.wikipedia.org/wiki/Credit_card_debt
type FinancialCreditDebt interface {
	UserID() [16]byte                // user domain
	Currency() [16]byte              // financial-currency
	Amount() protocol.AmountOfMoney  //
	Balance() protocol.AmountOfMoney //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}

// if long overdraft occur, user must pay fee. also can set preference to let system sell its shareholders to overcome of overdraft.

type FinancialCreditDebt_StorageServices interface {
	Save(ftr FinancialCreditDebt) (err protocol.Error)

	Count(referenceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(referenceID [16]byte, versionOffset uint64) (ftr FinancialCreditDebt,nv protocol.NumberOfVersion ,err protocol.Error)

}

type (
	FinancialCreditDebt_Service_Register_Request interface {
		UserID() [16]byte                
		Currency() [16]byte  
		Amount() protocol.AmountOfMoney  
		Balance() protocol.AmountOfMoney              
	}
	
	
	FinancialCreditDebt_Service_Register_Response = protocol.NumberOfVersion

)

type (
	FinancialCreditDebt_Service_Count_Request interface {
		ReferenceID() [16]byte
	}

	FinancialCreditDebt_Service_Count_Response = protocol.NumberOfVersion

)

type (
	FinancialCreditDebt_Service_Get_Request interface {
		ReferenceID() [16]byte
		versionOffset() uint64
	}
	
	FinancialCreditDebt_Service_Get_Response1 = FinancialCreditDebt
	FinancialCreditDebt_Service_Get_Response2 = protocol.NumberOfVersion
)
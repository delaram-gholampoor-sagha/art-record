/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// FinancialCredit indicate the domain record data fields.
type FinancialCredit interface {
	UserID() [16]byte                // user domain
	Currency() [16]byte              // financial-currency
	Amount() protocol.AmountOfMoney  //
	Balance() protocol.AmountOfMoney //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}


type FinancialCredit_StorageServices interface {
	Save(ftr FinancialCredit) (err protocol.Error)

	Count(referenceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(referenceID [16]byte, versionOffset uint64) (ftr FinancialCredit,nv protocol.NumberOfVersion ,err protocol.Error)

}

type (
	FinancialCredit_Service_Register_Request interface {
		UserID() [16]byte                
		Currency() [16]byte  
		Amount() protocol.AmountOfMoney  
		Balance() protocol.AmountOfMoney              
	}
  
  
	FinancialCredit_Service_Register_Response = protocol.NumberOfVersion

)

type (
	FinancialCredit_Service_Count_Request interface {
		ReferenceID() [16]byte
	
	}
	FinancialCredit_Service_Count_Response = protocol.NumberOfVersion
)

type (
	FinancialCredit_Service_Get_Request interface {
		ReferenceID() [16]byte
		versionOffset() uint64
	}
	
	FinancialCredit_Service_Get_Response1 = FinancialCredit
	FinancialCredit_Service_Get_Response2 = protocol.NumberOfVersion
  
)
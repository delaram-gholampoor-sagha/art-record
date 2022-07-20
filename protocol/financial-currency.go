/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// MoneySettlementReference: https://www.federalreserve.gov/paymentsystems/fedfunds_about.htm
type FinancialCurrency interface {
	Currency() [16]byte // quiddity domain
	// ISO_4217() iso.Currency             //
	MoneySettlementReference() [16]byte // user domain. can be any org or society user.
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}

type FinancialCurrency_StorageServices interface {
	Save(ftr FinancialCurrency) (err protocol.Error)

	Count(referenceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(referenceID [16]byte, versionOffset uint64) (ftr FinancialCurrency, nv protocol.NumberOfVersion ,err protocol.Error)

}

type (
	FinancialCurrency_Service_Register_Request interface {
		Currency() [16]byte            
		MoneySettlementReference() [16]byte   
	}
  
  
	FinancialCurrency_Service_Register_Response = protocol.NumberOfVersion

)

type (
	FinancialCurrency_Service_Count_Request interface {
		ReferenceID() [16]byte
	}
	FinancialCurrency_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
  
	FinancialCurrency_Service_Get_Request interface {
		ReferenceID() [16]byte
		versionOffset() uint64
	}
  
	FinancialCurrency_Service_Get_Response1 = 	FinancialCurrency
	FinancialCurrency_Service_Get_Response2 = protocol.NumberOfVersion
  
)
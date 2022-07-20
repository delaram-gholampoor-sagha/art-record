/* For license and copyright information please see LEGAL file in repository */

package art

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
	Get(accountID [16]byte, versionOffset uint64) (fai FinancialAccountIBAN,nv protocol.NumberOfVersion ,err protocol.Error)
	

	FindByIBAN(iban iso.IBAN) (accountID [16]byte,nv protocol.NumberOfVersion ,err protocol.Error)
}



type (
	FinancialAccountIBAN_Service_Register_Request interface {
		AccountID() [16]byte
		IBAN() iso.IBAN
	}
	
	FinancialAccountIBAN_Service_Register_Response = protocol.NumberOfVersion

)

type (

	FinancialAccountIBAN_Service_Get_Request interface {
		AccountID() [16]byte
		versionOffset() uint64
	}

	FinancialAccountIBAN_Service_Get_Response1 = FinancialAccountIBAN
	FinancialAccountIBAN_Service_Get_Response2 = protocol.NumberOfVersion
)

type (
	FinancialAccountIBAN_Service_Count_Request interface {
		AccountID() [16]byte
	
	}
	
	FinancialAccountIBAN_Service_Count_Response = protocol.NumberOfVersion
)



type (
	FinancialAccountIBAN_Service_FindByIBAN_Request interface {
		Iban() iso.IBAN
	}

	FinancialAccountIBAN_Service_FindByIBAN_Response1 interface {
		AccountID() [16]byte
	}

	FinancialAccountIBAN_Service_FindByIBAN_Response2 = protocol.NumberOfVersion
	
)
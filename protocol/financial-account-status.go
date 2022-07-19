/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"
type FinancialAccountStatus interface {
	AccountID() [16]byte             // financial-account domain
	Status() FinancialAccount_Status //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}

type FinancialAccountStatus_StorageServices interface {
	Save(fas FinancialAccountStatus) (err protocol.Error)

	Count(accountID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fas FinancialAccountStatus, err protocol.Error)
	

	FilterByStatus(status FinancialAccount_Status, offset, limit uint64) (accountIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type FinancialAccount_Status Quiddity_Status

const (
	FinancialAccount_Status_Closed = FinancialAccount_Status(Quiddity_Status_FreeFlag << iota)
)

type (
		FinancialAccountStatus_Service_Register_Request interface {               
	  AccountID() [16]byte            
		Status() FinancialAccount_Status    
	}
	
	
	FinancialAccountStatus_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
		FinancialAccountStatus_Service_Count_Request interface {
		AccountID() [16]byte
	
	}
	FinancialAccountStatus_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)


type (
		FinancialAccountStatus_Service_Get_Request interface {
		AccountID() [16]byte
		VersionOffset() uint64
	}
	
	FinancialAccountStatus_Service_Get_Response interface {
		FinancialAccountStatus
	}
	
)


type (
	FinancialAccountStatus_Service_FilterByStatus_Request interface {
		FinancialAccount_Status()
		Offset() uint64 
		Limit() uint64
	}
	
	FinancialAccountStatus_Service_FilterByStatus_Response interface {
		AccountIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
	
	
)
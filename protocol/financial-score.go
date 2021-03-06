/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// Individual and General score
type FinancialScore interface {
	UserID() [16]byte    // user domain
	Score() int64        //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}



type FinancialScore_StorageServices interface {
	Save(ftr FinancialScore) (err protocol.Error)

	Count(referenceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(referenceID [16]byte, versionOffset uint64) (ftr FinancialScore, nv protocol.NumberOfVersion ,err protocol.Error)

}

type (
	FinancialScore_Service_Register_Request interface {
		UserID() [16]byte    
		Score() int64       
	}
	
	
	FinancialScore_Service_Register_Response = protocol.NumberOfVersion

)

type (
	FinancialScore_Service_Count_Request interface {
		ReferenceID() [16]byte
	}

	FinancialScore_Service_Count_Response = protocol.NumberOfVersion
	
)

type (
	FinancialScore_Service_Get_Request interface {
		ReferenceID() [16]byte
		versionOffset() uint64
	}
	
	FinancialScore_Service_Get_Response1 = FinancialScore
	FinancialScore_Service_Get_Response2 = protocol.NumberOfVersion
	  
)
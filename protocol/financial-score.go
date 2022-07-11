package protocol

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

	Count(referenceID [16]byte) (numbers uint64, err protocol.Error)
	Get(referenceID [16]byte, versionOffset uint64) (ftr FinancialScore, err protocol.Error)

}

type (
	FinancialScore_Service_Register_Request interface {
		UserID() [16]byte    
		Score() int64       
	}
	
	
	FinancialScore_Service_Register_Response interface {
		Numbers() uint32
	}
	
	
	FinancialScore_Service_Count_Request interface {
		ReferenceID() [16]byte
	
	}
	FinancialScore_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	
	FinancialScore_Service_Get_Request interface {
		ReferenceID() [16]byte
		VersionOffset() uint64
	}
	
	FinancialScore_Service_Get_Response interface {
		FinancialScore
	}
	  
)
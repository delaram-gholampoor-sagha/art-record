package protocol

import "../libgo/protocol"

// FinancialTransactionOrder indicate the domain record data fields.
type FinancialTransactionOrder interface {
	ID() [16]byte          // FinancialTransactionOrderID
	UserID() [16]byte      // user domain
	ReferenceID() [16]byte // Other Platform transaction ID e.g. bank
	PlatformID() [4]byte   // Detect to||from Currency
	Amount() float64       //
	SuggestPrice() float64 //
	Time() protocol.Time   // save time
	RequestID() [16]byte   // user-request domain
}

type FinancialTransactionOrder_StorageServices interface {
	Save(ftr FinancialTransactionOrder) (err protocol.Error)

	Count(referenceID [16]byte) (numbers uint64, err protocol.Error)
	Get(referenceID [16]byte, versionOffset uint64) (ftr FinancialTransactionOrder, err protocol.Error)

}

type (
	FinancialTransactionOrder_Service_Register_Request interface {
		UserID() [16]byte      
		ReferenceID() [16]byte 
		PlatformID() [4]byte   
		Amount() float64      
		SuggestPrice() float64 
	}

  
	FinancialTransactionOrder_Service_Register_Response interface {
		ID() [16]byte 
		Numbers() uint32
	}
		
	
	FinancialTransactionOrder_Service_Count_Request interface {
		ReferenceID() [16]byte
	
	}
	FinancialTransactionOrder_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	
	FinancialTransactionOrder_Service_Get_Request interface {
		ReferenceID() [16]byte
		VersionOffset() uint64
	}
	
	FinancialTransactionOrder_Service_Get_Response interface {
		FinancialTransactionOrder
	}
  
)
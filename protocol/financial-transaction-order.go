/* For license and copyright information please see LEGAL file in repository */

package art

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

	Count(referenceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(referenceID [16]byte, versionOffset uint64) (ftr FinancialTransactionOrder, nv protocol.NumberOfVersion  ,  err protocol.Error)

}

type (
	FinancialTransactionOrder_Service_Register_Request interface {
		UserID() [16]byte      
		ReferenceID() [16]byte 
		PlatformID() [4]byte   
		Amount() float64      
		SuggestPrice() float64 
	}
	FinancialTransactionOrder_Service_Register_Response1 interface {
		ID() [16]byte 
	}

	FinancialTransactionOrder_Service_Register_Response2 = protocol.NumberOfVersion

)


type (
	FinancialTransactionOrder_Service_Count_Request interface {
		ReferenceID() [16]byte
	}

	FinancialTransactionOrder_Service_Count_Response = protocol.NumberOfVersion
	
)



type (
	FinancialTransactionOrder_Service_Get_Request interface {
		ReferenceID() [16]byte
		versionOffset() uint64
	}
	
	FinancialTransactionOrder_Service_Get_Response1 = FinancialTransactionOrder
	FinancialTransactionOrder_Service_Get_Response2 = protocol.NumberOfVersion
  
)
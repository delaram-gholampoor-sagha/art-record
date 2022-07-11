package protocol

import "../libgo/protocol"

// FinancialTransactionReference indicate the domain record data fields.
// FinancialTransactionReference store the secondary index that complicated than regular indexes.
// Use version to save multi transactions for the same reference.
type FinancialTransactionReference interface {
	ReferenceID() [16]byte         // any domain base on FinancialTransaction_Type
	SenderAccountID() [16]byte     // financial-account domain
	SenderAccountOffset() uint64   //
	ReceiverAccountID() [16]byte   // financial-account domain. or AccountPartyID
	ReceiverAccountOffset() uint64 //
	Time() protocol.Time           // save time
}

type FinancialTransactionReference_StorageServices interface {
	Save(ftr FinancialTransactionReference) (err protocol.Error)

	Count(referenceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(referenceID [16]byte, versionOffset uint64) (ftr FinancialTransactionReference, err protocol.Error)

}


type (
	FinancialTransactionReference_Service_Register_Request interface {
		SenderAccountID() [16]byte     
		SenderAccountOffset() uint64   
		ReceiverAccountID() [16]byte   
		ReceiverAccountOffset() uint64 
	}
	
	
	FinancialTransactionReference_Service_Register_Response interface {
	   ReferenceID() [16]byte
	}
	
	
	FinancialTransactionReference_Service_Count_Request interface {
		ReferenceID() [16]byte
	
	}
	FinancialTransactionReference_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	
	FinancialTransactionReference_Service_Get_Request interface {
		ReferenceID() [16]byte
		VersionOffset() uint64
	}
	
	FinancialTransactionReference_Service_Get_Response interface {
		FinancialTransactionReference
	}
	
)
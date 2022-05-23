

package protocol



// FinancialTransactionReference store the secondary index that complex than regular indexes.
// Use version to save multi transactions for a reference.
type FinancialTransactionReference interface {
	ReferenceID() [16]byte         // any domain base on FinancialTransaction_Type
	SenderAccountID() [16]byte     // financial-account || financial-bank-account domain
	SenderAccountOffset() uint64   //
	ReceiverAccountID() [16]byte   // financial-account || financial-bank-account domain. or AccountPartyID
	ReceiverAccountOffset() uint64 //
}

type FinancialTransactionReference_StorageServices interface {
	Save(ftr FinancialTransactionReference) (err error)

	Count(referenceID [16]byte) (numbers uint64, err error)
	Get(referenceID [16]byte, versionOffset uint64) (ftr FinancialTransactionReference, err error)
	Last(referenceID [16]byte) (ftr FinancialTransactionReference, err error)
}

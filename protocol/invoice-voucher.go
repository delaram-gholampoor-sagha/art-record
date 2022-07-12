package protocol

import "../libgo/protocol"

// InvoiceVoucher indicate the domain record data fields.
// This domain will let users redeem their vouchers and track vouchers rules.
type InvoiceVoucher interface {
	InvoiceID() [16]byte                // invoice domain
	VoucherID() [16]byte                // voucher domain
	Discounted() protocol.AmountOfMoney //
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}

type InvoiceVoucher_StorageServices interface {
	Save(iv InvoiceVoucher) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(invoiceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iv InvoiceVoucher, nv protocol.NumberOfVersion, err protocol.Error)

	FindByVoucherID(voucherID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	// protocol.EventTarget
}

type (
		InvoiceVoucher_Service_Register_Request interface {
		InvoiceID() [16]byte               
		VoucherID() [16]byte               
		Discounted() protocol.AmountOfMoney 
	}
	InvoiceVoucher_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	
	}

)

type (
		InvoiceVoucher_Service_Count_Request interface {
		InvoiceID() [16]byte
		
	
	}
	InvoiceVoucher_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)


type (
		InvoiceVoucher_Service_Get_Request interface {
		InvoiceID() [16]byte    
		VersionOffset() uint64
	}
	InvoiceVoucher_Service_Get_Response interface {
		InvoiceVoucher
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (

	InvoiceVoucher_Service_FindByVoucherID_Request interface {
		VoucherID() [16]byte
		Offset() uint64
		Limit() uint64	  
		
	}
	InvoiceVoucher_Service_FindByVoucherID_Response interface {
		InvoiceIDs() [][16]byte  
		NumberOfVersion() protocol.NumberOfVersion
	}
)
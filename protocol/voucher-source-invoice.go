package protocol

import "../libgo/protocol"

// VoucherSourceInvoice indicate the domain record data fields.
type VoucherSourceInvoice interface {
	VoucherID() [16]byte                    // voucher domain
	InvoiceID() [16]byte                    // invoice domain
	InvoiceType() VoucherSourceInvoice_Type //
	Time() protocol.Time                    // save time
	RequestID() [16]byte                    // user-request domain
}

type VoucherSourceInvoice_StorageServices interface {
	Save(vsi VoucherSourceInvoice) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vsi VoucherSourceInvoice, nv protocol.NumberOfVersion, err protocol.Error)

	FindByInvoice(invoiceID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type VoucherSourceInvoice_Type uint8

const (
	VoucherSourceInvoice_Type_Unset VoucherSourceInvoice_Type = iota
	VoucherSourceInvoice_Type_PastCompleted
	VoucherSourceInvoice_Type_HurryToBuying
	VoucherSourceInvoice_Type_BackToBuy
)


type (

		VoucherSourceInvoice_Service_Register_Request interface{
		VoucherID() [16]byte                    
	  InvoiceID() [16]byte                    
	  InvoiceType() VoucherSourceInvoice_Type 
	}
	
	VoucherSourceInvoice_Service_Register_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}
)


type (
	VoucherSourceInvoice_Service_Count_Request interface{
		VoucherID() [16]byte
	}
	
	VoucherSourceInvoice_Service_Count_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
	
	VoucherSourceInvoice_Service_Get_Request interface{
		VoucherID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherSourceInvoice_Service_Get_Response interface{
		VoucherSourceInvoice
		NumberOfVersion() protocol.NumberOfVersion
	}
)

type(
	VoucherSourceInvoice_Service_FindByInvoice_Request interface{
		InvoiceID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	VoucherSourceInvoice_Service_FindByInvoice_Response interface{
		VoucherIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
)
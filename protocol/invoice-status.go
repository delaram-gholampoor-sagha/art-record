/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// InvoiceStatus indicate the domain record data fields.
type InvoiceStatus interface {
	InvoiceID() [16]byte    // invoice-status domain
	Status() Invoice_Status //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type InvoiceStatus_StorageServices interface {
	Save(is InvoiceStatus) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(invoiceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (is InvoiceStatus, nv protocol.NumberOfVersion, err protocol.Error)

	FindByStatus(status Invoice_Status, offset, limit uint64) (invoiceIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	protocol.EventTarget
}


type Invoice_Status Quiddity_Status

const (
	// cancelled : Indicates that you have cancelled the invoice. Though you can view the cancelled invoice,
	// the customer cannot view the invoice or make payments against the invoice.
	// You can only cancel an unpaid issued invoice.
	// invoice is no longer valid, and not payable by client or refund
	Invoice_Status_Voided = Invoice_Status(Quiddity_Status_FreeFlag << iota)

	// expired : Indicates that an invoice has expired. Once expired,
	// the customer cannot make any payments against the invoice,
	// and you cannot cancel or delete the invoice. You can view an expired invoice.
	// Invoice_Status_Expired

	// partially paid : Indicates that the customer has made a partial payment against an invoice.
	// After a payment has been made, you can neither delete nor cancel the invoice.
	// You can only view the invoice and add internal notes to it.
	Invoice_Status_PartiallyPaid

	// paid : Indicates that the customer has fully paid the invoice amount.
	// Paid is the final state of an invoice. A paid invoice can only be viewed.
	// You can neither delete nor cancel it.
	Invoice_Status_Paid

	Invoice_Status_UserCheckout
	Invoice_Status_PendingReview
	Invoice_Status_Confirmed // Approved
	Invoice_Status_Packaging
	Invoice_Status_Shipping
	Invoice_Status_Done // delivered or served
)


type (
		InvoiceStatus_Service_Register_Request interface { 
		Status() Invoice_Status    
	}
	InvoiceStatus_Service_Register_Response1 interface {
		InvoiceID() [16]byte 
	}

	InvoiceStatus_Service_Register_Response2 = protocol.NumberOfVersion

)

type (
	InvoiceStatus_Service_Count_Request interface {
		InvoiceID() [16]byte
	}

	InvoiceStatus_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	InvoiceStatus_Service_Get_Request interface { 
		InvoiceID() [16]byte 
		versionOffset() uint64
	}
	InvoiceStatus_Service_Get_Response1 = InvoiceStatus
	InvoiceStatus_Service_Get_Response2 = protocol.NumberOfVersion
	
)


type (
	InvoiceStatus_Service_FindByStatus_Request interface { 
		Status() Invoice_Status
		Offset() uint64
		Limit() uint64
	}
	InvoiceStatus_Service_FindByStatus_Response1 interface {
		InvoiceIDs() [][16]byte 
	}

	InvoiceStatus_Service_FindByStatus_Response2 = protocol.NumberOfVersion
)


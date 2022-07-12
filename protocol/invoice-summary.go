package protocol

import "../libgo/protocol"

// InvoiceSummary indicate the domain record data fields.
// InvoiceSummary is summary data that calculated to reduce fetch and calculate on each client device.
type InvoiceSummary interface {
	InvoiceID() [16]byte                     // invoice domain
	Quantity() uint64                        // Total product count
	SuggestedPrice() protocol.AmountOfMoney  // Total should pay before any discount or auction
	DiscountedPrice() protocol.AmountOfMoney // user profit from this purchase by vouchers+auctions
	VoucherPrice() protocol.AmountOfMoney    // the sum of all vouchers apply for this invoice
	PayablePrice() protocol.AmountOfMoney    // Total price user pay to organization
	VAT() protocol.AmountOfMoney             // Total VAT user pay to its country government
	Time() protocol.Time                     // save time
}

type InvoiceSummary_StorageServices interface {
	Save(is InvoiceSummary) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(invoiceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (is InvoiceSummary, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	InvoiceSummary_Service_Register_Request interface {
		InvoiceID() [16]byte                    
		Quantity() uint64                       
		SuggestedPrice() protocol.AmountOfMoney  
		DiscountedPrice() protocol.AmountOfMoney 
		VoucherPrice() protocol.AmountOfMoney   
		PayablePrice() protocol.AmountOfMoney    
		VAT() protocol.AmountOfMoney             
	}
	InvoiceSummary_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	
	}
	
	InvoiceSummary_Service_Count_Request interface {
		InvoiceID() [16]byte
	
	}
	InvoiceSummary_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	
	}
	
	InvoiceSummary_Service_Get_Request interface { 
		InvoiceID() [16]byte 
		VersionOffset() uint64
	}
	InvoiceSummary_Service_Get_Response interface {
		InvoiceSummary
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)
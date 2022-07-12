package protocol

import "../libgo/protocol"

// InvoiceItemSummary indicate the domain record data fields.
type InvoiceItemSummary interface {
	InvoiceID() [16]byte                     // invoice domain
	ProductID() [16]byte                     // product domain
	SuggestedPrice() protocol.AmountOfMoney  // Total should pay before any discount or auction
	DiscountedPrice() protocol.AmountOfMoney // user profit from this product purchase by auctions
	PayablePrice() protocol.AmountOfMoney    // Total price user pay to organization
	VAT() protocol.AmountOfMoney             // Total VAT user pay to its country government
	Time() protocol.Time                     // save time
}

type InvoiceItemSummary_StorageServices interface {
	Save(iis InvoiceItemSummary) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(invoiceID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iis InvoiceItemSummary, err protocol.Error)
}

type (
	InvoiceItemSummary_Service_Register_Request interface {
		InvoiceID() [16]byte                     
		ProductID() [16]byte                     
		SuggestedPrice() protocol.AmountOfMoney  
		DiscountedPrice() protocol.AmountOfMoney 
		PayablePrice() protocol.AmountOfMoney    
		VAT() protocol.AmountOfMoney              
	}
	InvoiceItemSummary_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	
	}
	
	InvoiceItemSummary_Service_Count_Request interface { 
		InvoiceID() [16]byte 
	
	}
	InvoiceItemSummary_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
	InvoiceItemSummary_Service_Get_Request interface { 
		InvoiceID() [16]byte
		VersionOffset() uint64
	
	
	}
	InvoiceItemSummary_Service_Get_Response interface {
		InvoiceItemSummary
	}
)
package protocol

import (
	"time"
)

type Invoice struct {
	ID [16]byte
	// User_ID [16]byte
	// Voucher_ID       [16]byte
	// Tax_ID           [16]byte
	// Product_Price_ID [16]byte
	// RefrenceID [16]byte
	RelatedID      [16]byte
	// maybe we could move this into invoice shopping cart ?
	// we have both of them ...
	Number_of_Products uint64
	Status             uint8
	Payment_Date       time.Time
	Expire_by          time.Time
	Issued_at          time.Time
	Cancelled_At       time.Time
}

type InvoiceServices interface {
	GetInvoice(invoiceID uint64) (Invoice, error)
	GetInvoices() ([]Invoice, error)
	GetInvoicesByAdmin(userId uint64, invoiceID uint64) (Invoice, error)
	RegisterInvoice(invoice Invoice) (Invoice, error)

	FindInvoiceByProductID(productID uint64) (Invoice, error)
	FindInvoiceByUserID(userID uint64) (Invoice, error)
	FindInvoiceByAcossiatedProductOwner(productID uint64) (Invoice, error)
	FindInvoiceByDateTime(time string) (Invoice, error)
}

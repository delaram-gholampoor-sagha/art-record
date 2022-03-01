package protocol

type Invoice_Content struct {
	ID              uint64 `json:"id"`
	Acossiated_user string `json:"acossiated_user"`
	Price           string `json:"price"`
	Created_at      string `json:"created_at"`
	Updated_at      string `json:"updated_at"`
	State           string `json:"state"`
}

type Invoice interface {
	// Invoice ??
	GetInvoice(invoiceID uint64) (Invoice_Content, error)
	GetInvoices() ([]Invoice_Content, error)
	GetInvoicesByAdmin(userId uint64, invoiceID uint64) (Invoice_Content, error)
	RegisterInvoice(invoice Invoice_Content) (Invoice_Content, error)

	// Find Invoice
	FindInvoiceByProductID(productID uint64) (Invoice_Content, error)
	FindInvoiceByUserID(userID uint64) (Invoice_Content, error)
	FindInvoiceByAcossiatedProductOwner(productID uint64) (Invoice_Content , error)
	FindInvoiceByDateTime(time string) (Invoice_Content, error)
	FindInvoiceByState(state string ) (Invoice_Content , error)
}

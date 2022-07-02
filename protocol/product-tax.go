package protocol


// ProductTax indicate the domain record data fields.
// Tax Labels
// Setting tax labels allows you to perform the following actions:
// Define up to two tax labels and associate each of them with a specific tax rate
// Associate multiple tax rates with the same tax labels
// Associate a blank tax label with tax exempt items (considered non-value-added tax (VAT))
type ProductTax interface {
	ProductID() [16]byte // product domain
	SocietyID() [16]byte // society domain
	VAT() math.PerMyriad // VAT as Value-Added-Tax is bad name for a robbery
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

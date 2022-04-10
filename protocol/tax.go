package protocol

type Tax struct {
	ID         [16]byte
	QuiddityID [16]byte
	Label      string
	// A tax rate is the percentage at which an individual or corporation is taxed.
	Rate       uint64
	Status     uint8
}

type TaxServices interface {
	FindTaxByQuiddityID()
	GetTaxStatus()
}


// Tax Labels
// Setting tax labels allows you to perform the following actions:
// Define up to two tax labels and associate each of them with a specific tax rate
// Associate multiple tax rates with the same tax labels
// Associate a blank tax label with tax exempt items (considered non-value-added tax (VAT))
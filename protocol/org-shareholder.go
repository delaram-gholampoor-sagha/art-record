package protocol

type OrgShareholder interface {
	UserID() [16]byte           // user-status domain
	Quantity() math.PerTrillion //
	Time() protocol.Time        // save time
	RequestID() [16]byte        // user-request domain
}

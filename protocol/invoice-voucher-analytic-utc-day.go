package protocol


// InvoiceVoucherAnalyticUtcDay indicate the domain record data fields.
type InvoiceVoucherAnalyticUtcDay interface {
	Day() utc.DayElapsed                    //
	Redeemed() uint64                       // Total
	RedeemedAmount() protocol.AmountOfMoney // Total
	Time() protocol.Time                    // save time
}

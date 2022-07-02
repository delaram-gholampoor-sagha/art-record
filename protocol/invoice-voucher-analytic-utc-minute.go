package protocol

// InvoiceVoucherAnalyticUtcMinute indicate the domain record data fields.
type InvoiceVoucherAnalyticUtcMinute interface {
	Minute() utc.MinuteElapsed              //
	Redeemed() uint64                       // Total
	RedeemedAmount() protocol.AmountOfMoney // Total
	Time() protocol.Time                    // save time
}

type InvoiceVoucherAnalyticUtcMinute_StorageServices interface {
	Save(iva InvoiceVoucherAnalyticUtcMinute) (numbers uint64, err protocol.Error)

	Count() (numbers uint64, err protocol.Error)
	Get(versionOffset uint64) (iva InvoiceVoucherAnalyticUtcMinute, numbers uint64, err protocol.Error)
}

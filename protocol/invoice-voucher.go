package protocol

import (
	"time"
)

type InvoiceVoucher interface {
	InvoiceID() [16]byte                // invoice domain
	VoucherID() [16]byte                // voucher domain
	Discounted() protocol.AmountOfMoney //
	Time() time.Time                    // Save time
	RequestID() [16]byte                // user-request domain
}

type InvoiceVoucher_StorageServices interface {
	Save(iv InvoiceVoucher) (err error)

	Count(invoiceID [16]byte) (numbers uint64, err error)
	Get(invoiceID [16]byte, versionOffset uint64) (iv InvoiceVoucher, err error)
	Last(invoiceID [16]byte) (iv InvoiceVoucher, numbers uint64, err error)

	FindByVoucherID(voucherID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err error)
}

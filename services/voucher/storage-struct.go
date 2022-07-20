/* For license and copyright information please see LEGAL file in repository */

package voucher

import (
		"../../libgo/protocol"
)


type voucher struct {
	voucherID [16]byte 
	voucher_type Voucher_Type  
	time utc.time 
	requestID [16]byte 
}


func (v *voucher) VoucherID() [16]byte  { return v.voucherID }
func (v *voucher) Type() Voucher_Type   { return v.voucher_type }
func (v *voucher) Time() protocol.Time  { return v.time }
func (v *voucher) RequestID() [16]byte  { return v.requestID }


// JSON codec
func (q *voucher) FromJSON(payload []byte) (err protocol.Error) {
	// TODO::: auto generate
	return
}
func (q *voucher) ToJSON(payload []byte) []byte {
	// TODO::: auto generate
	return nil
}
func (q *voucher) LenAsJSON() (ln int) {
	// TODO::: auto generate
	return
}

/* For license and copyright information please see LEGAL file in repository */

package voucher


type voucher struct {
	voucherID [16]byte 
	voucher_type Voucher_Type  
	time utc.time 
	requestID [16]byte 
}



type Voucher_Type uint32

const (
	Voucher_Type_Unset Voucher_Type = 0
	Voucher_Type_Area  Voucher_Type = (1 << iota)
	Voucher_Type_Campaign
	Voucher_Type_Category
	Voucher_Type_Duration
	Voucher_Type_Group
	Voucher_Type_Invoice
	Voucher_Type_Location
	Voucher_Type_POS
	Voucher_Type_Product
	Voucher_Type_SourceInvoice
	Voucher_Type_Usage    // limit to usage number
	Voucher_Type_UserType // limit to user type
	Voucher_Type_User     // limit to specific users
)


func (v *voucher) VoucherID() [16]byte { return v.voucherID }
func (v *voucher) Type() Voucher_Type   { return v.voucher_type }
func (v *voucher) Time() protocol.Time  { return v.time }
func (v *voucher) RequestID() [16]byte  { return v.requestID }


// JSON codec
func (q *Quiddity) FromJSON(payload []byte) (err protocol.Error) {
	// TODO::: auto generate
	return
}
func (q *Quiddity) ToJSON(payload []byte) []byte {
	// TODO::: auto generate
	return nil
}
func (q *Quiddity) LenAsJSON() (ln int) {
	// TODO::: auto generate
	return
}

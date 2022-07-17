package voucher






type CountRequest struct {
	voucherID [16]byte `json:"voucherID,string"`
}

func (req *CountRequest) VoucherID() [16]byte       { return req.voucherID}
func (req *CountRequest) Set_voucherID(vi [16]byte) { req.voucherID = vi }




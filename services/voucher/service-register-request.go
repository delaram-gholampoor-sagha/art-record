package voucher




type RegisterRequest struct {
	voucher_type Voucher_Type  `json:"Type,string"`
}

func (req *RegisterRequest) Voucher_Type() Voucher_Type      { return req.voucher_type}
func (req *RegisterRequest) Set_Voucher_Type(vt Voucher_Type) {  req.voucher_type = vt }





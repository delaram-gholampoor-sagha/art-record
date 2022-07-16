package voucher

import "github.com/Delaram-Gholampoor-Sagha/art-record/protocol"







type GetResponse struct {
	voucherID [16]byte `json:"Type,string"`
	nv protocol.NumberOfVersion `json:"Type,string"`
}

func (req *GetResponse) VoucherID() [16]byte      { return req.voucherID}
func (req *GetResponse) Set_VoucherID(v [16]byte) {  req.voucherID = v }

func (req *GetResponse) Nv() protocol.NumberOfVersion      { return req.nv}
func (req *GetResponse) Set_nv(nv protocol.NumberOfVersion) {  req.nv = nv }
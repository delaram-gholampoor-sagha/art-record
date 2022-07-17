package voucher

import "github.com/Delaram-Gholampoor-Sagha/art-record/protocol"



type RegisterResponse struct {
		voucherID [16]byte                       `json:"voucherID,string"`
		numberOfVersion protocol.NumberOfVersion `json:"protocol.NumberOfVersion,string"`
}


func (req *RegisterResponse) VoucherID() [16]byte     { return req.voucherID}
func (req *RegisterResponse) Set_voucherID(vi [16]byte) {  req.voucherID = vi }



func (req *RegisterResponse) NumberOfVersion() protocol.NumberOfVersion      { return req.numberOfVersion}
func (req *RegisterResponse) Set_NumberOfVersion(nv protocol.NumberOfVersion) {  req.numberOfVersion = nv }
package voucher

import "github.com/Delaram-Gholampoor-Sagha/art-record/protocol"




	


type GetResponse struct {
	voucher protocol.Voucher                 `json:"protocol.Voucher,string"`
	numberOfVersion protocol.NumberOfVersion `json:"protocol.NumberOfVersion,string"`
}

func (req *GetResponse) Voucher() [16]byte      { return req.voucher}
func (req *GetResponse) Set_Voucher(v [16]byte) {  req.voucher = v }

func (req *GetResponse) NumberOfVersion() protocol.NumberOfVersion      { return req.numberOfVersion}
func (req *GetResponse) Set_NumberOfVersion(nv protocol.NumberOfVersion) {  req.numberOfVersion = nv }
package voucher



import (
		"../../libgo/protocol"
)







type CountResponse struct {
	numberOfVersion protocol.NumberOfVersion `json:"numberOfVersion,string"`
}

func (req *CountResponse) NumberOfVersion() protocol.NumberOfVersion      { return req.numberOfVersion}
func (req *CountResponse) Set_numberOfVersion(v protocol.NumberOfVersion) { req.numberOfVersion = v }


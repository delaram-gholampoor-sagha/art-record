package voucher



import (
		"../../libgo/protocol"
)


	


type GetResponse struct {
	voucher protocol.Voucher                 `json:"protocol.Voucher,string"`
	numberOfVersion protocol.NumberOfVersion `json:"protocol.NumberOfVersion,string"`
}

func (req *GetResponse) Voucher() [16]byte      { return req.voucher}
func (req *GetResponse) Set_Voucher(v [16]byte) {  req.voucher = v }

func (req *GetResponse) NumberOfVersion() protocol.NumberOfVersion      { return req.numberOfVersion}
func (req *GetResponse) Set_NumberOfVersion(nv protocol.NumberOfVersion) {  req.numberOfVersion = nv }




// JSON codec
func (req *RegisterRequest) FromJSON(payload []byte) (err protocol.Error) {
	var decoder = json.DecoderUnsafeMinified{Buf: payload}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "Language":
			var num uint32
			num, err = decoder.DecodeUInt32()
			req.Language = protocol.Language(num)
		case "URI":
			req.URI, err = decoder.DecodeString()
		case "Title":
			req.Title, err = decoder.DecodeString()
		default:
			err = decoder.NotFoundKeyStrict()
		}

		if decoder.End() {
			return
		}
	}
	return
}

func (req *RegisterRequest) ToJSON(payload []byte) []byte {
	var encoder = json.Encoder{Buf: payload}

	encoder.EncodeString(`{"Language":`)
	encoder.EncodeUInt32(uint32(req.Language))

	encoder.EncodeString(`,"URI":"`)
	encoder.EncodeString(req.URI)

	encoder.EncodeString(`","Title":"`)
	encoder.EncodeString(req.Title)

	encoder.EncodeString(`"}`)
	return encoder.Buf
}

func (req *RegisterRequest) LenAsJSON() (ln int) {
	ln = len(req.URI) + len(req.Title)
	ln += 43
	return
}

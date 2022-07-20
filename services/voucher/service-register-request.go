package voucher

import (
		"../../libgo/protocol"
)


type RegisterRequest struct {
	voucher_type Voucher_Type  `json:"Type,string"`
}

func (req *RegisterRequest) Voucher_Type() Voucher_Type      { return req.voucher_type}
func (req *RegisterRequest) Set_Voucher_Type(vt Voucher_Type) {  req.voucher_type = vt }



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


package voucher



type GetRequest struct {
		voucherID [16]byte      `json:"voucherID,string"`
		versionOffset uint64    `json:"versionOffset,string"`
}

func (req *GetRequest) VoucherID() [16]byte        { return req.voucherID}
func (req *GetRequest) Set_voucherID(vi [16]byte)  {  req.voucherID = vi }

func (req *GetRequest) versionOffset() uint64      { return req.versionOffset}
func (req *GetRequest) Set_versionOffset(v uint64) {  req.versionOffset = v }


// JSON codec
func (req *GetRequest) FromJSON(payload []byte) (err protocol.Error) {
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
func (req *GetRequest) ToJSON(payload []byte) []byte {
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


func (req *GetRequest) LenAsJSON() (ln int) {
	ln = len(req.URI) + len(req.Title)
	ln += 43
	return
}
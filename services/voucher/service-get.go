package voucher

import (
	"../../libgo/protocol"
	"../../libgo/service"
	art "github.com/Delaram-Gholampoor-Sagha/art-record/protocol"
)


var GetService getService

type getService struct {
	service.Service
}


func (ser *registerService) ServeHTTP(st protocol.Stream, httpReq protocol.HTTPRequest, httpRes protocol.HTTPResponse) (err protocol.Error) {
	var req RegisterRequest
	err = req.FromJSON(httpReq.Body().Marshal())
	if err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	err = ser.Process(st, &req)
	if err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	httpRes.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
	return
}


func (ser *registerService) validateRequest(req art.Voucher_Service_Register_Request) (err protocol.Error) {
	// TODO::: auto generate
	return
}


func (ser *registerService) Process(st protocol.Stream, req art.Voucher_Service_Register_Request) (err protocol.Error) {
	err = st.Authorize()
	if err != nil {
		return
	}

	// if req.URI != "" {
	// 	err = checkQuiddityURI(st, req.URI)
	// 	if err != nil {
	// 		return
	// 	}
	// }


	var v = Voucher{
		VoucherID:  [16]byte 
	  Voucher_type: Voucher_Type  
	  Time: utc.time 
  	RequestID [16]byte 
	// 	AppInstanceID:    achaemenid.App.Nodes.LocalNode.InstanceID,
	// 	UserConnectionID: st.Connection.ID,
	// 	ID:               uuid.Random32Byte(),
	// 	OrgID:            st.Connection.UserID,

	// 	Language: req.Language,
	// 	URI:      req.URI,
	// 	Title:    req.Title,
	// 	Status:   data.QuiddityStatusRegister,
	// }
	err = storage.Save(&v)
	if err != nil {
		return
	}

	res = &registerVoucherRes{
		ID: q.ID,
	}

	return
}
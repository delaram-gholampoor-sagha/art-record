/* For license and copyright information please see LEGAL file in repository */

package voucher


import (
		"../../libgo/protocol"
		"../../libgo/service"
	
)

var RegisterService registerService

type registerService struct {
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
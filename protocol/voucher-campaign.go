package protocol

import "../libgo/protocol"

// VoucherCampaign indicate the domain record data fields.
type VoucherCampaign interface {
	VoucherID() [16]byte  // voucher domain
	CampaignID() [16]byte // marketing-campaign domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type VoucherCampaign_StorageServices interface {
	Save(vc VoucherCampaign) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vc VoucherCampaign, numbers uint64, err protocol.Error)

	FindByCampaign(campaignID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, numbers uint64, err protocol.Error)
}

type (
	VoucherCampaign_Service_Register_Request interface {
		VoucherID() [16]byte  
  	CampaignID() [16]byte 
	}
	
	VoucherCampaign_Service_Register_Response interface {
		Numbers() uint64
	}
	
	VoucherCampaign_Service_Count_Request interface {
		VoucherID() [16]byte
	}
	
	VoucherCampaign_Service_Count_Response interface {
		Numbers() uint64
	}
	
	VoucherCampaign_Service_Get_Request interface {
		VoucherID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherCampaign_Service_Get_Response interface {
		VoucherCampaign
		  Numbers() uint64
	}
	
	VoucherCampaign_Service_FindByCampaign_Request interface {
		CampaignID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	VoucherCampaign_Service_FindByCampaign_Response interface {
		VoucherIDs() [][16]byte
		Numbers() uint64
	}
)
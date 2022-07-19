/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// VoucherCampaign indicate the domain record data fields.
type VoucherCampaign interface {
	VoucherID() [16]byte  // voucher domain
	CampaignID() [16]byte // marketing-campaign domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type VoucherCampaign_StorageServices interface {
	Save(vc VoucherCampaign) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(voucherID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vc VoucherCampaign, nv protocol.NumberOfVersion, err protocol.Error)

	FindByCampaign(campaignID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}



type (

	VoucherCampaign_Service_Register_Request interface {
		VoucherID() [16]byte  
  	CampaignID() [16]byte 
	}
	
	VoucherCampaign_Service_Register_Response = protocol.NumberOfVersion
)

type (
	VoucherCampaign_Service_Count_Request interface {
		VoucherID() [16]byte
	}
	
	VoucherCampaign_Service_Count_Response = protocol.NumberOfVersion
)

type (
	VoucherCampaign_Service_Get_Request interface {
		VoucherID() [16]byte
		VersionOffset() uint64
	}
	
	VoucherCampaign_Service_Get_Response1 = VoucherCampaign
	VoucherCampaign_Service_Get_Response2 = protocol.NumberOfVersion

)

type (
	VoucherCampaign_Service_FindByCampaign_Request interface {
		CampaignID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	VoucherCampaign_Service_FindByCampaign_Response1 interface {
		VoucherIDs() [][16]byte
	}

	VoucherCampaign_Service_FindByCampaign_Response2 = protocol.NumberOfVersion
)
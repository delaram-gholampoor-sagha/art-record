package protocol

import "../libgo/protocol"

// ProductReserveValidity indicate the domain record data fields.
type ProductReserveValidity interface {
	ProductID() [16]byte         // product domain
	Duration() protocol.Duration // from add to invoice time
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type ProductReserveValidity_StorageServices interface {
	Save(prv ProductReserveValidity) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (prv ProductReserveValidity, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	ProductReserveValidity_Service_Register_Request interface {
		ProductID() [16]byte
		Duration() protocol.Duration
	}
	ProductReserveValidity_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
	ProductReserveValidity_Service_Count_Request interface {
		ProductID() [16]byte
	
	}
	ProductReserveValidity_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	
	}
	
)



type (
	ProductReserveValidity_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	ProductReserveValidity_Service_Get_Response interface {
		ProductReserveValidity
		NumberOfVersion() protocol.NumberOfVersion
	}
	
	
)
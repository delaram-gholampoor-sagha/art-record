package protocol

import "../libgo/protocol"

// ProductTimeValidity indicate the domain record data fields.
// flight departure and arrival time, concert start and end time, ...
type ProductTimeValidity interface {
	ProductID() [16]byte  // product domain
	Start() protocol.Time //
	End() protocol.Time   //
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type ProductTimeValidity_StorageServices interface {
	Save(pt ProductTimeValidity) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pt ProductTimeValidity, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	ProductTimeValidity_Service_Register_Request interface {
		ProductID() [16]byte  
	  Start() protocol.Time 
	  End() protocol.Time   
	}
	
	ProductTimeValidity_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	ProductTimeValidity_Service_Count_Request interface {
		ProductID() [16]byte
	}
	
	ProductTimeValidity_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	ProductTimeValidity_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	
	ProductTimeValidity_Service_Get_Response interface {
		ProductTimeValidity
		Nv() protocol.NumberOfVersion
	}
	
)
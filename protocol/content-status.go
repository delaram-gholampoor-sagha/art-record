package protocol

import "../libgo/protocol"

type ContentStatus interface {
	ContentID() [16]byte    // content domain
	Status() Content_Status //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ContentStatus_StorageServices interface {
	Save(cs ContentStatus) protocol.Error

	Count(contentID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(contentID [16]byte, versionOffset uint64) (cs ContentStatus, err protocol.Error)
	
	// FilterByStatus(status Content_Status, offset, limit uint64) (contentIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	// protocol.EventTarget
}

type Content_Status Quiddity_Status

const (
// Content_Status_ = Content_Status(Quiddity_Status_FreeFlag << iota)
)

type (
	ContentStatus_Service_Register_Request interface {
		ContentID() [16]byte    
  	Status() Content_Status 
	}

	ContentStatus_Service_Register_Response interface {
	 Nv() protocol.NumberOfVersion
	}
	
	
	
	ContentStatus_Service_Count_Request interface {
	ContentID() [16]byte
	}
	
	ContentStatus_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	ContentStatus_Service_Get_Request interface {
	 ContentID() [16]byte
	 VersionOffset() uint64
	}
	
	ContentStatus_Service_Get_Response interface {
		ContentStatus
	}
	
)
package protocol

import "../libgo/protocol"

// CommentStatus indicate the domain record data fields.
type CommentStatus interface {
	CommentID() [16]byte    // comment domain
	Status() Comment_Status //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type CommentStatus_StorageServices interface {
	Save(cs CommentStatus) protocol.Error

	Count(commentID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(commentID [16]byte, versionOffset uint64) (cs CommentStatus, err protocol.Error)
	

	// FilterByStatus(status Comment_Status, offset, limit uint64) (commentIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	protocol.EventTarget
}

type Comment_Status Quiddity_Status

const (
// Comment_Status_ = Comment_Status(Quiddity_Status_FreeFlag << iota)
)


type (
		CommentStatus_Service_Register_Request interface {
		CommentID() [16]byte    
  	Status() Comment_Status 
	}

		CommentStatus_Service_Register_Response interface {
	  NumberOfVersion() protocol.NumberOfVersion
	}


)

type (
	CommentStatus_Service_Count_Request interface {
	CommentID() [16]byte
	}

	CommentStatus_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)



type (
	CommentStatus_Service_Get_Request interface {
	CommentID() [16]byte
	VersionOffset() uint64
	}

	CommentStatus_Service_Get_Response interface {
		CommentStatus
	}
)
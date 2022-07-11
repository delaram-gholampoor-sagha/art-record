package protocol

import "../libgo/protocol"

// CommentPin indicate the domain record data fields.
type CommentPin interface {
	CommentID() [16]byte // comment domain
	GroupID() [16]byte   // group domain
	PinedID() [16]byte   // comment domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type CommentPin_StorageServices interface {
	Save(cp CommentPin) protocol.Error

	Get(commentID [16]byte) (cp CommentPin, err protocol.Error)

	FindByGroupID(groupID [16]byte, offset, limit uint64) (commentIDs [][16]byte, numbers uint64, err protocol.Error)
}

// Unpin service : Hidden status to unpin a comment
type (
	CommentPin_Service_Register_Request interface {
		CommentID() [16]byte 
		GroupID() [16]byte  
		PinedID() [16]byte   
	}

	CommentPin_Service_Register_Response interface {
		 Nv() protocol.NumberOfVersion  
	}
	
	CommentPin_Service_Get_Request interface {
		CommentID() [16]byte
	}
	
	CommentPin_Service_Get_Response interface {
		CommentPin
	}
	
	CommentPin_Service_FindByGroupID_Request interface {
		GroupID() [16]byte
		Offset()  uint64
		Limit() uint64
	}
	
	CommentPin_Service_FindByGroupID_Response interface {
		CommentIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
)
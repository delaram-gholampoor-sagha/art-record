package protocol

import "../libgo/protocol"

// CommentMentioned indicate the domain record data fields.
type CommentMentioned interface {
	UserID() [16]byte    // user domain
	CommentID() [16]byte // comment domain
	Time() protocol.Time // save time
}

type CommentMentioned_StorageServices interface {
	Save(cm CommentMentioned) protocol.Error

	Count(userID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (cm CommentMentioned, err protocol.Error)
	
}


type (
	CommentMentioned_Service_Register_Request interface {
		UserID() [16]byte    
   	CommentID() [16]byte 
	}

	CommentMentioned_Service_Register_Response interface {
  	Numbers() uint64
	}
	
	CommentMentioned_Service_Count_Request interface {
		UserID() [16]byte
	}
	
	CommentMentioned_Service_Count_Response interface {
		Numbers() uint64
	}
	
	CommentMentioned_Service_Get_Request interface {
		UserID() [16]byte 
		VersionOffset() uint64
	}
	
	CommentMentioned_Service_Get_Response interface {
		CommentMentioned
	}
	
)
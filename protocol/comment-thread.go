package protocol

import "../libgo/protocol"

// CommentThread indicate the domain record data fields.
// Comment thread name can be set by help of comment-text domain
type CommentThread interface {
	CommentID() [16]byte             // comment domain
	ArchiveAfter() protocol.Duration //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}

type CommentThread_StorageServices interface {
	Save(ct CommentThread) protocol.Error

	Get(commentID [16]byte) (ct CommentThread, err protocol.Error)
}


type (
	CommentThread_Service_Register_Request interface {
			CommentID() [16]byte             
    	ArchiveAfter() protocol.Duration 
	}

		CommentThread_Service_Register_Response interface {
      Nv() protocol.NumberOfVersion
	}
	
	
	CommentThread_Service_Get_Request interface {
		CommentID() uint64
	}
	
	CommentThread_Service_Get_Response interface {
		CommentThread
	}
)
/* For license and copyright information please see LEGAL file in repository */

package art

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
	Save(ct CommentThread) ( protocol.NumberOfVersion , protocol.Error)

	Get(commentID [16]byte) (ct CommentThread,nv  protocol.NumberOfVersion ,err protocol.Error)
}


type (
	CommentThread_Service_Register_Request interface {
			CommentID() [16]byte             
    	ArchiveAfter() protocol.Duration 
	}

	CommentThread_Service_Register_Response = protocol.NumberOfVersion

)

type (
	CommentThread_Service_Get_Request interface {
		CommentID() [16]byte 
	}
	
	CommentThread_Service_Get_Response1 = CommentThread
	CommentThread_Service_Get_Response2 = protocol.NumberOfVersion
)
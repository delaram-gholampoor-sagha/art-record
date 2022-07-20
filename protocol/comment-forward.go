/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// CommentForward indicate the domain record data fields.
type CommentForward interface {
	CommentID() [16]byte   // comment domain
	ForwardedID() [16]byte // comment domain. or ForwardedCommentID
	Time() protocol.Time   // save time
	RequestID() [16]byte   // user-request domain
}

type CommentForward_StorageServices interface {
	Save(cf CommentForward) ( protocol.NumberOfVersion, protocol.Error)

	Get(commentID [16]byte) (cf CommentForward, nv protocol.NumberOfVersion,,err protocol.Error)

}

type (
	CommentForward_Service_Register_Request interface {
		CommentID() [16]byte   
	  ForwardedID() [16]byte 
	}

	CommentForward_Service_Register_Response = protocol.NumberOfVersion

)

type (
	CommentForward_Service_Get_Request interface {
		CommentID() [16]byte
	}
	
	CommentForward_Service_Get_Response1 = CommentForward
	CommentForward_Service_Get_Response2 = protocol.NumberOfVersion
	
)
/* For license and copyright information please see LEGAL file in repository */

package art

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
	Save(cp CommentPin) (protocol.NumberOfVersion , protocol.Error)

	Get(commentID [16]byte) (cp CommentPin,nv protocol.NumberOfVersion ,err protocol.Error)

	FindByGroupID(groupID [16]byte, offset, limit uint64) (commentIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}
// Unpin service : Hidden status to unpin a comment

type (
	CommentPin_Service_Get_Request interface {
		CommentID() [16]byte
	}
	
	CommentPin_Service_Get_Response1 = CommentPin
  CommentPin_Service_Get_Response2 = protocol.NumberOfVersion
)



type (
	CommentPin_Service_FindByGroupID_Request interface {
		GroupID() [16]byte
		Offset()  uint64
		Limit() uint64
	}
	
	CommentPin_Service_FindByGroupID_Response1 interface {
		CommentIDs() [][16]byte
	}
	
	CommentPin_Service_FindByGroupID_Response2 = protocol.NumberOfVersion
)
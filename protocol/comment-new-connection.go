/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// CommentFollow indicate the domain record data fields.
type CommentFollow interface {
	CommentID() [16]byte      // comment domain
	FollowedUserID() [16]byte // user domain. FollowedUserID
	Type() CommentFollow_Type //
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type CommentFollow_Type uint8

// CommentFollow_Type_MakeFriend
// CommentFollow_Type_Follow


type CommentFollow_StorageServices interface {
	Save(co CommentFollow) (protocol.NumberOfVersion , protocol.Error)

	Count(commentID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(commentID [16]byte, versionOffset uint64) (co CommentFollow, nv protocol.NumberOfVersion,err protocol.Error)

}


type (
	CommentFollow_Service_Register_Request interface {
		CommentID() [16]byte 
	  ObjectID() [16]byte  
	}

	CommentFollow_Service_Register_Response = protocol.NumberOfVersion

)

type (
	CommentFollow_Service_Count_Request interface {
		CommentID() [16]byte
	}
	
	CommentFollow_Service_Count_Response = protocol.NumberOfVersion
	
)

type (
	CommentFollow_Service_Get_Request interface {
		CommentID() [16]byte
		versionOffset() uint64
	}
	
  CommentFollow_Service_Get_Response1 = CommentFollow
	CommentFollow_Service_Get_Response2 = protocol.NumberOfVersion

)


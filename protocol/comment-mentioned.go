/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// CommentMentioned indicate the domain record data fields.
type CommentMentioned interface {
	UserID() [16]byte    // user domain
	CommentID() [16]byte // comment domain
	Time() protocol.Time // save time
}

type CommentMentioned_StorageServices interface {
	Save(cm CommentMentioned) (protocol.NumberOfVersion , protocol.Error)

	Count(userID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (cm CommentMentioned,nv protocol.NumberOfVersion ,err protocol.Error)
	
}

type (
	CommentMentioned_Service_Register_Request interface {
		UserID() [16]byte    
   	CommentID() [16]byte 
	}

	CommentMentioned_Service_Register_Response =  protocol.NumberOfVersion

)


type (
	CommentMentioned_Service_Count_Request interface {
		UserID() [16]byte
	}
	
	CommentMentioned_Service_Count_Response = protocol.NumberOfVersion

)

type (
	CommentMentioned_Service_Get_Request interface {
		UserID() [16]byte 
		versionOffset() uint64
	}
	
	CommentMentioned_Service_Get_Response1 = CommentMentioned
	CommentMentioned_Service_Get_Response2 = protocol.NumberOfVersion
	
)
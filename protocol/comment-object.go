package protocol

import "../libgo/protocol"

// CommentObject indicate the domain record data fields.
// Object can be any such as image, video, voice, ...
type CommentObject interface {
	CommentID() [16]byte // comment domain
	ObjectID() [16]byte  // object domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type CommentObject_StorageServices interface {
	Save(co CommentObject) protocol.Error

	Count(commentID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(commentID [16]byte, versionOffset uint64) (co CommentObject, err protocol.Error)


	FindByObjectID(objectID [16]byte, offset, limit uint64) (commentIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	CommentObject_Service_Register_Request interface {
		CommentID() [16]byte 
	  ObjectID() [16]byte  
	}

	CommentObject_Service_Register_Response interface {
	  NumberOfVersion() protocol.NumberOfVersion
	}

)


type (
	CommentObject_Service_Count_Request interface {
		CommentID() [16]byte
	}
	
	CommentObject_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)


type (
	CommentObject_Service_Get_Request interface {
		CommentID() [16]byte
		VersionOffset() uint64
	}
	 CommentObject_Service_Get_Response interface {
		CommentObject
	}
	
)


type (
	CommentObject_Service_FindByObjectID_Request interface {
		ObjectID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	CommentObject_Service_FindByObjectID_Response interface {
		CommentIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
)
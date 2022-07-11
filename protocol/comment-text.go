package protocol

import "../libgo/protocol"

// CommentText indicate the domain record data fields.
type CommentText interface {
	CommentID() [16]byte // comment domain
	Text() string        // Text With Style. HTML & CSS (No JS) is more expressive than markdown, suggest use them in article text to style text.
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type CommentText_StorageServices interface {
	Save(c CommentText) protocol.Error

	Count(commentID [16]byte) (numbers uint64, err protocol.Error)
	Get(commentID [16]byte, versionOffset uint64) (c CommentText, err protocol.Error)
	

	FindByText(text string, offset, limit uint64) (commentIDs [16]byte, numbers uint64, err protocol.Error)
}

type (
		
	CommentText_Service_Register_Request interface {
		CommentID() [16]byte 
	  Text() string        
	}

	CommentText_Service_Register_Response interface {
	   Nv() protocol.NumberOfVersion        
	}


	CommentText_Service_Count_Request interface {
		CommentID() [16]byte
	}

	CommentText_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}

	CommentText_Service_Get_Request interface {
		CommentID() [16]byte
		VersionOffset() uint64
	}

	CommentText_Service_Get_Response interface {
		CommentText
	}

	CommentText_Service_FindByText_Request interface {
		Text() string
		Offset() uint64
		Limit() uint64
	}

	CommentText_Service_FindByText_Response interface {
		CommentIDs() [16]byte 
		Nv() protocol.NumberOfVersion
	}


)
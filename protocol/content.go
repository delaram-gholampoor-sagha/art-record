package protocol

import "../libgo/protocol"

// 3 type of content exist: 1:about persons, 2:about events, 3:about sciences
type Content interface {
	ContentID() [16]byte // quiddity domain
	UserID() [16]byte    // user domain
	Type() Content_Type  //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type Content_StorageServices interface {
	Save(c Content) protocol.Error

	Count(contentID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(contentID [16]byte, versionOffset uint64) (c Content, err protocol.Error)
	

	GetIDs(offset, limit uint64) (contentIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	// GetIDsByDateTime(time protocol.Time, offset, limit uint64) (contentIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	FindByUserID(userID [16]byte, offset, limit uint64) (contentIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type Content_Type uint16

const (
	Content_Type_Unset Content_Type = iota

	// Article use for many thing like NEWS, blog, education systems, features, ...
	Content_Type_Article
	Content_Type_BlogPost

	Content_Type_Book
	Content_Type_Course
	Content_Type_Tutorial

	Content_Type_Music
	Content_Type_Album
	Content_Type_Concert
	Content_Type_MusicVideo

	Content_Type_Movie
	Content_Type_Serial
	Content_Type_SerialEpisode

	Content_Type_Game
	Content_Type_SportGame
)

type (
	Content_Service_Register_Request interface {
  	UserID() [16]byte    
		Type() Content_Type  
	}

		Content_Service_Register_Response interface {
       ContentID() [16]byte 
			 Nv() protocol.NumberOfVersion
	}
	
	Content_Service_Count_Request interface {
		ContentID() [16]byte
	}
	
	Content_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	
	Content_Service_Get_Request interface {
		ContentID() [16]byte
		VersionOffset() uint64
	}
	
	Content_Service_Get_Response interface {
		Content
	}
	
	
	Content_Service_GetIDs_Request interface {
		Offset() uint64
		Limit() uint64
	}
	
	Content_Service_GetIDs_Response interface {
		ContentIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
	
	
	Content_Service_FindByUserID_Request interface {
		UserID() [16]byte
		Offset()  uint64
		Limit() uint64
	}
	
	Content_Service_FindByUserID_Response interface {
		ContentIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
	
	
)


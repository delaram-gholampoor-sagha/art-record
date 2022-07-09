package protocol

import "../libgo/protocol"

type Picture interface {
	RelatedID() [16]byte // any domain e.g. user, quiddity, ...
	ObjectID() [16]byte  //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type Picture_StorageServices interface {
	Save(p Picture) (err protocol.Error)

	// TODO::: how to change the order of pictures for a RelatedID

	Count(relatedID [16]byte) (numbers uint64, err protocol.Error)
	Get(relatedID [16]byte, versionOffset uint64) (p Picture, err protocol.Error)
	Last(relatedID [16]byte) (p Picture, numbers uint64, err protocol.Error)
}

type (
	Picture_Service_Register_Request interface {
		RelatedID() [16]byte 
		ObjectID() [16]byte  
	}
	Picture_Service_Register_Response interface {
		Numbers() uint64
	}
	Picture_Service_Count_Request interface {
		RelatedID() [16]byte    
	} 
	
	Picture_Service_Count_Response interface {
		Numbers() uint64
	
		
	}
	
	Picture_Service_Get_Request interface {
		RelatedID() [16]byte    
		VersionOffset() uint64
	}
	Picture_Service_Get_Response interface {
		Picture
		
	}
	
	Picture_Service_Last_Request interface {
		RelatedID() [16]byte
	
	}
	Picture_Service_Last_Response interface {
		Picture
		Numbers() uint64
	}
	
	
)
package protocol

type ContentComplement interface {
	ContentID() [16]byte    // content domain
	Priority() int32        // complement priority
	ComplementID() [16]byte // content domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}


type ContentComplement_StorageServices interface {
	Save(cp ContentComplement) protocol.Error

	Count(contentID [16]byte) (numbers uint64, err protocol.Error)
	Get(contentID [16]byte, versionOffset uint64) (cp ContentComplement, err protocol.Error)
	
}

type (
	ContentComplement_Service_Register_Request interface {
		ContentID() [16]byte    
   	Priority() int32        
	  ComplementID() [16]byte 
	}

	ContentComplement_Service_Register_Response interface {
		 Numbers() uint64
	}
	
	
	ContentComplement_Service_Count_Request interface {
		ContentID() [16]byte
	}
	
	ContentComplement_Service_Count_Response interface {
		Numbers() uint64
	}
	
	ContentComplement_Service_Get_Request interface {
		ContentID() [16]byte
		VersionOffset() uint64
	}
	
	ContentComplement_Service_Get_Response interface {
		ContentComplement
	}
	
)
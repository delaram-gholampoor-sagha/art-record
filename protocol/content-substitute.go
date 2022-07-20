/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

type ContentSubstitute interface {
	ContentID() [16]byte    // content domain
	Priority() int32        // substitute priority
	SubstituteID() [16]byte // content domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}


type ContentSubstitute_StorageServices interface {
	Save(c ContentSubstitute) protocol.Error

	Count(contentID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(contentID [16]byte, versionOffset uint64) (c ContentSubstitute,nv protocol.NumberOfVersion ,err protocol.Error)
	
}

type (
	ContentSubstitute_Service_Register_Request interface {
		ContentID() [16]byte    
	  Priority() int32        
	  SubstituteID() [16]byte 
	}

	ContentSubstitute_Service_Register_Response = protocol.NumberOfVersion

)

type (
	ContentSubstitute_Service_Count_Request interface {
		ContentID() [16]byte
	}

	ContentSubstitute_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	ContentSubstitute_Service_Get_Request interface {
		ContentID() [16]byte
		versionOffset() uint64
	}
	
	ContentSubstitute_Service_Get_Response1 interface {
		ContentSubstitute
	}

	ContentSubstitute_Service_Get_Response2 = protocol.NumberOfVersion
)

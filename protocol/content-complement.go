/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

type ContentComplement interface {
	ContentID() [16]byte    // content domain
	Priority() int32        // complement priority
	ComplementID() [16]byte // content domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}


type ContentComplement_StorageServices interface {
	Save(cp ContentComplement) protocol.Error

	Count(contentID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(contentID [16]byte, versionOffset uint64) (cp ContentComplement, nv protocol.NumberOfVersion,err protocol.Error)
	
}

type (
	ContentComplement_Service_Register_Request interface {
		ContentID() [16]byte    
   	Priority() int32        
	  ComplementID() [16]byte 
	}

	ContentComplement_Service_Register_Response =  protocol.NumberOfVersion

)

type (
	ContentComplement_Service_Count_Request interface {
		ContentID() [16]byte
	}
	
	ContentComplement_Service_Count_Response = protocol.NumberOfVersion
	
)


type (	
	ContentComplement_Service_Get_Request interface {
		ContentID() [16]byte
		versionOffset() uint64
	}
	
	ContentComplement_Service_Get_Response1 =	ContentComplement
	ContentComplement_Service_Get_Response2 = protocol.NumberOfVersion
	
)
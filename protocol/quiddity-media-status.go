/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"



type QuiddityMediaStatus interface {
	QuiddityID() [16]byte         // quiddity domain. any domain e.g. user, ...
	ObjectID() [16]byte           // object domain. media file
	Status() QuiddityMedia_Status //
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type QuiddityMediaStatus_StorageServices interface {
	Save(q QuiddityMediaStatus) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(quiddityID, objectID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(quiddityID, objectID [16]byte, vo protocol.versionOffset) (q QuiddityMediaStatus, nv protocol.NumberOfVersion, err protocol.Error)
}

type QuiddityMedia_Status Quiddity_Status

const (
// QuiddityMedia_Status_ QuiddityMedia_Status = (Quiddity_Status_FreeFlag << iota)
)


type (
	QuiddityMediaStatus_Service_Register_Request interface {
		QuiddityID() [16]byte 
	  ObjectID() [16]byte    
	  Status() QuiddityMedia_Status  
	}
	
	QuiddityMediaStatus_Service_Register_Response = protocol.NumberOfVersion

)

type (
	QuiddityMediaStatus_Service_Count_Request interface {
		QuiddityID() [16]byte
		ObjectID() [16]byte 
	}

	QuiddityMediaStatus_Service_Count_Response = protocol.NumberOfVersion
	
)

type (
	QuiddityMediaStatus_Service_Get_Request interface {
		QuiddityID() [16]byte
		ObjectID() [16]byte 
		versionOffset() uint64
	}

	QuiddityMediaStatus_Service_Get_Response1 = QuiddityMediaStatus
	QuiddityMediaStatus_Service_Get_Response2 = protocol.NumberOfVersion

)


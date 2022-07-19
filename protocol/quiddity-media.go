/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"


type QuiddityMedia interface {
	QuiddityID() [16]byte // quiddity domain. any domain e.g. user, ...
	Priority() uint64     // Media priority
	ObjectID() [16]byte   // object domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type QuiddityMedia_StorageServices interface {
	Save(p QuiddityMedia) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(quiddityID [16]byte, priority uint64) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(quiddityID [16]byte, priority uint64, vo protocol.VersionOffset) (p QuiddityMedia, nv protocol.NumberOfVersion, err protocol.Error)

	ListMedias(quiddityID [16]byte, offset, limit uint64) (objectIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListRelates(objectID [16]byte, offset, limit uint64) (quiddityIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListPriorities(objectID, quiddityID [16]byte, offset, limit uint64) (priorities []uint64, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	QuiddityMedia_Service_Register_Request interface {
		QuiddityID() [16]byte 
	  Priority() uint64     
	  ObjectID() [16]byte   
	}
	
	QuiddityMedia_Service_Register_Response = protocol.NumberOfVersion

)

type (
	QuiddityMedia_Service_Count_Request interface {
		QuiddityID() [16]byte
		Priority() uint64
	}

	QuiddityMedia_Service_Count_Response = protocol.NumberOfVersion
	
)

type (
	QuiddityMedia_Service_Get_Request interface {
		QuiddityID() [16]byte
		VersionOffset() uint64
		Priority() uint64
	}

	QuiddityMedia_Service_Get_Response1 = QuiddityMedia
	QuiddityMedia_Service_Get_Response2 = protocol.NumberOfVersion

)

type (
	QuiddityMedia_Service_ListRelates_Request interface {
		ObjectID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	QuiddityMedia_Service_ListRelates_Response1 interface {
    QuiddityIDs() [][16]byte
	}

	QuiddityMedia_Service_ListRelates_Response2 = protocol.NumberOfVersion
		
)

type (
	QuiddityMedia_Service_ListMedias_Request interface {
  	QuiddityID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	QuiddityMedia_Service_ListMedias_Response1 interface {
    ObjectIDs() [][16]byte
	}
	
	QuiddityMedia_Service_ListMedias_Response2 =  protocol.NumberOfVersion
)


type (
	QuiddityMedia_Service_ListPriorities_Request interface {
	  ObjectID() [16]byte
		QuiddityID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	QuiddityMedia_Service_ListPriorities_Response1 interface {
    Priorities() []uint64
	}	

	QuiddityMedia_Service_ListPriorities_Response2 = protocol.NumberOfVersion

)
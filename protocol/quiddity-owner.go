/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// QuiddityOwner indicate the domain record data fields.
type QuiddityOwner interface {
	QuiddityID() [16]byte // Unique content ID
	UserID() [16]byte     // user domain. Owner of the quiddity
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type QuiddityOwner_StorageServices interface {
	Save(qo QuiddityOwner) (nv protocol.NumberOfVersion , err protocol.Error)

	Count(quiddityID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(quiddityID [16]byte, versionOffset uint64) (qo QuiddityOwner,nv protocol.NumberOfVersion , err protocol.Error)


	FindByOwnerID(userID [16]byte, offset, limit uint64) (quiddityIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	QuiddityOwner_Service_Register_Request interface {
		QuiddityID() [16]byte
		UserID() [16]byte
	}
	
	QuiddityOwner_Service_Register_Response = protocol.NumberOfVersion

)

type (
	QuiddityOwner_Service_Count_Request interface {
		QuiddityID() [16]byte
	}

	QuiddityOwner_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	QuiddityOwner_Service_Get_Request interface {
		QuiddityID() [16]byte
		versionOffset() uint64
	}
	QuiddityOwner_Service_Get_Response1 = QuiddityOwner
	QuiddityOwner_Service_Get_Response2 = protocol.NumberOfVersion
	
)



type (	
	QuiddityOwner_Service_FindByUserID_Request interface {
		UserID() [32]byte
		Offset() uint64
		Limit() uint64
	}
	QuiddityOwner_Service_FindByUserID_Response1 interface {
		QuiddityIDs() [][32]byte
	}

	QuiddityOwner_Service_FindByUserID_Response2 = protocol.NumberOfVersion
)
/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// QuiddityModel indicate the domain record data fields.
// Model is the 3D model of the quiddity
type QuiddityModel interface {
	QuiddityID() [16]byte // quiddity domain
	ObjectID() [32]byte   // object domain. ThingModel
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

// ProductAttribute_Type_MetricUnit // Number, Gram, Liter, Second

type QuiddityModel_StorageServices interface {
	Save(qm QuiddityModel) (nv protocol.NumberOfVersion ,err protocol.Error)

	Count(quiddityID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(quiddityID [16]byte, versionOffset uint64) (qm QuiddityModel, err protocol.Error)


	FindByObjectID(objectID [16]byte, offset, limit uint64) (quiddityIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	QuiddityModel_Service_Register_Request interface {
		QuiddityID() [16]byte
		ObjectID() [16]byte
	}
	
	QuiddityModel_Service_Register_Response = protocol.NumberOfVersion

)

type (
	QuiddityModel_Service_Count_Request interface {
		QuiddityID() [16]byte
	}

	QuiddityModel_Service_Count_Response = protocol.NumberOfVersion
	
)

type (
	QuiddityModel_Service_Get_Request interface {
		QuiddityID() [16]byte
		versionOffset() uint64
	}

	QuiddityModel_Service_Get_Response1 = QuiddityModel
	QuiddityModel_Service_Get_Response2 = protocol.NumberOfVersion
	
)

type (	
	QuiddityModel_Service_FindByUserID_Request interface {
		UserID() [32]byte
		Offset() uint64
		Limit() uint64
	}
	QuiddityModel_Service_FindByUserID_Response1 interface {
		QuiddityIDs() [][32]byte
	}

	QuiddityModel_Service_FindByUserID_Response2 = protocol.NumberOfVersion
)
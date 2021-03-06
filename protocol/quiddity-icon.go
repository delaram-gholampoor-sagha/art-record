/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// QuiddityIcon indicate the domain record data fields.
// Icon or emoji use in many domains e.g. category, departments, ...
type QuiddityIcon interface {
	QuiddityID() [16]byte // quiddity domain
	Icon() [16]byte       // object domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type QuiddityIcon_StorageServices interface {
	Save(qi QuiddityIcon) ( nv protocol.NumberOfVersion ,err protocol.Error)

	Count(quiddityID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(quiddityID [16]byte, versionOffset uint64) (qi QuiddityIcon, err protocol.Error)

}

type (
	QuiddityIcon_Service_Register_Request interface {
		QuiddityID() [16]byte
		Icon() [16]byte
	}
	
	QuiddityIcon_Service_Register_Response = protocol.NumberOfVersion

)

type (
	QuiddityIcon_Service_Count_Request interface {
		QuiddityID() [16]byte
	}

	QuiddityIcon_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	QuiddityIcon_Service_Get_Request interface {
		QuiddityID() [16]byte
		versionOffset() uint64
	}

	QuiddityIcon_Service_Get_Response1 = QuiddityIcon
	QuiddityIcon_Service_Get_Response2 = protocol.NumberOfVersion
	
)

package protocol

import "../libgo/protocol"

type PictureStatus interface {
	ObjectID() [16]byte     // object domain
	Status() Picture_Status //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type PictureStatus_StorageServices interface {
	Save(ps PictureStatus) (err protocol.Error)

	Count(objectID [16]byte) (numbers uint64, err protocol.Error)
	Get(objectID [16]byte, versionOffset uint64) (ps PictureStatus, err protocol.Error)


	// FilterByStatus(status Picture_Status, offset, limit uint64) (objectIDs [][16]byte, numbers uint64, err protocol.Error)
	// protocol.EventTarget
}

type (
	PictureStatus_Service_Register_Request interface {
		ObjectID() [16]byte     
		Status() Picture_Status 
	
	}
	PictureStatus_Service_Register_Response interface {
			Nv() protocol.NumberOfVersion
	}
	
	PictureStatus_Service_Count_Request interface {
		ObjectID() [16]byte
		
	
	}
	PictureStatus_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	PictureStatus_Service_Get_Request interface {
		ObjectID() [16]byte    
		VersionOffset() uint64
	}
	PictureStatus_Service_Get_Response interface {
		PictureStatus
		Nv() protocol.NumberOfVersion
	}
	

)




type Picture_Status Quiddity_Status

const (
// Picture_Status_ Picture_Status = (Quiddity_Status_FreeFlag << iota)
)

package protocol

import "../libgo/protocol"

// RoleStaff indicate the domain record data fields.
type RoleStaff interface {
	RoleID() [16]byte    // role domain
	Minimum() int        // Minimum number of employees for this job position
	Expected() int       // Expected number of employees for this job position
	Maximum() int        // Maximum number of employees for this job position
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type RoleStaff_StorageServices interface {
	Save(rs RoleStaff) protocol.Error

	Count(roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (rs RoleStaff, err protocol.Error)
	
}


 
type (

	RoleStaff_Service_Register_Request interface{
		RoleID() [16]byte
		Minimum() int        
		Expected() int       
		Maximum() int        
  }


	RoleStaff_Service_Register_Response interface{
		Nv() protocol.NumberOfVersion      
	}
	
	
	RoleStaff_Service_Count_Request interface{
		RoleID() [16]byte
	}
	
	RoleStaff_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	RoleStaff_Service_Get_Request interface{
		RoleID() [16]byte
		VersionOffset() uint64
	}
	
	RoleStaff_Service_Get_Response interface{
		RoleStaff
		Nv() protocol.NumberOfVersion
	}
	
	RoleStaff_Service_Last_Request interface{
		RoleID() [16]byte
	}
	
	RoleStaff_Service_Last_Response interface{
		RoleStaff
		Nv() protocol.NumberOfVersion
	}
)
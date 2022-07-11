package protocol

import "../libgo/protocol"

// RoleSalary is fixed negotiated salary that don't follow org||department||role salary rules.
type RoleSalary interface {
	RoleID() [16]byte                     // role domain
	WeeklySalary() protocol.AmountOfMoney //
	Time() protocol.Time                  // save time
	RequestID() [16]byte                  // user-request domain
}



type RoleSalary_StorageServices interface {
	Save(rs RoleSalary) protocol.Error

	Count(roleID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (rs RoleSalary, nv protocol.NumberOfVersion  , err protocol.Error)
	
}

type (

	RoleSalary_Service_Register_Request interface{
		RoleID() [16]byte                    
	  WeeklySalary() protocol.AmountOfMoney 
  }

	RoleSalary_Service_Register_Response interface{
		Nv() protocol.NumberOfVersion      
	}
	
	
	RoleSalary_Service_Count_Request interface{
		RoleID() [16]byte
	}
	
	RoleSalary_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	RoleSalary_Service_Get_Request interface{
		RoleID() [16]byte
		VersionOffset() uint64
	}
	
	RoleSalary_Service_Get_Response interface{
		RoleSalary
		Nv() protocol.NumberOfVersion
	}
	
	RoleSalary_Service_Last_Request interface{
		RoleID() [16]byte
	}
	
	RoleSalary_Service_Last_Response interface{
		RoleSalary
		Nv() protocol.NumberOfVersion
	}
)
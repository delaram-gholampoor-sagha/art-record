package protocol

import (
	"../libgo/protocol"
)

type CategorySubstitute interface {
	CategoryID() [16]byte   // category domain
	Priority() uint64       // Substitute priority
	SubstituteID() [16]byte // category domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type CategorySubstitute_StorageServices interface {
	Save(cs CategorySubstitute) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(categoryID [16]byte, priority uint64) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(categoryID [16]byte, priority uint64, vo protocol.VersionOffset) (cs CategorySubstitute, nv protocol.NumberOfVersion, err protocol.Error)

	ListSubstitute(categoryID [16]byte, offset, limit uint64) (substituteIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListCategories(substituteID [16]byte, offset, limit uint64) (categoryIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListPriorities(substituteID, categoryID [16]byte, offset, limit uint64) (priorities []uint64, nv protocol.NumberOfVersion, err protocol.Error)
}

type (

		CategorySubstitute_Service_Register_Request interface {
		CategoryID() [16]byte  
		Priority() uint64       
		SubstituteID() [16]byte 
	}
	CategorySubstitute_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (

	CategorySubstitute_Service_Count_Request interface {
		CategoryID() [16]byte
		Priority() uint64
			
	}
	CategorySubstitute_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
			
	}
	
)




type (

		CategorySubstitute_Service_Get_Request interface {
		CategoryID() [16]byte
		priority() uint64    
		Vo() protocol.VersionOffset
	}
	CategorySubstitute_Service_Get_Response interface {
		CategorySubstitute
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)


type (
	CategorySubstitute_Service_ListSubstitutes_Request interface {
		SubstituteID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	CategorySubstitute_Service_ListSubstitutes_Response interface {
		SubstituteIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)


type (

	CategorySubstitute_Service_ListCategories_Request interface {
		SubstituteID() [16]byte
		Offset() uint64
		Limit() uint64
		
	}
	CategorySubstitute_Service_ListCategories_Response interface {
		CategoryIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
		
	}
	
)


type (
	
	CategorySubstitute_Service_ListPriorities_Request interface {
		SubstituteID() [16]byte
		Offset() uint64
		Limit() uint64
		
	}
	CategorySubstitute_Service_ListPriorities_Response interface {
		Priorities() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
		
	}
)
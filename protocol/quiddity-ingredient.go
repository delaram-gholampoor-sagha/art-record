package protocol

import "../libgo/protocol"

// ProductIngredient indicate the domain record data fields.
type QuiddityIngredient interface {
	QuiddityID() [16]byte    // quiddity domain
	IngredientID() [16]byte  // quiddity domain
	Time() protocol.Time     // save time
	RequestID() [16]byte     // user-request domain
}

type ProductIngredient_StorageServices interface {
	Save(pi QuiddityIngredient) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(QuiddityID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(QuiddityID [16]byte, versionOffset uint64) (pi QuiddityIngredient, nv protocol.NumberOfVersion, err protocol.Error)

	FindByIngredient(ingredientID [16]byte, offset, limit uint64) (QuiddityIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}



type (
	ProductIngredient_Service_Register_Request interface {
		QuiddityID() [16]byte    
		IngredientID() [16]byte 
	}
	ProductIngredient_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
		ProductIngredient_Service_Count_Request interface {
		QuiddityID() [16]byte
	}
	ProductIngredient_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)


type (
		ProductIngredient_Service_Get_Request interface {
		QuiddityID() [16]byte
		VersionOffset() uint64
	}
	ProductIngredient_Service_Get_Response interface {
    QuiddityIngredient
 		NumberOfVersion() protocol.NumberOfVersion
	}
	
)


type (
	ProductIngredient_Service_FindByIngredient_Request interface {
		IngredientID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	ProductIngredient_Service_FindByIngredient_Response interface {
		QuiddityIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
)
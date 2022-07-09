package protocol

import "../libgo/protocol"

// ProductIngredient indicate the domain record data fields.
type ProductIngredient interface {
	ProductID() [16]byte    // product domain
	IngredientID() [16]byte // product domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ProductIngredient_StorageServices interface {
	Save(pi ProductIngredient) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pi ProductIngredient, numbers uint64, err protocol.Error)

	FindByIngredient(ingredientID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}

type (
	ProductIngredient_Service_Register_Request interface {
		ProductID() [16]byte    
		IngredientID() [16]byte 
	}
	ProductIngredient_Service_Register_Response interface {
		Numbers() uint64
	}
	
	ProductIngredient_Service_Count_Request interface {
		ProductID() [16]byte
	}
	ProductIngredient_Service_Count_Response interface {
		Numbers() uint64
	}
	
	ProductIngredient_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	ProductIngredient_Service_Get_Response interface {
		ProductIngredient
		Numbers() uint64
	}
	
	ProductIngredient_Service_FindByIngredient_Request interface {
		IngredientID() [16]byte
		Offset() uint64
		Limit() uint64}
	ProductIngredient_Service_FindByIngredient_Response interface {
		ProductIDs() [][16]byte
		Numbers() uint64
	}
)
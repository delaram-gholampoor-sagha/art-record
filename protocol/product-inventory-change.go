package protocol

// ProductInventoryChange indicate the domain record data fields.
type ProductInventoryChange interface {
	ProductInventoryChangeID() [16]byte //
	ProductID() [16]byte                // product domain
	From() [16]byte                     // building-location domain
	To() [16]byte                       // building-location domain. Can be zero for voiding products.
	Amount() int64                      // this transaction
	Priority() uint64                   //
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}

type ProductInventoryChange_StorageServices interface {
	Save(pi ProductInventoryChange) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pi ProductInventoryChange, numbers uint64, err protocol.Error)

	FindByProduct(productID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}

type (
	
	ProductInventoryChange_Service_Register_Request interface {
		ProductInventoryChangeID() [16]byte 
		ProductID() [16]byte                
		From() [16]byte                   
		To() [16]byte                       
		Amount() int64                      
		Priority() uint64                   
	}
	ProductInventoryChange_Service_Register_Response interface {
		Numbers() uint64
	}

	ProductInventoryChange_Service_Count_Request interface {
		ProductID() [16]byte
	}
	ProductInventoryChange_Service_Count_Response interface {
		Numbers()uint64
	}

	ProductInventoryChange_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	ProductInventoryChange_Service_Get_Response interface {
		ProductInventoryChange
		Numbers() uint64
	}

	ProductInventoryChange_Service_FindByProduct_Request interface {
		ProductID() [16]byte
		Offset() uint64
		Limit() uint64

	}
	ProductInventoryChange_Service_FindByProduct_Response interface {
		ProductIDs() [][16]byte
		Numbers() uint64
	}
)
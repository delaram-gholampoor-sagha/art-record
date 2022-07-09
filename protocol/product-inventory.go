package protocol

// ProductInventory indicate the domain record data fields.
// ProductInventory or good stock
type ProductInventory interface {
	ProductID() [16]byte                // product domain
	BuildingLocationID() [16]byte       // building-location domain
	ReferenceID() [16]byte              // Depend on ReferenceType()
	ReferenceType() ProductInventory_RT //
	Amount() int64                      // this transaction
	Stock() int64                       // overall in each location
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}

type ProductInventory_StorageService interface {
	Save(p ProductInventory) (numbers uint64, err protocol.Error)

	Lock(productID, buildingLocationID [16]byte) (p ProductInventory, err protocol.Error)
	Unlock(p ProductInventory) (numbers uint64, err protocol.Error)

	Count(productID, buildingLocationID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID, buildingLocationID [16]byte, versionOffset uint64) (p ProductInventory, numbers uint64, err protocol.Error)

	ListProductIDs(buildingLocationID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
	ListBuildingLocationIDs(productID [16]byte, offset, limit uint64) (buildingLocationIDs [][16]byte, numbers uint64, err protocol.Error)
}



type (

	ProductInventory_Service_Register_Request interface {
		ProductID() [16]byte               
		BuildingLocationID() [16]byte       
		ReferenceID() [16]byte              
		ReferenceType() ProductInventory_RT 
		Amount() int64                      
		Stock() int64                       


	}

	ProductInventory_Service_Register_Response interface {
		Numbers() uint64
	}
	ProductInventory_Service_Lock_Request interface {
		ProductID()[16]byte
		BuildingLocationID() [16]byte
	}
	ProductInventory_Service_Lock_Response interface {
		ProductInventory
	}
	
	ProductInventory_Service_Unlock_Request interface {
		ProductInventory
	}
	ProductInventory_Service_Unlock_Response interface {
		Numbers() uint64
	
	}
	
	ProductInventory_Service_Count_Request interface {
		ProductID() [16]byte
		BuildingLocationID() [16]byte
	
	}
	ProductInventory_Service_Count_Response interface {
		Numbers() uint64
	}
	
	ProductInventory_Service_Get_Request interface {
		ProductID() [16]byte
		BuildingLocationID() [16]byte    
		VersionOffset() uint64
	}
	ProductInventory_Service_Get_Response interface {
		ProductInventory
		Numbers() uint64
	}
)



type ProductInventory_RT uint8

const (
	ProductInventory_RT_Unset        ProductInventory_RT = iota
	ProductInventory_RT_Production                       // ProductionID
	ProductInventory_RT_Invoice                          // InvoiceID
	ProductInventory_RT_Order                            // ProductOrderID
	ProductInventory_RT_Transfer                         // ProductInventoryChangeID
	ProductInventory_RT_ChangeOwner                      // UserID
	ProductInventory_RT_SplitProduct                     // Split to small size product || Split from upper size product
	ProductInventory_RT_Void                             //
	// ProductInventory_RT_
)

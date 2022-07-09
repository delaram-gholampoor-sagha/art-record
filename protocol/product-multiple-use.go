package protocol

// ProductMultipleUse indicate the domain record data fields.
type ProductMultipleUse interface {
	ProductID() [16]byte        // product domain
	Maximum() uint64            // Maximum time this service can be served to buyer
	Timeout() protocol.Duration // from first use
	Time() protocol.Time        // save time
	RequestID() [16]byte        // user-request domain
}

type ProductMultipleUse_StorageServices interface {
	Save(pm ProductMultipleUse) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pm ProductMultipleUse, numbers uint64, err protocol.Error)
}


type (
	ProductMultipleUse_Service_Register_Request interface {
		ProductID() [16]byte        
		Maximum() uint64           
		Timeout() protocol.Duration 
	}
	ProductMultipleUse_Service_Register_Response interface {
		Numbers() uint64
	}
	//	Count(productID [16]byte, currency uint64) (numbers uint64, err protocol.Error)
	
	ProductMultipleUse_Service_Count_Request interface {
		ProductID() [16]byte
	
	}
	ProductMultipleUse_Service_Count_Response interface {
		Numbers() uint64
	}
	
	ProductMultipleUse_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	ProductMultipleUse_Service_Get_Response interface {
		ProductMultipleUse
		Numbers() uint64
	}
)
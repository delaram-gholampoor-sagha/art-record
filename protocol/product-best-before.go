package protocol

// ProductBestBefore indicate the domain record data fields.
// https://en.wikipedia.org/wiki/Expiration_date
type ProductBestBefore interface {
	ProductID() [16]byte         // product domain
	Duration() protocol.Duration // from production time
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type ProductBestBefore_StorageServices interface {
	Save(pb ProductBestBefore) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pb ProductBestBefore, numbers uint64, err protocol.Error)
}

type (
	ProductBestBefore_Service_Register_Request interface {
		ProductID() [16]byte 
		Duration() protocol.Duration     
	
	}
	ProductBestBefore_Service_Register_Response interface {
		Numbers() uint64
	}
	
	ProductBestBefore_Service_Count_Request interface {
		ProductID() [16]byte
	
	}
	ProductBestBefore_Service_Count_Response interface {
		Numbers() uint64
	}
	
	ProductBestBefore_Service_Get_Request interface {
		ProductID() [16]byte    
		VersionOffset() uint64
	}
	ProductBestBefore_Service_Get_Response interface {
		ProductBestBefore
		Numbers() uint64
	}
	
)

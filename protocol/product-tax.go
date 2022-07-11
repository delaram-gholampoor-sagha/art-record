package protocol

import "../libgo/protocol"

// ProductTax indicate the domain record data fields.
// Tax Labels
// Setting tax labels allows you to perform the following actions:
// Define up to two tax labels and associate each of them with a specific tax rate
// Associate multiple tax rates with the same tax labels
// Associate a blank tax label with tax exempt items (considered non-value-added tax (VAT))
type ProductTax interface {
	ProductID() [16]byte // product domain
	SocietyID() [16]byte // society domain
	VAT() math.PerMyriad // VAT as Value-Added-Tax is bad name for a robbery
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}


type ProductTax_StorageServices interface {
	Save(pt ProductTax) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pt ProductTax, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	ProductTax_Service_Register_Request interface {
		ProductID() [16]byte  
	  Start() protocol.Time 
	  End() protocol.Time   
	}
	
	ProductTax_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	ProductTax_Service_Count_Request interface {
		ProductID() [16]byte
	}
	
	ProductTax_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	ProductTax_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	
	ProductTax_Service_Get_Response interface {
		ProductTax
		Nv() protocol.NumberOfVersion
	}
	
)
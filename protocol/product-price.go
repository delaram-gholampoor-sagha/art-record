/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductPrice indicate the domain record data fields.
// ProductPrice can be register just by producer organization whom own the related quiddity
type ProductPrice interface {
	ProductID() [16]byte           // product domain
	Currency() [16]byte            // financial-currency
	Price() protocol.AmountOfMoney // Basic Price not payable price in the currency
	Time() protocol.Time           // save time
	RequestID() [16]byte           // user-request domain
}

type ProductPrice_StorageServices interface {
	Save(pp ProductPrice) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte, currency uint64) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, currency uint64, versionOffset uint64) (pp ProductPrice, nv protocol.NumberOfVersion, err protocol.Error)

	ListProductCurrencies(productID [16]byte, offset, limit uint64) (currencies []uint64, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	ProductPrice_Service_Register_Request interface {
		ProductID() [16]byte
		Currency() [16]byte
	}
	ProductPrice_Service_Register_Response = protocol.NumberOfVersion

)

type (
	ProductPrice_Service_Count_Request interface {
		ProductID() [16]byte
		Currency() uint64
	}
	ProductPrice_Service_Count_Response = protocol.NumberOfVersion

)

type (
	ProductPrice_Service_Get_Request interface {
		ProductID() [16]byte
		Currency() uint64
		versionOffset() uint64
	}

	ProductPrice_Service_Get_Response1 = 	ProductPrice
	ProductPrice_Service_Get_Response2 =  protocol.NumberOfVersion
	
)

type (
	ProductPrice_Service_ListProductCurrencies_Request interface {
		ProductID() [16]byte
		Offset()
		Limit() uint64
	}
	ProductPrice_Service_ListProductCurrencies_Response1 interface {
		Currencies() []uint64
	}

	ProductPrice_Service_ListProductCurrencies_Response2 = protocol.NumberOfVersion
	
)
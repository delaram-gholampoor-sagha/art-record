/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductPricePayable indicate the domain record data fields.
// ProductPricePayable store historically price with discount to show to users in price change chart.
type ProductPricePayable interface {
	ProductID() [16]byte // product domain
	// OrgID() [16]byte               //
	Currency() uint64              //
	Price() protocol.AmountOfMoney // Basic Price not payable price in the society currency
	Time() protocol.Time           // save time
	RequestID() [16]byte           // user-request domain
}

type ProductPricePayable_StorageServices interface {
	Save(q ProductPricePayable) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte, currency uint64) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, currency uint64, versionOffset uint64) (q ProductPricePayable, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	ProductPricePayable_Service_Register_Request interface {
		ProductID() [16]byte
		Currency() uint64 
		Price()  protocol.AmountOfMoney
	}

	ProductPricePayable_Service_Register_Response = protocol.NumberOfVersion

)

type (
	ProductPricePayable_Service_Count_Request interface {
		ProductID() [16]byte
		Currency() uint64
	}

	ProductPricePayable_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	ProductPricePayable_Service_Get_Request interface {
		ProductID() [16]byte
		Currency() uint64
		versionOffset() uint64
	}
	
	ProductPricePayable_Service_Get_Response1 = ProductPricePayable
	ProductPricePayable_Service_Get_Response2 = protocol.NumberOfVersion
	
)
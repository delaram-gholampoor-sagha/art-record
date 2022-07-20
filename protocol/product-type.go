/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

type ProductType interface {
	ProductID() [16]byte // product-status domain
	Kind() Product_Kind  //
	Type() Product_Type  //
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type Product_Kind uint8

const (
	Product_Kind_Unset Product_Kind = 0 // usually for good kind

	Product_Kind_Inventory Product_Kind = (1 << iota) //
	Product_Kind_BestBefore

	Product_Kind_TimeValidity
	Product_Kind_Timing
	Product_Kind_Place
	Product_Kind_Provider  // a user prodive the product
	Product_Kind_Transport // indicate InvoiceID use to retrieve coordinate domain
	Product_Kind_Station   // service that provides in stations e.g. bus, train, airport,
	Product_Kind_Share     // ticket for panic room, ...	
)

// TODO::: remove in favor of product tags??
type Product_Type int8

const (
	Product_Type_Good             Product_Type = (-1 - iota) //
	Product_Type_Good_Content                                // Digital data
	Product_Type_Good_Consumption                            // any Good usually use daily to satisfy human wants.
	Product_Type_Good_Commodity                              // a product of agriculture or mining
	Product_Type_Good_Medical                                // can't add by user to invoices
	Product_Type_Good_Capital                                // Book, Mobile, Car, ...

	Product_Type_Unset Product_Type = iota

	// Service is the structure of TimeService like:
	// Consultation in chat or real life
	// It can yse for any service or goods e.g.
	// ActService like guarding, maintenance, cleaning, ...
	// Transport
	Product_Type_Service                  // FixingService, ...
	Product_Type_Service_Time             // Therapy, Consulting, Coaching, Mentoring, ...
	Product_Type_Service_TimePlace        // Rent, ...
	Product_Type_Service_TimeDefiendPlace // Cinema, ...

	Product_Type_Service_Shipping_Road // the transportation of goods
	Product_Type_Service_Shipping_Air  // the transportation of goods
	Product_Type_Service_Shipping_Sea  // the transportation of goods
	Product_Type_Service_Shipping_Rail // the transportation of goods

	Product_Type_Service_Transport_Road // the transportation of passenger
	Product_Type_Service_Transport_Air  // the transportation of passenger
	Product_Type_Service_Transport_Sea  // the transportation of passenger
	Product_Type_Service_Transport_Rail // the transportation of passenger
)


type ProductType_StorageServices interface {
	Save(pt ProductType) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pt ProductType, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	ProductType_Service_Register_Request interface {
		ProductID() [16]byte 
	  Kind() Product_Kind  
	  Type() Product_Type  
	}

	ProductType_Service_Register_Response = protocol.NumberOfVersion
)

type (
	ProductType_Service_Count_Request interface {
		ProductID() [16]byte
	}

	ProductType_Service_Count_Response = protocol.NumberOfVersion

)


type (
	ProductType_Service_Get_Request interface {
		ProductID() [16]byte
		versionOffset() uint64
	}

	ProductType_Service_Get_Response1 = ProductType
	ProductType_Service_Get_Response2 = protocol.NumberOfVersion
)
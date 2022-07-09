package protocol

// Product indicate the domain record data fields.
type Product interface {
	ProductID() [16]byte // quiddity domain
	Type() Product_Type  //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type Product_StorageServices interface {
	Save(p Product) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (p Product, numbers uint64, err protocol.Error)

	FilterByType(typ Product_Type, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}

// Cinema, Rent, Therapy, Consulting, Coaching, Mentoring, FixingService, the transportation of goods and passenger,  ...
type Product_Type uint64

// Common flags
const (
	Product_Type_Unset Product_Type = 0           // usually for good kind
	Product_Type_Good  Product_Type = (1 << iota) // if flag not set means product is service
	Product_Type_ReserveValidity
	Product_Type_ReserveQuantity  // max product can reserve by one person before pay them
	Product_Type_PreSale          // ??
	Product_Type_NeedApprove      //
	Product_Type_NeedPrescription // Medical good. can't add by users to invoice

	Product_Type_InvoiceSell
	Product_Type_OrderSell
	Product_Type_Product  // Product limit sell
	Product_Type_Group // group limit sell

	_
	_
	_
	_
	_
	_
	Product_Type_StartSpecific // it is start of good or service flags
)

// Good flags
const (
	Product_Type_Inventory  Product_Type = (Product_Type_StartSpecific << iota) // usually capital good must track owner. Book, Mobile, Car, ...
	Product_Type_Custom                                                         // made on buyer requirements, product-ingredient
	Product_Type_BestBefore                                                     // Consumption good. any Good usually use daily to satisfy human wants.
	Product_Type_Commodity                                                      // a product of agriculture or mining
)

// Services flags
// Service is the structure of TimeService like:
// Consultation in chat or real life
// It can yse for any service or goods e.g.
// ActService like guarding, maintenance, cleaning, ...
// Transport
const (
	Product_Type_TimeValidity   Product_Type = (Product_Type_StartSpecific << iota)
	Product_Type_Transport                   // indicate InvoiceID use to retrieve coordinate domain
	Product_Type_Shipping                    //
	Product_Type_Content                     // for any content types, Digital data
	Product_Type_Timing                      // the time service will provide
	Product_Type_TimingInterval              //
	Product_Type_MultipleUse                 // gym, ...
	Product_Type_Number                      // cinema, concert, transport, ...
	Product_Type_Area                        // transport, ...
	Product_Type_Place                       // limit to get service on specific place like doctor-room, gym, station(bus, train, airport, ...)
	Product_Type_Seat                        // cinema, concert, transport, ...
	Product_Type_Provider_Staff              // the provider of the product
	Product_Type_Provider_Role               // the provider of the product
	Product_Type_Share                       // ticket for panic room, ...
)


type (
	Product_Service_Register_Request interface{
		Type() Product_Type
	}
	Product_Service_Register_Response interface{
		ProductID() [16]byte
		Numbers() uint64
	}

	Product_Service_Count_Request interface{
		ProductID() [16]byte
	}

	Product_Service_Count_Response interface{
		Numbers() uint64
	}
	Product_Service_Get_Request interface{
		ProductID() [16]byte
		VersionOffset() uint64
	}

	Product_Service_Get_Response interface{
		Product
		Numbers() uint64
	}

	Product_Service_FilterByType_Request interface{
		Product_Type() uint64
		Offset() uint64
		Limit() uint64
	}

	Product_Service_FilterByType_Response interface{
		ProductIDs() [][16]byte
		Numbers() uint64
	}
)
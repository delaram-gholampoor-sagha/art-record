/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// InvoiceSides indicate the domain record data fields.
type InvoiceSides interface {
	InvoiceID() [16]byte  // invoice domain
	UserID() [16]byte   // user domain
	Type() InvoiceSides_Type // user domain
	Status() [16]byte    // user domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type InvoiceSides_StorageServices interface {
	Save(is InvoiceSides) (nv protocol.NumberOfVersion , err protocol.Error)

	Get(invoiceID [16]byte) (is InvoiceSides , nv protocol.NumberOfVersion,err protocol.Error)
	Count(invoiceID [16]byte) (nv protocol.NumberOfVersion,err protocol.Error)

	FindByUser(userID [16]byte , Type InvoiceSides_Type , offset uint64 , limit uint64) (invoiceIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

}

type InvoiceSides_Type uint8

const (
	InvoiceSides_Type_Seller InvoiceSides_Type = iota
	InvoiceSides_Type_Customer
	InvoiceSides_Type_Agent
	InvoiceSides_Type_Franchise
)

type InvoiceSides_Status = Quiddity_Status


type (
	InvoiceSides_Service_Register_Request interface {
		InvoiceID() [16]byte  
	  UserID() [16]byte   
	  Type() InvoiceSides_Type 
    Status() [16]byte   
	}
		InvoiceSides_Service_Register_Response = protocol.NumberOfVersion
)


type (
	InvoiceSides_Service_Get_Request interface {
		InvoiceID() [16]byte      
	}

	InvoiceSides_Service_Get_Response1 = InvoiceSides
	InvoiceSides_Service_Get_Response2 = protocol.NumberOfVersion

)

type (
	InvoiceSides_Service_FindByUser_Request interface { 
	  UserID() [16]byte 
	  Type() InvoiceSides_Type
		Offset() uint64 
		Limit() uint64
	}
	InvoiceSides_Service_FindByUser_Response1 interface {
		InvoiceIDs() [][16]byte 
	}

	InvoiceSides_Service_FindBySeller_Response2 = protocol.NumberOfVersion

)

type (
	InvoiceSides_Service_Count_Request interface { 
	  InvoiceID() [16]byte 
	}

	InvoiceSides_Service_Count_Response2 = protocol.NumberOfVersion
	
)




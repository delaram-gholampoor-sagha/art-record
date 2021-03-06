/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductProvider indicate the domain record data fields.
type ProductProviderStaff interface {
	ProductID() [16]byte // product domain
	StaffID() [16]byte   // staff domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductProviderStaff_StorageServices interface {
	Save(pp ProductProviderStaff) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pp ProductProviderStaff, nv protocol.NumberOfVersion, err protocol.Error)

	FindByStaff(staffID [16]byte, offset, limit uint64) (productIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

// related UserIDs by StaffID use to get Coordinate(domain)

type (
	ProductProviderStaff_Service_Register_Request interface {
		ProductID() [16]byte
		StaffID() [16]byte
	}

	ProductProviderStaff_Service_Register_Response = protocol.NumberOfVersion

)

type (
	ProductProviderStaff_Service_Count_Request interface {
		ProductID() [16]byte
	}

	ProductProviderStaff_Service_Count_Response = protocol.NumberOfVersion

)


type (
  ProductProviderStaff_Service_Get_Request interface {
		ProductID() [16]byte
		versionOffset() uint64	
	}

	ProductProviderStaff_Service_Get_Response1  = 	ProductProviderStaff
	ProductProviderStaff_Service_Get_Response2 = protocol.NumberOfVersion
)


type (

	ProductProviderStaff_Service_FindByStaff_Request interface {
		StaffID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	ProductProviderStaff_Service_FindByStaff_Response1 interface {
		ProductIDs() [][16]byte
	}
	
	ProductProviderStaff_Service_FindByStaff_Response2 = protocol.NumberOfVersion
)


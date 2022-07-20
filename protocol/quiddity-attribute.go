/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductAttribute store some key/value for a product to describe it e.g. mobile-cpu, ...
// TODO::: is it good idea to add detail manually by users?? or quiddity-model is better??
type QuiddityAttribute interface {
	QuiddityID() [16]byte         // product domain
	Type() ProductAttribute_Type //
	Value() string               //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}


type QuiddityAttribute_StorageServices interface {
	Save(qi QuiddityAttribute) (nv protocol.NumberOfVersion , err protocol.Error)

	Count(quiddityID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(quiddityID [16]byte, versionOffset uint64) (qi QuiddityAttribute, err protocol.Error)

}

type ProductAttribute_Type uint16

const (
	ProductAttribute_Type_Unset ProductAttribute_Type = iota

	ProductAttribute_Type_MetricUnit // Number, Gram, Liter, Second
)

type (
	QuiddityAttribute_Service_Register_Request interface {
		QuiddityID() [16]byte
		Type() ProductAttribute_Type 
		Value() string  
	}
	
	QuiddityAttribute_Service_Register_Response = protocol.NumberOfVersion

)

type (
	QuiddityAttribute_Service_Count_Request interface {
		QuiddityID() [16]byte
	}

	QuiddityAttribute_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	QuiddityAttribute_Service_Get_Request interface {
		QuiddityID() [16]byte
		versionOffset() uint64
	}

	QuiddityAttribute_Service_Get_Response1 = QuiddityAttribute
	QuiddityAttribute_Service_Get_Response2 = protocol.NumberOfVersion
	
)

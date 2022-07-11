package protocol

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


type ProductAttribute_StorageServices interface {
	Save(qi ProductAttribute) (err protocol.Error)

	Count(quiddityID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(quiddityID [16]byte, versionOffset uint64) (qi ProductAttribute, err protocol.Error)

}

type ProductAttribute_Type uint16

const (
	ProductAttribute_Type_Unset ProductAttribute_Type = iota

	ProductAttribute_Type_MetricUnit // Number, Gram, Liter, Second
)

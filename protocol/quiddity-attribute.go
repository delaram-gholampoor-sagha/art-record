package protocol


// ProductAttribute store some key/value for a product to describe it e.g. mobile-cpu, ...
// TODO::: is it good idea to add detail manually by users?? or quiddity-model is better??
type ProductAttribute interface {
	ProductID() [16]byte         // product domain
	Type() ProductAttribute_Type //
	Value() string               //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type ProductAttribute_Type uint16

const (
	ProductAttribute_Type_Unset ProductAttribute_Type = iota

	ProductAttribute_Type_MetricUnit // Number, Gram, Liter, Second
)

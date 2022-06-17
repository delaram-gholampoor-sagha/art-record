package protocol

// ProductIngredient indicate the domain record data fields.
type ProductIngredient interface {
	ProductID() [16]byte    // product domain
	IngredientID() [16]byte // product domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ProductIngredient_StorageServices interface {
	Save(pi ProductIngredient) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pi ProductIngredient, err protocol.Error)
	Last(productID [16]byte) (pi ProductIngredient, numbers uint64, err protocol.Error)

	FindByIngredient(ingredientID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}

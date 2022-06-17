package protocol

// ProductRole indicate the domain record data fields.
type ProductRole interface {
	ProductID() [16]byte // product domain
	RoleID() [16]byte    // role domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductRole_StorageServices interface {
	Save(pr ProductRole) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pr ProductRole, err protocol.Error)
	Last(productID [16]byte) (pr ProductRole, numbers uint64, err protocol.Error)

	FindByRole(roleID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}

package protocol

// Category or Topic
type Category interface {
	CategoryID() [16]byte // quiddity domain
	ParentID() [16]byte   // CategoryID, Exist if it isn't top category.
	Priority() uint32     //
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type Category_StorageServices interface {
	Save(c Category) (err protocol.Error)

	GetIDs(offset, limit uint64) (categoryIDs [][16]byte, numbers uint64, err protocol.Error)

	Delete(categoryIDs [16]byte) (err protocol.Error)
}

// TODO::: cache strategy?? 1 day or what?

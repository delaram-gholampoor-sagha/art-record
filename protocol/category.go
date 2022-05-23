package protocol

import "time"

// Category or Topic
type Category interface {
	QuiddityID() [16]byte // CategoryID
	ParentID() [16]byte   // CategoryID, Exist if it isn't top category.
	Time() time.Time      // Save time
	RequestID() [16]byte  // user-request domain
}

// TODO::: cache strategy?? 1 day or what?
type Category_StorageServices interface {
	Save(c Category) (err error)

	// count 
	Count(QuiddityID [16]byte) (numbers uint64, err error)
	
	// az inja 70 , 30 ta behem bede
	GetIDs(offset, limit uint64) (quiddityIDs [][16]byte, length uint64, err error)

	FindCategoryByParentID(id [16]byte) (ci [16]byte, err error)

	Delete(quiddityID [16]byte) (err error)
}


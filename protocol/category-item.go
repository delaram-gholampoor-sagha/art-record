package protocol

import "time"

type CategoryItem interface {
	ItemID() [16]byte     // any domain usually product, content, ... domain
	ItemDomainID() uint64 // Domain MediatypeID
	CategoryID() [16]byte // category domain
	Time() time.Time      // Save time
	RequestID() [16]byte  // user-request domain
}

type CategoryItem_StorageServices interface {
	Save(ci CategoryItem) (err error)

	Count(itemID [16]byte) (length uint64, err error)
	Get(itemID [16]byte, versionOffset uint64) (ci CategoryItem, err error)
	Last(itemID [16]byte) (ci CategoryItem, length uint64, err error)

	// TODO::: bad scenario: one admin made 90% discount for a product category, other admin assign a product to that category,
	// after buy a lot of categorized product, it will delete the product category.
	Delete(itemID [16]byte) (err error)

	FindByCategoryID(categoryID [16]byte, offset, limit uint64) (itemIDs [][16]byte, numbers uint64, err error)
}

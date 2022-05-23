package protocol

import "time"

type Picture interface {
	RelatedID() [16]byte     // any domain e.g. user, quiddity, ...
	RelatedDomainID() uint64 //
	ObjectID() [16]byte      //
	Time() time.Time         // Save time
	RequestID() [16]byte     // user-request domain
}

type Picture_StorageServices interface {
	Save(p Picture) (err error)

	// TODO::: how to change the order of pictures for a RelatedID

	Count(relatedID [16]byte) (numbers uint64, err error)
	Get(relatedID [16]byte, versionOffset uint64) (p Picture, err error)
	Last(relatedID [16]byte) (p Picture, numbers uint64, err error)
}

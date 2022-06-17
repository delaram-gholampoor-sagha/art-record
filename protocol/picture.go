package protocol

type Picture interface {
	RelatedID() [16]byte // any domain e.g. user, quiddity, ...
	ObjectID() [16]byte  //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type Picture_StorageServices interface {
	Save(p Picture) (err protocol.Error)

	// TODO::: how to change the order of pictures for a RelatedID

	Count(relatedID [16]byte) (numbers uint64, err protocol.Error)
	Get(relatedID [16]byte, versionOffset uint64) (p Picture, err protocol.Error)
	Last(relatedID [16]byte) (p Picture, numbers uint64, err protocol.Error)
}

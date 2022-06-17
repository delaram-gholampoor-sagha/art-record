package protocol

type ContentAccess interface {
	ContentID() [16]byte    // content domain
	Access() Content_Access //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ContentAccess_StorageServices interface {
	Save(ca ContentAccess) protocol.Error

	Count(contentID [16]byte) (numbers uint64, err protocol.Error)
	Get(contentID [16]byte, versionOffset uint64) (ca ContentAccess, err protocol.Error)
	Last(contentID [16]byte) (ca ContentAccess, numbers uint64, err protocol.Error)
}

type Content_Access uint8

const (
	Content_Access_Unset    Content_Access = 0
	Content_Access_UserList Content_Access = (1 << iota)
	Content_Access_Public
	Content_Access_RegisteredUser
	Content_Access_PremiumUser // usually charge for monthly access.  /// other plan?? basic, ... /// role based??
)

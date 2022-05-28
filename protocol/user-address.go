package protocol

type UserAddress interface {
	UserID()    [16]byte `index-hash:"RecordID"` // UUID of User
	// TODO:::
	MobileNumber() string
	Status()    UserAddressStatus
	RequestID() [16]byte // user-request domain
}

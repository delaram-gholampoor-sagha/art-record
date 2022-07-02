package protocol

type UserOtherAppToken interface {
	UserID() [16]byte     // user domain
	OrgID() [16]byte      // user domain
	Token() string       // DelegateToken, AccessToken
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

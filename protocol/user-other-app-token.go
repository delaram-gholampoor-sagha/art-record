package protocol

type UserOtherAppToken interface {
	UserOtherAppID() [16]byte // user-other-app domain
	Token() string            // DelegateToken , AccessToken
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

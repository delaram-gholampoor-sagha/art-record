package protocol

// UserPublicKey indicate the domain record data fields.
type UserPublicKey interface {
	UserID() [16]byte             // user domain
	PublicKey() []byte            // suggest use DER format
	Issuer() [16]byte             // user domain. Use to get notify about status of the PK e.g. revoked notification, ...
	Status() UserPublicKey_Status //
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

// UserPublicKey_Status use to indicate UserPublicKey record status
type UserPublicKey_Status uint8

// UserPublicKey status
const (
	UserPublicKey_Status_Unset UserPublicKey_Status = iota
	UserPublicKey_Status_Active
	UserPublicKey_Status_Inactive
	UserPublicKey_Status_Blocked
	UserPublicKey_Status_Revoked
)

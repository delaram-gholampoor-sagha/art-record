package protocol

// UserRelationStatus indicate the domain record data fields.
type UserRelationStatus interface {
	UserID() [16]byte            // user domain
	SideID() [16]byte            // user domain
	Status() UserRelation_Status //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type UserRelationStatus_StorageServices interface {
	Save(gs UserRelationStatus) protocol.Error

	Count(userID, sideID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID, sideID [16]byte, versionOffset uint64) (gs UserRelationStatus, err protocol.Error)
	Last(userID, sideID [16]byte) (gs UserRelationStatus, numbers uint64, err protocol.Error)

	ListSides(userID [16]byte, offset, limit uint64) (sideIDs [][16]byte, numbers uint64, err protocol.Error)
	ListUsers(sideID [16]byte, offset, limit uint64) (userIDs [][16]byte, numbers uint64, err protocol.Error)
	
	FilterByStatus(userID [16]byte, status UserRelation_Status, offset, limit uint64) (sideIDs [][16]byte, numbers uint64, err protocol.Error)

	protocol.EventTarget
}

type UserRelation_Status Quiddity_Status

const (
	UserRelation_Status_Blocked = UserRelation_Status(Quiddity_Status_FreeFlag << iota)
	UserRelation_Status_Muted
)

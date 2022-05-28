package protocol


// UserRelation indicate the domain record data fields.
type UserRelation interface {
	UserID() [16]byte            // user-status domain
	SideID() [16]byte            // user-status domain
	Status() UserRelation_Status //
	// TODO::: save relation map by ids?
	RequestID() [16]byte         // user-request domain
}

type UserRelation_Status uint8

const (
	UserRelation_Status_Unset UserRelation_Status = iota
	UserRelation_Status_1stFriendshipCycle
	UserRelation_Status_2stFriendshipCycle
	UserRelation_Status_3stFriendshipCycle
	UserRelation_Status_4stFriendshipCycle
	UserRelation_Status_5stFriendshipCycle
	UserRelation_Status_6stFriendshipCycle
)

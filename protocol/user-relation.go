package protocol

// UserRelation indicate the domain record data fields.
type UserRelation interface {
	UserID() [16]byte        // user domain
	SideID() [16]byte        // user domain
	Relation() User_Relation //
	// TODO::: save relation map by ids?
	// Map() [][16]byte
	RequestID() [16]byte // user-request domain
}

type User_Relation uint8

const (
	User_Relation_Unset User_Relation = iota
	User_Relation_1stFriendshipCycle
	User_Relation_2stFriendshipCycle
	User_Relation_3stFriendshipCycle
	User_Relation_4stFriendshipCycle
	User_Relation_5stFriendshipCycle
	User_Relation_6stFriendshipCycle
)

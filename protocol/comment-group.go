package protocol

import (
	"time"

)

// CommentGroup indicate the comment-group domain record data fields.
type CommentGroup interface {
	GroupID() [16]byte                  // or ChannelID. Use to store and retrieve comments in time order. It can be UUID, URL hash or any random id.
	OwnerUserID() [16]byte              // user-status domain
	BlockedWords() []string             //
	GetPolicy() CommentGroup_Policy     // Visibility level
	SetPolicy() CommentGroup_Policy     //
	ForwardPolicy() CommentGroup_Policy //
	Time() time.Time                    // Save time
	RequestID() [16]byte                // user-request domain
}

type CommentGroup_Policy uint16

const (
	CommentGroup_Policy_Unset     CommentGroup_Policy = 0
	CommentGroup_Policy_JustOwner CommentGroup_Policy = (1 << iota)
	CommentGroup_Policy_GroupManagers
	CommentGroup_Policy_GroupMembers
	CommentGroup_Policy_Mentioned
	CommentGroup_Policy_AnyOne
	// Below policies are person type authorization
	CommentGroup_Policy_1stFriendshipCycle
	CommentGroup_Policy_2stFriendshipCycle
	CommentGroup_Policy_3stFriendshipCycle
	CommentGroup_Policy_4stFriendshipCycle
	CommentGroup_Policy_5stFriendshipCycle
	CommentGroup_Policy_6stFriendshipCycle
)

type CommentGroup_StorageServices interface {
	Save(cg CommentGroup) error

	Count(groupID [16]byte) (length uint64, err error)
	Get(groupID [16]byte, versionOffset uint64) (cg CommentGroup, err error)
	Last(groupID [16]byte) (cg CommentGroup, length uint64, err error)

	FindByUserID(ownerUserID [16]byte, offset, limit uint64) (ids [][16]byte, length uint64, err error)
}

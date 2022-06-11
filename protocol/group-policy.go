package protocol

// GroupPolicy indicate the domain record data fields.
type GroupPolicy interface {
	GroupID() [16]byte                 // group-status domain
	GetPolicy() GroupPolicy_Policy     // Visibility level
	SetPolicy() GroupPolicy_Policy     //
	ForwardPolicy() GroupPolicy_Policy //
	Time() protocol.Time               // Save time
	RequestID() [16]byte               // user-request domain
}

type GroupPolicy_StorageServices interface {
	Save(cg GroupPolicy) protocol.Error

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (cg GroupPolicy, err protocol.Error)
	Last(groupID [16]byte) (cg GroupPolicy, numbers uint64, err protocol.Error)

	FindByUserID(ownerUserID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)
}

type GroupPolicy_Policy uint16

const (
	GroupPolicy_Policy_Unset     GroupPolicy_Policy = 0
	GroupPolicy_Policy_JustOwner GroupPolicy_Policy = (1 << iota)
	GroupPolicy_Policy_GroupManagers
	GroupPolicy_Policy_GroupMembers
	GroupPolicy_Policy_Mentioned
	GroupPolicy_Policy_AnyOne
	// Below policies are person type authorization
	GroupPolicy_Policy_1stFriendshipCycle
	GroupPolicy_Policy_2stFriendshipCycle
	GroupPolicy_Policy_3stFriendshipCycle
	GroupPolicy_Policy_4stFriendshipCycle
	GroupPolicy_Policy_5stFriendshipCycle
	GroupPolicy_Policy_6stFriendshipCycle
)

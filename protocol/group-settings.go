package protocol

// GroupRole indicate the domain record data fields.
type GroupSettings interface {
	GroupID() [16]byte                     // group domain
	ArchiveThreadAfter() protocol.Duration //
	BlockedWords() []string                //
	Join() GroupSettings                   //
	Time() protocol.Time                   // save time
	RequestID() [16]byte                   // user-request domain
}

type GroupSettings_StorageServices interface {
	Save(gs GroupSettings) protocol.Error

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (gs GroupSettings, err protocol.Error)
	Last(groupID [16]byte) (gs GroupSettings, numbers uint64, err protocol.Error)
}

type GroupSettings_Join uint8

const (
	GroupSettings_Join_Unset  GroupSettings_Join = 0
	GroupSettings_Join_Invite GroupSettings_Join = (1 << iota)
	GroupSettings_Join_AddByMembers
)

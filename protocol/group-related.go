package protocol

// GroupRelated indicate the domain record data fields.
type GroupRelated interface {
	GroupID() [16]byte        // group-status domain
	RelatedGroupID() [16]byte // group-status domain. Use to categorize and graph groups
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}

package protocol

type Comment_Group struct {
	CommentGroupID [16]byte
	BlockedWords   string
	// options
	Status uint8
}


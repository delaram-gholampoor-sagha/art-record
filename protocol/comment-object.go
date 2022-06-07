
package protocol



// CommentObject indicate the comment-object domain record data fields.
// Object can be any such as image, video, voice, ...
type CommentObject interface {
	CommentID() [16]byte // comment domain
	ObjectID() [16]byte  // object domain
	Time() time.Time     // Save time
	RequestID() [16]byte // user-request domain
}

type CommentObject_StorageService interface {
	Save(co CommentObject) error

	Count(commentID [16]byte) (length uint64, err error)
	Get(commentID [16]byte, versionOffset uint64) (co CommentObject, err error)
	Last(commentID [16]byte) (co CommentObject, length uint64, err error)

	FindByObjectID(objectID [16]byte, offset, limit uint64) (ids [][16]byte, length uint64, err error)	
}



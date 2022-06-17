package protocol

type PictureRating interface {
	ObjectID() [16]byte         //
	Rating() PictureRating_Rate //
	Time() protocol.Time        // save time
	RequestID() [16]byte        // user-request domain
}

type PictureRating_Rate uint8

const (
// 0 : Suitable for display on all websites with any audience type.
// 1 : May contain rude gestures, provocatively dressed individuals, the lesser swear words, or mild violence.
// 2 : May contain such things as harsh profanity, intense violence, nudity, or hard drug use.
// 3 : May contain hardcore sexual imagery or extremely disturbing violence.
)

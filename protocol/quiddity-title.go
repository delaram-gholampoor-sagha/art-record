

package protocol


/*
	Storage Data Structure
*/

type QuiddityTitle interface {
	QuiddityID() [16]byte        // quiddity domain. Wil be unique in all languages
	Language() protocol.Language //
	Title() string               // Locale name in the Human world. It can be not unique in all quiddity contents.
	RequestID() [16]byte         // user-request domain
}

type QuiddityTitle_StorageServices interface {
	Save(q QuiddityTitle) (err error)

	Count(quiddityID [16]byte, language protocol.Language) (numbers uint64, err error)
	Get(quiddityID [16]byte, language protocol.Language, versionOffset uint64) (q QuiddityTitle, err error)
	Last(quiddityID [16]byte, language protocol.Language) (q QuiddityTitle, err error)

	ListQuiddityLanguages(quiddityID [16]byte, offset, limit uint64) (languages []protocol.Language, err error)
}

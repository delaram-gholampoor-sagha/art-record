package protocol

// QuiddityTitle indicate the domain record data fields.
type QuiddityTitle interface {
	QuiddityID() [16]byte        // quiddity domain. Wil be unique in all languages
	Language() protocol.Language //
	Title() string               // Locale name in the Human world. It can be not unique in all quiddity contents.
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type QuiddityTitle_StorageServices interface {
	Save(qt QuiddityTitle) (err protocol.Error)

	Count(quiddityID [16]byte, language protocol.Language) (numbers uint64, err protocol.Error)
	Get(quiddityID [16]byte, language protocol.Language, versionOffset uint64) (qt QuiddityTitle, err protocol.Error)
	Last(quiddityID [16]byte, language protocol.Language) (qt QuiddityTitle, numbers uint64, err protocol.Error)

	ListQuiddityLanguages(quiddityID [16]byte, offset, limit uint64) (languages []protocol.Language, numbers uint64, err protocol.Error)
}

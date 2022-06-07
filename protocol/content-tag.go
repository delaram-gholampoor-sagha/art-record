package protocol

type ContentTag interface {
	Tag() string         //
	ContentID() [16]byte // content domain
	Time() protocol.Time // Save time
}

type ContentTag_StorageServices interface {
	CreateTutorial(tutorial Tutorial) (Tutorial, error)
	UpdateTutorial(isPartial bool, tutorial Tutorial) (Tutorial, error)
	DeleteTutorial(tutorialID uint64) error
	FindTutorialByCategory(categoryname string) (Tutorial, error)
	FindTutorialByTitle(title string) (Tutorial, error)
}

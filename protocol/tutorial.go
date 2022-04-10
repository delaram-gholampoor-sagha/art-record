package protocol



type Tutorial struct {
	ID         uint64
	// QuiddityID [16]byte
	// PriceID         uint64
	// CategoryID      uint64
	// VideoID         uint64
	// PictureID       uint64
	RelatedID       [16]byte
	// sees it as comment ?
	Body            string
	Status          uint8
	Instructor_Name string
}

type TutorialServices interface {
	CreateTutorial(tutorial Tutorial) (Tutorial, error)
	UpdateTutorial(isPartial bool, tutorial Tutorial) (Tutorial, error)
	DeleteTutorial(tutorialID uint64) error
	FindTutorialByCategory(categoryname string) (Tutorial, error)
	FindTutorialByTitle(title string) (Tutorial, error)
}

package protocol

type tutorial_Content struct {
	ID            uint64 `json:"id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	Product_Price string `json:"product_price"`
	//published or saved ? or maybe other state ?
	State      string `json:"state"`
	UserID     uint64 `json:"user_id"`
	CategoryId uint64 `json:"Category_id"`
	Video_URL  string `json:"video_url"`
	Pic_URL    string `json:"pic_url"`
}

type Tutorial interface {
	CreateTutorial(tutorial Tutorial) (tutorial_Content, error)
	UpdateTutorial(isPartial bool, tutorial tutorial_Content) (tutorial_Content, error)
	DeleteTutorial(tutorialID uint64) error
	// GetTutorial(tutorialID uint64) (tutorial_Content, error)
	FindTutorialByCategory(categoryname string) (tutorial_Content, error)
	FindTutorialByTitle(title string) (tutorial_Content, error)
}

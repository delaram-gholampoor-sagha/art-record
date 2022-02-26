package protocol

import "time"

type User struct {
	ID          uint      `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	Status      string    `json:"status"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Avatar      string    `json:"avatar"`
}

// Create: bring (something) into existence
// Register: to record someone’s name or ownership of property on an official list
// ... To record or show a measurement, an amount, etc
// ... A book or record containing an official list of names, items, etc



// How do we interact with our data ?
type Business_Logic interface {
	// Identification
	GetUser()
	UpdateUser()
	GetUsers()
	SearchUser()

	// Authentication
	LogInUser()
	RegisterUser()

	// Authorization
	// libgo - 

	// blog
	GetPost()
	CreatePost()
	GetPosts()
	SearchPost()
	DeletePost()
	UpdatePost()

	SearchMusic()
	DeleteMusic()
	CreateMusic()
	GetMusic()
	GetMusics()

	CreateAlbum()
	DeleteAlbum()
	SearchAlbum()
	GetAlbums()
	GetAlbum()

	RegisterTag()
	UpdateTag()
	DeleteTag()
	SearchTags()

	// Invoice ??
	GetInvoice()
	GetInvoices()
	GetInvoicesByAdmin()
	RegisterInvoice()
	UpdateInvoice()

	RegisterReservation()
	SuccessReservation()
	DeleteReservation()
	UpdateReservation()
	GetReservation()
	GetReservations()

	// Event
	RegisterEvent()
	DeleteEvent()
	FindEventByTime()
	FindEventByLocation()
    GetEvents()

	CreateCategory()
	UpdateCategory()
	DeleteCategory()
	GetCategory()
	GetCategories()

	CreateTutorial()
	UpdateTutorial()
	DeleteTutorial()
	GetTutorial()
	GetTutotials()

	// ObjectStorage
	// There are some apis for it - we can separately read it or download it => 256kb
	// Taghsim e bar 
	// stree amazon  / storage-object libgo 
}

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
	LogInUser(email string , password string)
	RegisterUser(email string, phone_number string,full_name string, password string)

	// Authorization
	// libgo - 

	// blog
	GetPost()
	CreatePost()
	GetPosts(postID int , )
	SearchPost(tilte string , )
	DeletePost(postID int)
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

	// Find Invoice
    FindInvoiceByProductID()
	FindInvoiceByUserID()
    FindInvoiceByAcossiatedOwner()
	FindInvoiceByDateTime()
	FindInvoiceByState()

	// the state of our purchase ??? the products are staying in the cart with a time limitation
	// how should i present the state of Return from purchase ?
	// event's invoice
	// add item ? 


	// product price
    GetProductPrice()
	UpdatePrice()
	DeletePrice()
	DiscountedPrice()
	FindAllPricesByDate()


	// user's wallet


	// vouchers


	RegisterReservation()
	ChangeReservationState()
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
	// s3 amazon  / storage-object libgo 
}




package protocol


	// ماهیت یعنی ذات و چیستی یک شئ که دارای نام یا عنوان مشخص در زبان های انسانی و جدیدا در زبان های ماشین می باشد.
type Quiddity interface {
	ID() [16]byte               
	OrgID() [16]byte             
	Status() QuiddityStatus      
	RequestID() [16]byte        
}


// QuiddityStatus indicate Quiddity record status
type QuiddityStatus uint8

// Quiddity status
const (
	QuiddityStatusUnset = iota
	QuiddityStatusRegister
	QuiddityStatusUnRegister
	QuiddityStatusSuggestion
	QuiddityStatusBlocked
)

type QuiddityServices interface {

	// 	{
	// 		How:  protocol.IndexType_Hash_OneToMany,
	// 		What: "ObjectID",
	// 		For:  "ID",
	// 		Pair: []string{"Language"},
	// 	}, {
	// 		How:  protocol.IndexType_Hash_OneToMany,
	// 		What: "Language",
	// 		For:  "ID",
	// 	}, {
	// 		How:  protocol.IndexType_Hash_OneToMany,
	// 		What: "ID",
	// 		For:  "OrgID",
	// 	}, {
	// 		How:  protocol.IndexType_Hash_OneToMany,
	// 		What: "ID",
	// 		For:  "URI",
	// 		If:   "URI",
	// 	}, {
	// 		How:  protocol.IndexType_Text,
	// 		What: "ID",
	// 		For:  "Title",
	// 		If:   "Title",
	// 	},
}

/*
	Services
*/

const QuiddityServiceFindByOrgID = "urn:giti:quiddity.sabz.city:service:find-by-org-id"

// Service request and response shape
type QuiddityServiceFindByOrgIDRequest interface {
	OrgID() [32]byte
	Offset() uint64
	Limit() uint64
}
type QuiddityServiceFindByOrgIDResponse interface {
	IDs() [][32]byte
}

/*
	Errors
*/

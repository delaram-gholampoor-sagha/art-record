package protocol

// UserName indicate the domain record data fields.
type UserName interface {
	UserID() [16]byte        // user domain
	Username() string        // It is not replace of user ID. It usually use to find user by their friends in more human friendly manner.
	Status() UserName_Status //
	Time() protocol.Time     // save time
	RequestID() [16]byte     // user-request domain
}

type UserName_StorageServices interface {
	Save(un UserName) (numbers uint64, err protocol.Error)

	Count(userID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (un UserName, numbers uint64, err protocol.Error)

	FindByUsername(username string, offset, limit uint64) (userID [][16]byte, numbers uint64, err protocol.Error)
}

// UserName_Status indicate UserName record status
type UserName_Status uint8

// UserName status
const (
	UserName_Status_Unset    UserName_Status = 0           //
	UserName_Status_Register UserName_Status = (1 << iota) //= (1 << 0)
	UserName_Status_Remove                                 //= (1 << 1)
	UserName_Status_Blocked                                //= (1 << 2)
	UserName_Status_Reserved                               //= (1 << 3)

	UserName_Status_Active
	UserName_Status_Inactive
)

/*
	Business Services
*/

type UserName_Service_Register_Request interface {
	UserID() [16]byte // TODO::: get user id from connection??
	Username() string
}

type UserName_Service_Get_Request interface {
	UserID() [16]byte
	VersionOffset() uint64
}
type UserName_Service_Get_Response interface {
	UserName
}

type UserName_Service_GetStatus_Request interface {
	Username() string
}
type UserName_Service_GetStatus_Response interface {
	Status() UserName_Status
}

// Admins services

type UserName_Service_ChangeStatus_Request interface {
	Username() string
	Status() UserName_Status
}

/*
	Errors
*/

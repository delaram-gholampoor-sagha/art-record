package protocol

// UserReferent indicate the domain record data fields.
type UserReferent interface {
	UserID() [16]byte         // user domain
	ReferentUserID() [16]byte // user domain
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type UserReferent_StorageServices interface {
	Save(ur UserReferent) (numbers uint64, err protocol.Error)

	Count(userID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (ur UserReferent, numbers uint64, err protocol.Error)

	FindByReferentUserID(referentUserID [16]byte, offset, limit uint64) (userIDs [][16]byte, numbers uint64, err protocol.Error)
}

type (
	UserReferent_Service_Register_Request interface{
		UserID() [16]byte         
	  ReferentUserID() [16]byte 
	}
	
	UserReferent_Service_Register_Response interface{
		Numbers() uint64
	}
	
	UserReferent_Service_Count_Request interface{
		UserID() [16]byte
	}
	
	UserReferent_Service_Count_Response interface{
		Numbers() uint64
	}
	UserReferent_Service_Get_Request interface{
		UserID() [16]byte
		VersionOffset() uint64
	}
	
	UserReferent_Service_Get_Response interface{
		UserReferent
		Numbers() uint64
	}
	
	
	UserReferent_Service_FindByReferentUserID_Request interface{
		ReferentUserID() uint64
		Offset() uint64
		Limit() uint64
	}
	
	UserReferent_Service_FindByReferentUserID_Response interface{
		UserIDs() [][16]byte
		Numbers() uint64
	}
	
	
)
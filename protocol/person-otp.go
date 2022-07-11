package protocol

import "../libgo/protocol"

// PersonOTP indicate the domain record data fields.
type PersonOTP interface {
	PersonID() [16]byte   // user domain
	OTPPattern() []byte   // https://tools.ietf.org/html/rfc6238
	OTPAdditional() int32 // easy to remember and must be 2 to 7 digit. https://en.wikipedia.org/wiki/Personal_identification_number
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type PersonOTP_StorageServices interface {
	Save(po PersonOTP) (err protocol.Error)

	Count(personID [16]byte) (numbers uint64, err protocol.Error)
	Get(personID [16]byte, versionOffset uint64) (po PersonOTP, err protocol.Error)

}

type (
	PersonOTP_Service_Register_Request interface {
		PersonID() [16]byte   
		OTPPattern() []byte   
		OTPAdditional() int32	
	}
	PersonOTP_Service_Register_Response interface {
		 Nv() protocol.NumberOfVersion
	}
	
	PersonOTP_Service_Count_Request interface {
		PersonID() [16]byte
		
	
	}
	PersonOTP_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	PersonOTP_Service_Get_Request interface {
		PersonID() [16]byte    
		VersionOffset() uint64
	}
	PersonOTP_Service_Get_Response interface {
		PersonOTP
		Nv() protocol.NumberOfVersion
	}
	
	
)

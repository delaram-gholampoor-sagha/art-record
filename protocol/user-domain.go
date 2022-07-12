package protocol

import "../libgo/protocol"

type UserDomain interface {
	Domain() string            // unique name in the computer world
	UserID() [16]byte          // user domain
	Status() UserDomain_Status //
	Time() protocol.Time       // save time
	RequestID() [16]byte       // user-request domain
}

type UserDomain_StorageServices interface {
	Save(ud UserDomain) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(domain string) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(domain string, versionOffset uint64) (ud UserDomain, nv protocol.NumberOfVersion, err protocol.Error)

	FindByUserID(userID [16]byte, offset, limit uint64) (domains []string, nv protocol.NumberOfVersion, err protocol.Error)
}

// UserDomain_Status indicate UserDomain record status
type UserDomain_Status uint8

// UserDomain statuses
const (
	UserDomain_Status_Unset UserDomain_Status = iota
	UserDomain_Status_Active
	UserDomain_Status_Inactive
	UserDomain_Status_Blocked
)

const UserDomain_Service_Register = "urn:giti:user-domain.sabz.city:service:register"
const UserDomain_Service_ChangeDomain = "urn:giti:user-domain.sabz.city:service:change-domain"
const UserDomain_Service_Close = "urn:giti:user-domain.sabz.city:service:close"
const UserDomain_Service_Reopen = "urn:giti:user-domain.sabz.city:service:reopen"
const UserDomain_Service_Block = "urn:giti:user-domain.sabz.city:service:block"
const UserDomain_Service_Unblock = "urn:giti:user-domain.sabz.city:service:unblock"

/*
	Errors
*/

type (
	UserDomain_Service_Register_Request interface {
		Domain() string
		UserID() [16]byte
		Status() UserDomain_Status
	}

	UserDomain_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

	UserDomain_Service_Count_Request interface {
		Domain() string
	}

	UserDomain_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
	UserDomain_Service_Get_Request interface {
		Domain() string
		VersionOffset() uint64
	}

	UserDomain_Service_Get_Response interface {
		UserDomain
		NumberOfVersion() protocol.NumberOfVersion
	}

	UserDomain_Service_FindByUserID_Request interface {
		UserID() uint64
		Offset() uint64
		Limit() uint64
	}

	UserDomain_Service_FindByUserID_Response interface {
		Domains() []string
		NumberOfVersion() protocol.NumberOfVersion
	}
)
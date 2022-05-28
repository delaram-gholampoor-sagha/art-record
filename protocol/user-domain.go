package protocol

import (
	"../libgo/protocol"
)

type UserDomain interface {
	Domain() string            // unique name in the computer world
	UserID() [16]byte          // user-status domain
	Status() UserDomain_Status //
	Time() protocol.Time       // Save time
	RequestID() [16]byte       // user-request domain
}

type UserDomain_StorageServices interface {
	Save(ud UserDomain) (err protocol.Error)

	Count(domain string) (numbers uint64, err protocol.Error)
	Get(domain string, versionOffset uint64) (ud UserDomain, err protocol.Error)
	Last(domain string) (ud UserDomain, numbers uint64, err protocol.Error)

	FindByUserID(userID [16]byte offset, limit uint64) (domains []string, numbers uint64, err protocol.Error)
}

// UserDomain_Status indicate UserDomain record status
type UserDomain_Status uint8

// UserDoamin status
const (
	UserDomain_Status_Unset UserDomain_Status = iota
	UserDomain_Status_Active
	UserDomain_Status_Inactive
	UserDomain_Status_Blocked
)

// ShortDetail returns localize short and long details for given status code
func (oas UserDomain_Status) ShortDetail() (short string) {
	switch protocol.AppLanguage {
	case protocol.LanguageEnglish:
		switch oas {
		case UserDomain_StatusUnset:
			short = "Unset"
		case UserDomain_StatusRegister:
			short = "Register"
		case UserDomain_StatusChangedDomian:
			short = "Change Domian"
		case UserDomain_StatusClosed:
			short = "Closed"
		case UserDomain_StatusBlocked:
			short = "Blocked"
		default:
			short = "Invalid Status"
		}
	}
	return
}

// LongDetail returns localize long details for given status code
func (oas UserDomain_Status) LongDetail() (long string) {
	switch protocol.AppLanguage {
	case protocol.LanguageEnglish:
		switch oas {
		case UserDomain_StatusUnset:
			long = "Record status didn't set yet"
		case UserDomain_StatusRegister:
			long = "User register and start in this society"
		case UserDomain_StatusChangedDomian:
			long = "User change its domian to other domain from last change"
		case UserDomain_StatusClosed:
			long = "User closed and can't & don't have any further activity"
		case UserDomain_StatusBlocked:
			long = "User had been inactive and can't be use now"
		default:
			long = "Given status code is not valid for this type of record"
		}
	}
	return
}

/*
	Services
*/

const UserDomain_Service_Register = "urn:giti:user-domain.sabz.city:service:register"
const UserDomain_Service_ChangeDomain = "urn:giti:user-domain.sabz.city:service:change-domain"
const UserDomain_Service_Close = "urn:giti:user-domain.sabz.city:service:close"
const UserDomain_Service_Reopen = "urn:giti:user-domain.sabz.city:service:reopen"
const UserDomain_Service_Block = "urn:giti:user-domain.sabz.city:service:block"
const UserDomain_Service_Unblock = "urn:giti:user-domain.sabz.city:service:unblock"

/*
	Errors
*/

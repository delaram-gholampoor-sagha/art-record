/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// UserPublicKey indicate the domain record data fields.
type UserPublicKey interface {
	UserID() [16]byte             // user domain
	PublicKey() []byte            // suggest use DER format
	Issuer() [16]byte             // user domain. Use to get notify about status of the PK e.g. revoked notification, ...
	Status() UserPublicKey_Status //
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

// UserPublicKey_Status use to indicate UserPublicKey record status
type UserPublicKey_Status uint8

// UserPublicKey status
const (
	UserPublicKey_Status_Unset UserPublicKey_Status = iota
	UserPublicKey_Status_Active
	UserPublicKey_Status_Inactive
	UserPublicKey_Status_Blocked
	UserPublicKey_Status_Revoked
)


type UserPublicKey_StorageServices interface {
	Save(up UserPublicKey) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(userID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (up UserPublicKey, nv protocol.NumberOfVersion, err protocol.Error)
}



type (
	UserPublicKey_Service_Register_Request interface{
		UserID() [16]byte             
		PublicKey() []byte            
		Issuer() [16]byte             
		Status() UserPublicKey_Status 
	}
	
	UserPublicKey_Service_Register_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (

		UserPublicKey_Service_Count_Request interface{
		UserID() [16]byte
	}
	
	UserPublicKey_Service_Count_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
	UserPublicKey_Service_Get_Request interface{
		UserID() [16]byte
		VersionOffset() uint64
	}
	
	UserPublicKey_Service_Get_Response interface{
		UserPublicKey
		NumberOfVersion() protocol.NumberOfVersion
	}

)


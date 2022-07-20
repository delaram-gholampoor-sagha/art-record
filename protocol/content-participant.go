/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ContentParticipant indicate the domain record data fields.
type ContentParticipant interface {
	ContentID() [16]byte           // content domain
	Type() ContentParticipant_Type //
	UserID() [16]byte              // user domain
	Time() protocol.Time           // save time
	RequestID() [16]byte           // user-request domain
}

// TODO::: immutable?? if wrong data save for a content??
type ContentParticipant_StorageServices interface {
	Save(cp ContentParticipant) protocol.Error

	Count(contentID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(contentID [16]byte, versionOffset uint64) (cp ContentParticipant,nv protocol.NumberOfVersion ,err protocol.Error)
	

	FindByUserID(userID [16]byte, offset, limit uint64) (contentIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

// ContentParticipant_Type indicate participant type or role
type ContentParticipant_Type uint32

// ContentParticipant types
const (
	ContentParticipant_Type_Unset ContentParticipant_Type = iota

	// Content_Type_Music
	ContentParticipant_Type_SongWriter
	ContentParticipant_Type_Composer
	ContentParticipant_Type_Regulator
	ContentParticipant_Type_MusicSupplier

	// Content_Type_Movie
	ContentParticipant_Type_Director

	// Content_Type_Game
	ContentParticipant_Type_SportTeam
)



type (
	ContentParticipant_Service_Register_Request interface {
		ContentID() [16]byte           
  	Type() ContentParticipant_Type 
  	UserID() [16]byte              
	}
	
	ContentParticipant_Service_Register_Response = protocol.NumberOfVersion

)

type (
		ContentParticipant_Service_Count_Request interface {
		ContentID() [16]byte
	}
	
	ContentParticipant_Service_Count_Response = protocol.NumberOfVersion
	
)



type (
	ContentParticipant_Service_Get_Request interface {
		ContentID() [16]byte
		versionOffset() uint64
	}
	
	ContentParticipant_Service_Get_Response1 = 	ContentParticipant
	ContentParticipant_Service_Get_Response2 = protocol.NumberOfVersion
	
)


type (
	ContentParticipant_Service_FindByUserID_Request interface {
		UserID() [16]byte 
		Offset() uint64
		Limit() uint64
	}
	
	ContentParticipant_Service_FindByUserID_Response1 interface {
		ContentIDs() [][16]byte
	}

	ContentParticipant_Service_FindByUserID_Response2 =  protocol.NumberOfVersion
	
)
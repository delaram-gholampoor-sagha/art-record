/* For license and copyright information please see LEGAL file in repository */

package art

// QuiddityTitle indicate the domain record data fields.
type QuiddityTitle interface {
	QuiddityID() [16]byte        // quiddity domain. Wil be unique in all languages
	Language() protocol.Language //
	Title() string               // Locale name in the Human world. It can be not unique in all quiddity contents.
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type QuiddityTitle_StorageServices interface {
	Save(qt QuiddityTitle) (nv protocol.NumberOfVersion ,err protocol.Error)

	Count(quiddityID [16]byte, language protocol.Language) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(quiddityID [16]byte, language protocol.Language, versionOffset uint64) (qt QuiddityTitle, nv protocol.NumberOfVersion ,err protocol.Error)
	

	ListQuiddityLanguages(quiddityID [16]byte, offset, limit uint64) (languages []protocol.Language, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	QuiddityTitle_Service_Register_Request interface {
		QuiddityID() [16]byte
		Languages() protocol.Language
		Title()string
	}
	
	QuiddityTitle_Service_Register_Response = protocol.NumberOfVersion
)

type (
	QuiddityTitle_Service_Count_Request interface {
		QuiddityID() [16]byte
		Language() protocol.Language
	}
	
	QuiddityTitle_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	QuiddityTitle_Service_Get_Request interface {
		QuiddityID() [16]byte
		Language() protocol.Language
		VersionOffset() uint64 
	}
	
	QuiddityTitle_Service_Get_Response1 = QuiddityTitle

	QuiddityTitle_Service_Get_Response2 = protocol.NumberOfVersion
	
)


type (
	QuiddityTitle_Service_ListQuiddityLanguages_Request interface {
		QuiddityID() [16]byte 
		Offset() uint64
		Limit() uint64
	}
	
	QuiddityTitle_Service_ListQuiddityLanguages_Response1 interface {
		Languages() []protocol.Language 
	}

	QuiddityTitle_Service_ListQuiddityLanguages_Response2 =  protocol.NumberOfVersion
)

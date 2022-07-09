package protocol

// QuiddityIcon indicate the domain record data fields.
// Icon or emoji use in many domains e.g. category, departments, ...
type QuiddityIcon interface {
	QuiddityID() [16]byte // quiddity domain
	Icon() [16]byte       // object domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type QuiddityIcon_StorageServices interface {
	Save(qi QuiddityIcon) (err protocol.Error)

	Count(quiddityID [16]byte) (numbers uint64, err protocol.Error)
	Get(quiddityID [16]byte, versionOffset uint64) (qi QuiddityIcon, err protocol.Error)

}

type (
	QuiddityIcon_Service_Register_Request interface {
		QuiddityID() [16]byte
		Icon() [16]byte
	}
	
	QuiddityIcon_Service_Register_Response interface {
		Numbers() [16]byte
	}
	
	QuiddityIcon_Service_Count_Request interface {
		QuiddityID() [16]byte
	}
	QuiddityIcon_Service_Count_Response interface {
		Numbers() uint64
	}
	
	
	QuiddityIcon_Service_Get_Request interface {
		QuiddityID() [16]byte
		VersionOffset() uint64
	}
	QuiddityIcon_Service_Get_Response interface {
		QuiddityIcon
	}
	
)

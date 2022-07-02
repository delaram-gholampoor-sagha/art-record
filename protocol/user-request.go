package protocol

// UserRequest use to prove how other records created as a chain to other records or standalone data.
type UserRequest interface {
	RequestID() [16]byte        //
	DomainID() uint64           // Domain MediaTypeID e.g. product, content, ... domain
	PastRequestID() [16]byte    // user-request domain. to chain related requests that create||update the same record.
	UserID() [16]byte           // user domain
	UserConnectionID() [16]byte // Store to remember request belong to which Connection/AppInstance.
	// ClientType() string // persiaos, web, android, ios, ...
	Geo()                // location of user or post
	ServiceID() uint64   //
	CompressID() uint64  // gzip, ...
	Request() []byte     // Decode by service request structure and compress type.
	Time() protocol.Time // Request time or save Time of the request not the created record by this record.
	Signature() []byte   // Store to prove user requested to set||update this record.
}

type UserRequest_StorageServices interface {
	Save(q UserRequest) (numbers uint64, err protocol.Error)

	Count(requestID [16]byte) (numbers uint64, err protocol.Error)
	Get(requestID [16]byte, versionOffset uint64) (ur UserRequest, numbers uint64, err protocol.Error)

	FindByDomain(domainID uint64, offset, limit uint64) (requestIDs [][16]byte, numbers uint64, err protocol.Error)
	FindByUser(userID [16]byte, offset, limit uint64) (requestIDs [][16]byte, numbers uint64, err protocol.Error)
}

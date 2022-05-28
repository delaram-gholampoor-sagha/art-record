package protocol

import (
	"../libgo/protocol"
)

/*
	Storage Data Structure
*/

// UserRequest use to prove how other records created as a chain to other records or standalone data.
type UserRequest interface {
	ID() [16]byte               // RequestID
	PastRequestID() [16]byte    // user-request domian. to chain related requests that create||update the same record.
	UserID() [16]byte           // user-status domain
	UserConnectionID() [16]byte // Store to remember request belong to which Connection/AppInstance.
	// ClientType() string // persiaos, web, android, ios, ...
	Geo()                // location of user or post
	ServiceID() uint64   //
	CompressID() uint64  // gzip, ...
	Request() []byte     // Decode by service request structure and compress type.
	Time() protocol.Time // Request time or save Time of the request not the created record by this record.
	Signature() []byte   // Store to prove user requested to set||update this record.
}

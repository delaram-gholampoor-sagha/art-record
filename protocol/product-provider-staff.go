package protocol

// ProductProvider indicate the domain record data fields.
type ProductProviderStaff interface {
	ProductID() [16]byte // product domain
	StaffID() [16]byte   // staff-status domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

// related UserIDs by RoleID use to get Coordinate(domain)

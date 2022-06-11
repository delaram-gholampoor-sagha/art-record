package protocol
type Building interface {
	BuildingID() [16]byte // quiddity domain
	AreaID() [16]byte     // area domain
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}

package protocol

type CoordinateAnalytic interface {
	CoordinateID() [16]byte // quiddity||user||invoice||area||building||building-location||staff||... domain
	Distance() uint64       // km or ...
	AverageSpeed() uint64   // km/h or ...
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

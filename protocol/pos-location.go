package protocol

type PosLocation interface {
	PosID() [16]byte              // pos domain
	BuildingLocationID() [16]byte // building-location domain
	Time() protocol.Time          // Save time
	RequestID() [16]byte          // user-request domain
}

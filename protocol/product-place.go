package protocol


// ProductPlace indicate service provides in specific place e.g. stations(bus, train, airport, ...), ...
type ProductPlace interface {
	ProductID() [16]byte          // product domain.
	BuildingLocationID() [16]byte // building-location domain. use also for Coordinate domain
	Time() protocol.Time          // Save time
	RequestID() [16]byte          // user-request domain
}

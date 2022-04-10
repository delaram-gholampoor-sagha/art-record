package protocol

// urn:isbn:number
type Quiddity_URI interface {
	ID() [16]byte
	URI() [16]byte // Locale name in the Computer world.
	Status() uint8
	RequestID() [16]byte
}


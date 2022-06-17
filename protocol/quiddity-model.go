package protocol

// QuiddityModel indicate the domain record data fields.
// Model is the 3D model of the quiddity
type QuiddityModel interface {
	QuiddityID() [16]byte // quiddity domain
	ObjectID() [32]byte   // object domain. ThingModel
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

// ProductAttribute_Type_MetricUnit // Number, Gram, Liter, Second

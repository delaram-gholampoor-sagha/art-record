package protocol

type ProductTimingUTC interface {
	ProductID() [16]byte         // product domain
	Minimum() protocol.Duration  // Minimum time duration for this service
	Expected() protocol.Duration // Expected time duration for this service
	Maximum() protocol.Duration  // Maximum time duration for this service
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}

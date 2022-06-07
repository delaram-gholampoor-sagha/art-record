

package protocol


type CategorySubstitute interface {
	CategoryID() [16]byte   // category domain
	Priotry() int32         // Substitute priotry
	SubstituteID() [16]byte // category domain
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

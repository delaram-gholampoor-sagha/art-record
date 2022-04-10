package protocol

type quiddity_Language struct {
	ID          [16]byte
	Title       string
	Language    string
	Status      uint8
	RequestedID [16]byte
}

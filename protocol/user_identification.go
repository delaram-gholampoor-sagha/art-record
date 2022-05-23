package protocol

// تعیین هویت 



type User interface {
	ID()           [16]byte
}



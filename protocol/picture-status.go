package protocol

import "time"



type PictureStatus interface {
	ObjectID() [16]byte           //
	Status() PictureStatus_Status //
	Time() time.Time              // Save time
	RequestID() [16]byte          // user-request domain
}

type PictureStatusServices_StorageServices interface {
	
}

type PictureStatus_Status uint8

const (
	PictureStatus_Status_Unset PictureStatus_Status = iota
	PictureStatus_Status_Registered
	PictureStatus_Status_Hidden
	PictureStatus_Status_Deleted
	PictureStatus_Status_Blocked
)

package protocol

import "time"

type Invoice_Trans struct {
	ID                [16]byte
	SenderAccountID   [16]byte
	RecieverAccountID [16]byte
	Amount            int64
	Status            uint8
	Issue_At          time.Time
}

type Inoice_Trans_Service interface {

}

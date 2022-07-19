/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// FinancialAccountID indicate the domain record data fields.
// It is about settlement reference internal ID like bank account ID.
type FinancialAccountID interface {
	AccountID() [16]byte // financial-account domain
	ID() string          //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}


type FinancialAccountID_StorageServices interface {
	Save(fai FinancialAccountID) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(accountID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fai FinancialAccountID, nv protocol.NumberOfVersion, err protocol.Error)

	FindByID(id string) (accountIDs [][16]byte, err protocol.Error)
}

type (
		FinancialAccountID_Service_Register_Request interface {
		AccountID() [16]byte
		ID() string
	}

	FinancialAccountID_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
	FinancialAccountID_Service_Get_Request interface {
		AccountID() [16]byte
		VersionOffset() uint64
	}
	FinancialAccountID_Service_Get_Response interface {
		FinancialAccountID
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
	FinancialAccountID_Service_FindByID_Request interface {
		ID() string
	}
	FinancialAccountID_Service_FindByID_Response interface {
		AccountIDs() [][16]byte
	}

)
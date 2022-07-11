package protocol

import "../libgo/protocol"

type StaffHireStatus interface {
	UserID() [16]byte         // user domain
	Status() StaffHire_Status //
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type StaffHireStatus_StorageServices interface {
	Save(shs StaffHireStatus) protocol.Error

	Count(userID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (shs StaffHireStatus, err protocol.Error)

	// FilterByStatus(status StaffHire_Status, offset, limit uint64) (userIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	// protocol.EventTarget
}

type StaffHire_Status Quiddity_Status

const (
	StaffHire_Status_Requester = StaffHire_Status(Quiddity_Status_FreeFlag << iota)
	StaffHire_Status_Referral
	StaffHire_Status_Talent // high skill match
	StaffHire_Status_NeedHRInterview
	StaffHire_Status_ApproveForHRInterview
	StaffHire_Status_NeedTechnicalInterview
	StaffHire_Status_ApproveForTechnicalInterview
	StaffHire_Status_NeedForManagerInterview
	StaffHire_Status_ApproveForManagerInterview
	StaffHire_Status_NeedForChiefInterview
	StaffHire_Status_ApproveForChiefInterview
	StaffHire_Status_NeedForBoardInterview
	StaffHire_Status_ApproveForBoardInterview

	// if not accepted or hired
	StaffHire_Status_LockFor1Month
	StaffHire_Status_LockFor3Month
	StaffHire_Status_LockFor6Month
	StaffHire_Status_LockFor12Month

	// if accepted
	StaffHire_Status_Designation
	StaffHire_Status_Negotiation
	StaffHire_Status_Reserved
	StaffHire_Status_ContractProposal
)

type (
	StaffHireStatus_Service_Register_Request interface {
		UserID() [16]byte
		Status() StaffHire_Status
	}

	StaffHireStatus_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}

	StaffHireStatus_Service_Count_Request interface {
		StaffID() [16]byte
	}

	StaffHireStatus_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	StaffHireStatus_Service_Get_Request interface {
		StaffID() [16]byte
		VersionOffset() uint64
	}

	StaffHireStatus_Service_Get_Response interface {
		StaffHireStatus
	}

	StaffHireStatus_Service_Last_Request interface {
		StaffID() [16]byte
	}

	StaffHireStatus_Service_Last_Response interface {
		StaffHireStatus
		Nv() protocol.NumberOfVersion
	}
)
package protocol

import "../libgo/protocol"

type StaffHireStatus interface {
	ID() [16]byte             // user-status domain
	Status() StaffHire_Status //
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}

type StaffHire_Status uint8

const (
	StaffHire_Status_Unset StaffHire_Status = iota

	StaffHire_Status_Talent
	StaffHire_Status_NeedHRInterview
	StaffHire_Status_ApproveForHRInterview
	StaffHire_Status_NeedTechnicalInterview
	StaffHire_Status_ApproveForTechnicalInterview
	StaffHire_Status_NeedForManagerInterview
	StaffHire_Status_ApproveForManagerInterview
	StaffHire_Status_NeedForChiefInterview
	StaffHire_Status_ApproveForChiefInterview

	StaffHire_Status_Designation
	StaffHire_Status_Negotiation
	StaffHire_Status_Reserved
	StaffHire_Status_ContractProposal

	StaffHire_Status_Hire
	StaffHire_Status_Fire
	StaffHire_Status_End // death, end of contract, ...
)

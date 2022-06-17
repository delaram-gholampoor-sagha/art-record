package protocol


type ContractExpertOpinion interface {
	ContractID() [16]byte     // contract domain
	ExpertID() [16]byte       //
	Opinion() string          //
	Progress() math.PerMyriad //
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

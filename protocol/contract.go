package protocol

type Contract interface {
	ContractID() [16]byte              // contract domain
	ContractTemplateID() [16]byte      // contract-template domain
	ContractTemplateVersion() uint64   // VersionOffset
	Data() map[Contract_KeyType]string //
	Time() protocol.Time               // save time
	RequestID() [16]byte               // user-request domain
}

type Contract_StorageServices interface {
	Save(c Contract) (numbers uint64, err protocol.Error)

	Count(contractID [16]byte) (numbers uint64, err protocol.Error)
	Get(contractID [16]byte, versionOffset uint64) (c Contract, numbers uint64, err protocol.Error)

	FindByContractTemplate(contractTemplateID [16]byte, offset, limit uint64) (contractIDs [][16]byte, numbers uint64, err protocol.Error)
	FindByData(data string, offset, limit uint64) (contractIDs [16]byte, numbers uint64, err protocol.Error)
}

type Contract_KeyType uint32

const (
	Contract_KeyType_Unset Contract_KeyType = iota

	Contract_KeyType_FirstSide
	Contract_KeyType_SecondSide
	Contract_KeyType_ThirdSide
	Contract_KeyType_Witness
	Contract_KeyType_Arbitrator
	Contract_KeyType_Expert

	Contract_KeyType_Mass
	Contract_KeyType_Density

	Contract_KeyType_WeeklyRent
	Contract_KeyType_MonthlyRent

	Contract_KeyType_DueTo
)

// Service: ایفای تعهد

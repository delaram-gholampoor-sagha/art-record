package protocol


type ContractTemplateStatus interface {
	ContractTemplateID() [16]byte    // contract-template domain
	Status() ContractTemplate_Status //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}

type ContractTemplateStatus_StorageServices interface {
	Save(cts ContractTemplateStatus) (numbers uint64, err protocol.Error)

	Count(contractTemplateID [16]byte) (numbers uint64, err protocol.Error)
	Get(contractTemplateID [16]byte, versionOffset uint64) (cts ContractTemplateStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status Content_Status, offset, limit uint64) (contractTemplateIDs [][16]byte, numbers uint64, err protocol.Error)
}

type ContractTemplate_Status Quiddity_Status

const (
	ContractTemplate_Status_Draft = ContractTemplate_Status(Quiddity_Status_FreeFlag << iota)
)

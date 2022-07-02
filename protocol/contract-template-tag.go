package protocol


// ContractTemplateTag indicate the domain record data fields.
// Tag is any text with hashtag in the contract template text, can't use by old fashion way to assign by users.
type ContractTemplateTag interface {
	ContractTemplateID() [16]byte       // contract-template domain
	Tag() string                        //
	Status() ContractTemplateTag_Status //
	Time() protocol.Time                // save time
}

type ContractTemplateTag_StorageServices interface {
	Save(ct ContractTemplateTag) (numbers uint64, err protocol.Error)

	Count(contractTemplateID [16]byte) (numbers uint64, err protocol.Error)
	Get(contractTemplateID [16]byte, versionOffset uint64) (ct ContractTemplateTag, numbers uint64, err protocol.Error)

	ListTags(contractTemplateID [16]byte, offset, limit uint64) (tags []string, numbers uint64, err protocol.Error)
	ListContractTemplates(tag string, offset, limit uint64) (contractTemplateIDs [][16]byte, numbers uint64, err protocol.Error)

	FilterByStatus(contractTemplateID [16]byte, status ContractTemplateTag_Status, offset, limit uint64) (tags []string, numbers uint64, err protocol.Error)
}

type ContractTemplateTag_Status Quiddity_Status

const (
// ContractTemplateTag_Status_ = ContractTemplateTag_Status(Quiddity_Status_FreeFlag << iota)
)

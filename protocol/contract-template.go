package protocol

type ContractTemplate interface {
	ContractTemplateID() [16]byte // contract-template domain
	Terms() []string              // HTML, CSS, JS base, Solidity PL is other option
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type ContractTemplate_StorageServices interface {
	Save(ct ContractTemplate) (numbers uint64, err protocol.Error)

	Count(contractTemplateID [16]byte) (numbers uint64, err protocol.Error)
	Get(contractTemplateID [16]byte, versionOffset uint64) (ct ContractTemplate, numbers uint64, err protocol.Error)

	FindByHashTag()
}

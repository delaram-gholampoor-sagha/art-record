package protocol


type ContractTemplate interface {
	ContractTemplateID() [16]byte // contract-template domain
	Terms() []string              // HTML, CSS, JS base, Solidity PL is other option
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

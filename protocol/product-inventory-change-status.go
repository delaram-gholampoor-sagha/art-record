package protocol

type ProductInventoryChangeStatus interface {
	ProductInventoryChangeID() [16]byte           // product-inventory-change domain
	Status() ProductInventoryChange_Status // overall
	Time() protocol.Time                   // Save time
	RequestID() [16]byte                   // user-request domain
}

type ProductInventoryChange_Status uint8

const (
	ProductInventoryChange_Status_Unset ProductInventoryChange_Status = iota
	ProductInventoryChange_Status_Register

	ProductInventoryChange_Status_ChangeDCRequested
	ProductInventoryChange_Status_ChangeDCApproved
	ProductInventoryChange_Status_Void

	ProductInventoryChange_Status_PreSale // use in budget analysis and also can be trade

	// ManagerApprove
	// WarehouseApprove
	// ActiveWithApprove
	// ActiveWithoutApprove
	// ActiveWithOrderer
	// Inactive
	// Normal
	// Reject
)

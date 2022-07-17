package voucher

type Voucher_Type uint32

const (
	Voucher_Type_Unset Voucher_Type = 0
	Voucher_Type_Area  Voucher_Type = (1 << iota)
	Voucher_Type_Campaign
	Voucher_Type_Category
	Voucher_Type_Duration
	Voucher_Type_Group
	Voucher_Type_Invoice
	Voucher_Type_Location
	Voucher_Type_POS
	Voucher_Type_Product
	Voucher_Type_SourceInvoice
	Voucher_Type_Usage    // limit to usage number
	Voucher_Type_UserType // limit to user type
	Voucher_Type_User     // limit to specific users
)

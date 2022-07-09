package protocol

// VoucherUsage indicate the domain record data fields.
type VoucherUsage interface {
	VoucherID() [16]byte    // product domain
	MaximumInvoice() uint64 //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type VoucherUsage_StorageServices interface {
	Save(vu VoucherUsage) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vu VoucherUsage, numbers uint64, err protocol.Error)
}


type (
	
	VoucherUsage_Service_Register_Request interface {
			VoucherID() [16]byte    
	    MaximumInvoice() uint64 
	}

	VoucherUsage_Service_Register_Response interface {
		Numbers() uint64
	}

	VoucherUsage_Service_Count_Request interface {
		voucherID() [16]byte
		Numbers() uint64
	}

	VoucherUsage_Service_Count_Response interface {
		Numbers() uint64
	}

	VoucherUsage_Service_Get_Request interface {
		voucherID() [16]byte
		versionOffset() uint64
	}


	VoucherUsage_Service_Get_Response interface {
			VoucherUsage
		Numbers() uint64
	}
)

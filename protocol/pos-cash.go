/* For license and copyright information please see LEGAL file in repository */

package art

import (
	"../libgo/protocol"
)

// PosProvider or pos operator is to supply a particular service.
type PosCash interface {
	PosID() [16]byte                    // pos domain
	AccountID() [16]byte                // financial-account domain
	MaxCash() protocol.AmountOfMoney    //
	MaxCashTransfer() protocol.Duration // from reach the max limit
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}


type PosCash_StorageServices interface {
	Save(ps PosFinancialAccount) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(posID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(posID [16]byte, vo protocol.VersionOffset) (ps PosCash, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
		PosCash_Service_Register_Request interface {
			PosID() [16]byte
			AccountID() [16]byte
		
		}
		PosCash_Service_Register_Response interface {
			NumberOfVersion() protocol.NumberOfVersion
		}

)

type (
		PosCash_Service_Count_Request interface {
			PosID() [16]byte
		
		}
		PosCash_Service_Count_Response interface {
			NumberOfVersion() protocol.NumberOfVersion
		
		}
	
)


type (
		PosCash_Service_Get_Request interface {
			PosID() [16]byte    
			Vo() protocol.VersionOffset
		}
		PosCash_Service_Get_Response interface {
			PosCash
			NumberOfVersion() protocol.NumberOfVersion
		
		}
)
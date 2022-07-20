/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

type FinancialAccountMain interface {
	UserID() [16]byte    // user domain
	Currency() [16]byte  // financial-currency
	AccountID() [16]byte // financial-account domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type FinancialAccountMain_StorageServices interface {
	Save(fa FinancialAccountMain) (err protocol.Error)

	Count(userID, currency [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(userID, currency [16]byte, versionOffset uint64) (fa FinancialAccountMain, nv protocol.NumberOfVersion,err protocol.Error)

}


type (
		FinancialAccountMain_Service_Register_Request interface {               
		  UserID() [16]byte    
	   	Currency() [16]byte  
	   	AccountID() [16]byte 
   }
   
   FinancialAccountMain_Service_Register_Response = protocol.NumberOfVersion

)

type (
	   FinancialAccountMain_Service_Count_Request interface {
	   AccountID() [16]byte
   
   }
   FinancialAccountMain_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
   FinancialAccountMain_Service_Get_Request interface {
	   AccountID() [16]byte
	   versionOffset() uint64
   }
   
   FinancialAccountMain_Service_Get_Response1 =  FinancialAccountMain
	 FinancialAccountMain_Service_Get_Response2 = protocol.NumberOfVersion
   
)
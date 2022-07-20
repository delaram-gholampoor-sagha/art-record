/* For license and copyright information please see LEGAL file in repository */

package voucher


import (
		"../../libgo/protocol"
)


func Init() {
	protocol.App.RegisterService(&RegisterService)
	protocol.App.RegisterService(&GetService)
	protocol.App.RegisterService(&CountService)
}


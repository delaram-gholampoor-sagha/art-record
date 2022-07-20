/* For license and copyright information please see LEGAL file in repository */

package voucher

import (
	"sync"

	"github.com/Delaram-Gholampoor-Sagha/art-record/protocol"
	"../../libgo/protocol"
)


var storage vm

var _ art.Voucher_StorageServices = &storage

type vm struct {
	mutex sync.Mutex
	ids [][16]byte
	index_VoucherID map[[16]byte][]art.Voucher // PrimaryKey
}


func (s *vm) Save(v art.Voucher) (nv protocol.NumberOfVersion , err protocol.Error) {
	s.mutex.Lock()
		var voucherID = v.VoucherID()
		var records, ok = s.index_VoucherID[voucherID]
	if !ok {
		s.index_VoucherID[voucherID] = []art.Voucher{v}
		s.ids = append(s.ids, voucherID)
	} else {
		records = append(records, v)
		s.index_VoucherID[voucherID] = records
}

nv = protocol.NumberOfVersion(len(records))
	s.mutex.Unlock()
	return
}

func (s *vm) Count(voucherID [16]byte) (nv protocol.NumberOfVersion , err protocol.Error) {
	s.mutex.Lock()
	nv = len(s.index_VoucherID[voucherID])
	return
}

func (s *vm) Get(voucherID [16]byte, versionOffset uint64) (nv protocol.NumberOfVersion , v art.Voucher ,err protocol.Error) {
	s.mutex.Lock()
	v = s.index_VoucherID[voucherID][versionOffset]
	s.mutex.Unlock()
	return 
}

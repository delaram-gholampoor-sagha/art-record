/* For license and copyright information please see LEGAL file in repository */

package voucher

import (
	"database/sql"

	"github.com/Delaram-Gholampoor-Sagha/art-record/protocol"
)



var voucher_storage mysql
var _ protocol.Voucher_StorageServices = &storage

type mysql struct {
	db *sql.DB
}


func (s *mysql) Save(r protocol.Voucher) (err protocol.Error) {

	return
}


func (s *mysql) Get(r protocol.Voucher) (err protocol.Error) {
	// const query = "INSERT INTO Quiddity VALUES (?, ?, ?, ?)"
	// _, err = s.db.Exec(query, r.QuiddityID(), r.DomainID(), r.RequestID(), r.Time())
	return
}



func (s *mysql) Count(r protocol.Voucher) (err protocol.Error) {
  return
}


// go:build mysql

/* For license and copyright information please see LEGAL file in repository */

package voucher

import (
	"database/sql"

	art "github.com/Delaram-Gholampoor-Sagha/art-record/protocol"
	"../../libgo/protocol"
)


const (
	queryInsertVoucher = `INSERT INTO Voucher (VoucherID , Type , Time , RequestID) VALUES (?, ? , ? , ?);`
	queryNumberOfVersion = `SELECT COUNT(VoucherID) INTO NumberOfVersion FROM Voucher WHERE voucherID = ? ;`
	queryGetVoucher = `SELECT * FROM Voucher WHERE VoucherID = ?  LIMIT = versionOffset`
	queryCountVoucher = `SELECT COUNT(VoucherID) INTO NumberOfVersion FROM Voucher WHERE voucherID = ? ;`
)


var voucher_storage mysql
var _ art.Voucher_StorageServices = &storage

type mysql struct {
	db *sql.DB
}


func (s *mysql) Save(r art.Voucher) (err protocol.Error) {
	stmt, err := s.db.Prepare(queryInsertVoucher)
	if err != nil {
		protocol.Log("error when trying to prepare save voucher statement", err)
		return protocol.error("I don't know what to say",)
	}
	defer stmt.Close()
		insertResult, saveErr := stmt.Exec(r.Type(), r.Time(), r.RequestID())
	if saveErr != nil {
		protocol.Log("error when trying to save user", saveErr)
		return protocol.error("error when tying to save user",)
	}

	voucherId, err := insertResult.LastInsertId()
	if err != nil {
		protocol.Log("error when trying to get last insert id after creating a new voucher", err)
		return protocol.error("I don't know what to say",)
	}
	r.VoucherID() = voucherId
	return nil


	return
}


func (s *mysql) Get(r art.Voucher) (err protocol.Error) {
	stmt, err := s.db.Prepare(queryInsertVoucher)
	if err != nil {
		protocol.Log("error when trying to prepare get voucher statement", err)
		return protocol.error("I don't know what to say",)
	}

 	defer stmt.Close()
	result := stmt.QueryRow(r.VoucherID())
	if getErr := result.Scan(r.VoucherID, r.Type, r.Time, r.RequestID); getErr != nil {
		protocol.Log("error when trying to get voucher by id", getErr)
		return protocol.error("I don't know what to say",)

	}

	return
}



func (s *mysql) Count(r art.Voucher) (err protocol.Error) {
  stmt, err := s.db.Prepare(queryCountVoucher)
	if err != nil {
		protocol.Log("error when trying to prepare count voucher statement", err)
		return protocol.error("I don't know what to say",)
	}

 	defer stmt.Close()
	result := stmt.QueryRow(r.VoucherID())
	if getErr := result.Scan(r.VoucherID, r.Type, r.Time, r.RequestID); getErr != nil {
		protocol.Log("error when trying to get voucher by id", getErr)
		return protocol.error("I don't know what to say",)
	}

	return
}


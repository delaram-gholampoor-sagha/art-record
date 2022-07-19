/* For license and copyright information please see LEGAL file in repository */

CREATE TABLE `Voucher_Pos` (
  `VoucherID` BINARY(16) P,
  `PosID` BINARY(16) ,
  `Time` TIMESTAMP ,
  `RequestID` BINARY(16) 
   Key(VoucherID)
);
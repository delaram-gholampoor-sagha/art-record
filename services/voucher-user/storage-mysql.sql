/* For license and copyright information please see LEGAL file in repository */

CREATE TABLE `Voucher_User` (
  `VoucherID` BINARY(16) ,
  `UserID` BINARY(16),
  `Time` TIMESTAMP ,
  `RequestID` BINARY(16) 
   Key(VoucherID)
);

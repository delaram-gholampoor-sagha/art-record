/* For license and copyright information please see LEGAL file in repository */

CREATE TABLE `Voucher_User_Type` (
  `VoucherID` BINARY(16) ,
  `UserType` SMALLINT ,
  `Time` TIMESTAMP ,
  `RequestID` BINARY(16) 
   Key(VoucherID)
);
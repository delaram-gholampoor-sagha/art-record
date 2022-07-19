/* For license and copyright information please see LEGAL file in repository */

CREATE TABLE `Voucher_Category` (
  `VoucherID` BINARY(16) ,
  `Each` TINYINT ,
  `CategoryID` BINARY(16) ,
  `Time` TIMESTAMP ,
  `RequestID` BINARY(16) 
   Key(VoucherID)
);


/* For license and copyright information please see LEGAL file in repository */

CREATE TABLE `Voucher_Area` (
  `VoucherID` BINARY(16) ,
  `AreaID` BINARY(16) ,
  `Time` TIMESTAMP ,
  `RequestID` BINARY(16) 
   Key(VoucherID)
);

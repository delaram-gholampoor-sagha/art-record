/* For license and copyright information please see LEGAL file in repository */


CREATE TABLE `Voucher_Price` (
  `VoucherID` BINARY(16) ,
  `Each` TINYINT ,
  `Price` DECIMAL(10 , 2) ,
  `Time` TIMESTAMP ,
  `RequestID` BINARY(16) 
   Key(VoucherID)
);

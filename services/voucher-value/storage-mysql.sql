/* For license and copyright information please see LEGAL file in repository */

CREATE TABLE `Voucher_Value` (
  `VoucherID` BINARY(16) ,
  `Each` TINYINT ,
  `Percentage` DECIMAL(2, 2) ,
  `Price` DECIMAL(10 , 2) ,
  `Time` TIMESTAMP ,
  `RequestID` BINARY(16) 
   Key(VoucherID)
);




/* For license and copyright information please see LEGAL file in repository */

CREATE TABLE `Voucher_Source_Invoice` (
  `VoucherID` BINARY(16) ,
  `InvoiceID` BINARY(16) ,
  `InvoiceType` TINYINT ,
  `Time` TIMESTAMP ,
  `RequestID` BINARY(16) 
   Key(VoucherID)
);
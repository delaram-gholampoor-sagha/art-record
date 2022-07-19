/* For license and copyright information please see LEGAL file in repository */

CREATE TABLE `Voucher_Location` (
  `VoucherID` BINARY(16) ,
  `BuildingLocationID` BINARY(16) ,
  `Time` TIMESTAMP L,
  `RequestID` BINARY(16) 
   Key(VoucherID)
);
CREATE TABLE `Voucher_Location` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `BuildingLocationID` BINARY(16) NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);
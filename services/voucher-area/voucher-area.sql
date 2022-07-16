CREATE TABLE `Voucher_Area` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `AreaID` BINARY(16) NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);

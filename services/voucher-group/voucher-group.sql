CREATE TABLE `Voucher_Group` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `GroupID` BINARY(16) NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);
CREATE TABLE `Voucher_Pos` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `PosID` BINARY(16) NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);
CREATE TABLE `Voucher_Category` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `Each` tinyint NOT NULL,
  `CategoryID` BINARY(16) NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);
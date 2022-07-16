CREATE TABLE `Voucher_Product` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `Each` tinyint NOT NULL,
  `ProductID` BINARY(16) NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);
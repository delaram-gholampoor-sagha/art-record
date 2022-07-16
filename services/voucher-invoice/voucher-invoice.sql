CREATE TABLE `Voucher_Invoice` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `Each` tinyint NOT NULL,
  `MinPrice` INT NOT NULL,
  `MinAmount` INT NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);

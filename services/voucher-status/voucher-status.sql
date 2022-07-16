CREATE TABLE `Voucher_Status` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `Status` INT NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);



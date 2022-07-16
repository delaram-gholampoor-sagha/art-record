CREATE TABLE `Voucher_Duration` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `Each` tinyint NOT NULL,
  `Epoch` tinyint NOT NULL,
  `Duration` INT NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);

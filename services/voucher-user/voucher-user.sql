

CREATE TABLE `Voucher_User` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `UserID` BINARY(16) NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);

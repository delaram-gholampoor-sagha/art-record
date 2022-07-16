

CREATE TABLE `Voucher_User_Type` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `UserType` SMALLINT NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);
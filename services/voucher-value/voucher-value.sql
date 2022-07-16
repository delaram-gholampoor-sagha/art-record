
CREATE TABLE `Voucher_Value` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `Each` tinyint NOT NULL,
  `Percentage` decimal(2, 2) NOT NULL,
  `Price` decimal(10 , 2) NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);




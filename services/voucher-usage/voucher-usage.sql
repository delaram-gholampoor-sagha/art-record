
CREATE TABLE `Voucher_Usage` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `MaximumInvoice` INT NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);
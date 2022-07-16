CREATE TABLE `Voucher_Source_Invoice` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `InvoiceID` BINARY(16) NOT NULL,
  `InvoiceType` tinyint NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);
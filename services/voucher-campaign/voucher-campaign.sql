CREATE TABLE `Voucher_Campaign` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `CampaignID` BINARY(16) NOT NULL,
  `Time` timestamp NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);
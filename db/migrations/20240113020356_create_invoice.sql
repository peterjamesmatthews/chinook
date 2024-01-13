-- Create "Invoice" table
CREATE TABLE `Invoice` (
  `InvoiceId` int NOT NULL AUTO_INCREMENT,
  `CustomerId` int NOT NULL,
  `InvoiceDate` datetime NOT NULL,
  `BillingAddress` varchar(70) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `BillingCity` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `BillingState` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `BillingCountry` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `BillingPostalCode` varchar(10) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Total` decimal(10, 2) NOT NULL,
  PRIMARY KEY (`InvoiceId`),
  INDEX `IFK_InvoiceCustomerId` (`CustomerId`),
  CONSTRAINT `FK_InvoiceCustomerId` FOREIGN KEY (`CustomerId`) REFERENCES `Customer` (`CustomerId`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
